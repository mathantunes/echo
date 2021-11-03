# Echo

Implementation of the [Echo Protocol](https://datatracker.ietf.org/doc/html/rfc862) in Golang on top of TCP

## Executing

`go run ./app/app.go`

Can also run in Docker.

* `--port` to specify a port in which to listen to incoming messages (defaults to 7)

## Testing

* Unit testing
* `go run ./test/test.go` is a tcp client that reads from stdin, sends to running `app` and outputs the `echo` on stdout.


## CI:

* CI builds and tests each commit and PR to `main` branch

# TODO:
* Package the new docker image and store on github container registry
