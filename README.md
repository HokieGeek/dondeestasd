# ¿Dónde Estás? Daemon [![Build Status](https://travis-ci.org/HokieGeek/donde-estas-daemon.svg?branch=master)](https://travis-ci.org/HokieGeek/donde-estas-daemon) [![Coverage](http://gocover.io/_badge/github.com/HokieGeek/donde-estas-daemon)](http://gocover.io/github.com/HokieGeek/donde-estas-daemon) [![GoDoc](http://godoc.org/github.com/HokieGeek/donde-estas-daemon?status.png)](http://godoc.org/github.com/HokieGeek/donde-estas-daemon)
The server side to the [¿Dónde Estás?](https://github.com/HokieGeek/DondeEstas) android app

##### Suggested usage
```sh
docker run -d --name couchdb couchdb
docker run -d -p 8080:8080 --link couchdb:db hokiegeek/donde-estas-daemon
```
