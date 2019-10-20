# beers
Enjoy having beers using gRPC and domain-driven design


## Stack

- **CI/CD**: Github Actions (e.g Review, Release automation `goreleaser`, etc)
- **Testing**: Go `testing` package
- **Serialization**: Protocol Buffers `protobuf`
- **OCI orchestration**: Kubernetes, Docker


## Architecture
Following DDD principles I've separated bounded contextes and behaviors (e.g. adding, removing, serializing, grpc) into multiple packages using
Service Objects, Database Layer Abstraction, Repository pattern, CLI interafces so that
decopling behaviors lead to easier testing and shipping new features.

## Installation
Installing `beers` is easy, using go get you can install the cmd line app `beers-cli` to interact with gRPC server. First you'll need Google's Protocol Buffers installed.
```
$ brew install protobuf
$ go get -u github.com/azbshiri/beers/...
```



## Testing
I haven't used any external libraries for testing/diffing so if you already have latest version Go installed,
just simply run `go test -v ./...` :
```
ok      github.com/azbshiri/beers/pkg/adding    (cached)
=== RUN   Test_Add
=== RUN   Test_Add/adds_beer
=== RUN   Test_Add/checks_against_blank_name
=== RUN   Test_Add/fails_due_to_service
--- PASS: Test_Add (0.00s)
    --- PASS: Test_Add/adds_beer (0.00s)
    --- PASS: Test_Add/checks_against_blank_name (0.00s)
    --- PASS: Test_Add/fails_due_to_service (0.00s)
PASS
ok 
```
