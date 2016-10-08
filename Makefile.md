# Makefile

...OK currently it's a document...

## Generate server from Swagger

* Install `go-swagger`:
    * `go get -u github.com/go-swagger/go-swagger/cmd/swagger`
* Generate the server:
    *  `swagger generate server -f .\swagger.yaml --skip-support`
    *  or just `go generate`

## Update dependencies

* Install `govendor`:
    * `go get -u github.com/kardianos/govendor`
* Gather up all the dependencies:
    * `govendor fetch`

## Test Heroku locally

* Install `heroku`:
    * Installer linked [here](https://devcenter.heroku.com/articles/heroku-command-line)
* Run locally:
    * *Commit changes*
    * `heroku local web`

## Formatting

* `gofmt -w .\restapi\configure_revised.go`
* `gofmt -w .\main.go`
