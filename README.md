# beers
Enjoy serving 🍻 through micro-services using gRPC
[![asciicast](https://asciinema.org/a/PvdZwgpwX5rv4vyJ6cd91s9Qh.svg)](https://asciinema.org/a/PvdZwgpwX5rv4vyJ6cd91s9Qh)
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
Installing `beers-srv` is easy, using go get you can install the cmd line app `beers-cli` to interact with gRPC server. First you'll need Google's Protocol Buffers installed.
```
$ brew install protobuf
$ go get -u github.com/azbshiri/beers-srv/...
```



## Testing
I haven't used any external libraries for testing/diffing so if you already have latest version Go installed,
just simply run `go test -v ./...` :
```
ok      github.com/azbshiri/beers-srv/pkg/adding    (cached)
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

## Kubernetes
Deploying gRPC applications to K8s and the best way to configure health checks is using `grpc-health-probe` and GRPC Health Checking Protocol
https://kubernetes.io/blog/2018/10/01/health-checking-grpc-servers-on-kubernetes/

```
$ kubectl apply -f deployment.yml
$ kubectl get svc beers-srv
NAME        TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
beers-srv   NodePort   10.110.71.154   <none>        8000:30768/TCP   15m
```

executing “grpc_health_probe” will call our gRPC server over localhost (Cluster IP):

```
$ minikube status
kubectl: Correctly Configured: pointing to minikube-vm at 192.168.122.118

$ grpc-health-probe --addr=192.168.122.118:30768
status: SERVING
```

kaboom!
