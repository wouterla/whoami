test:
	@go test -v .

run:
	@WHOAMI_SERVICE=supersonic WHOAMI_VERSION=1.0-alpha WHOAMI_IMAGE=foo WHOAMI_COUNT=3 go run http.go
