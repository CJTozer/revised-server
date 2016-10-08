# Makefile

...OK currently it's a document...

## Generate server from Swagger

* Install `go-swagger`:
    * `go get -u github.com/go-swagger/go-swagger/cmd/swagger`
* Generate the server:
    *  `swagger generate server -f .\swagger.yaml`

## Update dependencies

* Install `govendor`:
    * `go get -u github.com/kardianos/govendor`
* Gather up all the dependencies:
    * `govendor fetch`
