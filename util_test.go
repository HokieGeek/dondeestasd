package dondeestas

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func stringToReadCloser(str string) io.ReadCloser {
    return ioutil.NopCloser(bytes.NewReader(bytes.NewBufferString(str)))
}

func TestReadCloserToJson(t *testing.T) {
    // TODO: make this just a JSON object, not a Person
	expectedPerson, _ := createRandomPerson()
	expectedPersonJson, _ := json.Marshal(expectedPerson)
	expectedPersonStr := string(expectedPersonJson)

	var person Person
	if err := ReadCloserJsonToStruct(stringToReadCloser(expectedPersonStr), person); err != nil {
		t.Fatalf("Encountered error when retrieving json from string: %s", err)
	}
	
	t.Skip("TODO")
	
	// Test forcing the function to read a closed stream
	/*
    req = httptest.NewRequest("GET", "http://blah.com/foo", bytes.NewBufferString(expectedPersonStr))
	if err := req.Body.Close(); err != nil {
		t.Fatalf("Could not close test request body!")
	}
	if err := ReadCloserJsonToStruct(req.Body, person); err == nil {
		t.Error("Did not receive expected error when reading closed stream")
	}
	
	// Incorrect JSON object
	if err := ReadCloserJsonToStruct(ioutil.NopCloser(bytes.NewReader(bytes.NewBufferString(`{"id":"foo"}`).Bytes())), person); err == nil  {
		t.Error("Did not receive expected error on bad JSON unmarshalling")
	}
    */
}
