# beers
Enjoy serving 🍻 through micro-services using gRPC  
<br>
<br>
[![asciicast](https://asciinema.org/a/XXLhQTinGqJdn7F5YR00Rp171.svg)](https://asciinema.org/a/XXLhQTinGqJdn7F5YR00Rp171)
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

### Docker
If you have Docker already installed then you don't need Go or Protocol Buffers you just need to run below command:

```
$ docker build -t beers-srv:0.1.0 .
Sending build context to Docker daemon  360.4kB
Step 1/11 : FROM golang:1.13-alpine as builder
 ---> 4acab7f5278b
Step 2/11 : COPY . /go/src/github.com/azbshiri/beers
...

$ docker run -d --rm --name=beers-srv beers-srv:0.1.0
$ docker exec -i beers-srv grpc-health-probe -addr=:8000
status: SERVING

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


## Event sorucing
I'd go with Google Pub/Sub API to keep track of events on cluster in an event store and https://watermill.io/ API to make it easier to deal with events/subscribers.

- BeerRemoved
- BeerAdded
- BeerNameChanged
etc


## Cloud naviteness
Using Kubernetes and Docker you can easily run this applicaiton on cloud with ease scaling up/down, add database integeration and etc.

## The Twelve-Factor App

> I. Codebase
One codebase tracked in revision control, many deploys
II. Dependencies
Explicitly declare and isolate dependencies

Using Go modules I've isolated dependencies as well as moved out shared code into separate codebase with semantic versioning also since this is an small micro-service I decided to use one branch `master` for deployment and build and adding features using separate branches using PRs.

> III. Config
Store config in the environment

There's not much of a configuration in this micro-service but I've used env variables to setup the minimum.

> IV. Backing services
Treat backing services as attached resources

There's no backing services (database, etc) but otherwise still using env vars and service discovery tools one can easily decople attached services and their communication.

> V. Build, release, run
Strictly separate build and run stages

`master` is now reilable stage for building and release also using GitHub actions I run linting, testing and building for each PR.

> VI. Processes
Execute the app as one or more stateless processes

Since I'm using in-memory allocation and no database is envovled so this app is not stateless but because I've already implemented database abstraction layer and repository pattern it's just so easy to configure databases and other backing services using env vars / service discovery to make this app stateless. (e.g. store removing, adding into separate process pg, etc)

> VII. Port binding
Export services via port binding

By setting `PORT` you can easily change port of services (in this case only gRPC server)

> VIII. Concurrency
Scale out via the process model

OCI orchestration is a factor here as you can scale down/up easily on clusters (GCP, AWS) using Docker and K8s.

> IX. Disposability
Maximize robustness with fast startup and graceful shutdown

Using K8s and Docker I've implemented gRPC health protocol so K8s can easily restart/replace service instances when needed.

> X. Dev/prod parity
Keep development, staging, and production as similar as possible

As I have only one `master` branch and many sub-branches this is also true

> XI. Logs
Treat logs as event streams

Unfortunately I haven't had time to make this consitent but I guess in some cases I'm treating logs like streams?!

> XII. Admin processes
Run admin/management tasks as one-off processes

When run using Docker it's simple as running `docker exec` to do admin tasks also when developing using `beers-cli` and `grpc-health-probe` to make sure everything is working fine.


## Things that are missed 

- [ ] Move protocol buffers generation/linting to Github actions
- [ ] Fix GoReleaser on Github actions (checkout action doesn't sync tags)
- [ ] Do dependency injection for logging and make it consistent
- [ ] More human-friendly errors when dealing with 
- [ ] Add GCP/AWS integration using Terraform
- [ ] Manage protocol buffers using [buf](https://buf.build/docs/introduction?ref=producthunt) or [gunk](https://github.com/gunk/gunk)
- [ ] Service discovery using etcd, consul or etc
