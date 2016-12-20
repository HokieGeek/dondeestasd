package dondeestas

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

type personDataRequest struct {
	Ids []string `json:"ids"`
}

type personDataResponse struct {
	People []Person `json:"people"`
}

func personRequestHandler(log *log.Logger, db *DbClient, w http.ResponseWriter, r *http.Request) {
	var req personDataRequest

	if bytes, err := httputil.DumpRequest(r, true); err == nil {
		log.Println(string(bytes))
	}

	if err := readCloserJSONToStruct(r.Body, &req); err != nil {
		http.Error(w, fmt.Sprintf("%s\n", err), http.StatusUnprocessableEntity)
	} else {
		log.Printf("Received request for people with ids: %v\n", req.Ids)

		var resp personDataResponse
		resp.People = make([]Person, 0)

		for _, id := range req.Ids {
			if person, err := (*db).Get(id); err == nil {
				resp.People = append(resp.People, *person)
			}
		}

		if len(resp.People) == len(req.Ids) {
			postJSON(w, http.StatusOK, resp)
		} else {
			postJSON(w, http.StatusPartialContent, resp)
		}
	}
}

func updatePersonHandler(log *log.Logger, db *DbClient, w http.ResponseWriter, r *http.Request) {
	var update Person

	if bytes, err := httputil.DumpRequest(r, true); err == nil {
		log.Println(string(bytes))
	}

	if err := readCloserJSONToStruct(r.Body, &update); err != nil {
		http.Error(w, fmt.Sprintf("%s\n", err), http.StatusUnprocessableEntity)
	} else {
		log.Printf("Received update for person with id: %s\n", update.ID)

		var err error
		if (*db).Exists(update.ID) {
			err = (*db).Update(update)
		} else {
			err = (*db).Create(update)
		}

		if err != nil {
			log.Printf("ERROR: %d: %s\n", http.StatusInternalServerError, err)
			http.Error(w, fmt.Sprintf("%s\n", err), http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	}
}

func postJSON(w http.ResponseWriter, httpStatus int, send interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "PUT")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(httpStatus)
	if err := json.NewEncoder(w).Encode(send); err != nil {
		return err
	}

	return nil
}

// ListenAndServe opens an HTTP server which listens for connections and provides
// data using the given DbClient.
//
// The following routes are available:
//  /person
//        This route returns data on the request people. It expects a JSON body with the following format:
//        { ids: []{ "identifier" } }
//        Where "identifier" is a string used to identify the people being requested
//        It returns a JSON with an array of Person objects
//
//  /update
//        This route expects a JSON body with a single Person object to update
func ListenAndServe(log *log.Logger, port int, db *DbClient) error {
	http.HandleFunc("/person",
		func(w http.ResponseWriter, r *http.Request) {
			personRequestHandler(log, db, w, r)
		})

	http.HandleFunc("/update",
		func(w http.ResponseWriter, r *http.Request) {
			updatePersonHandler(log, db, w, r)
		})

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
