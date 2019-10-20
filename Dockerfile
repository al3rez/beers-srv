FROM golang:1.13-alpine as builder
COPY . /go/src/github.com/azbshiri/beers
WORKDIR /go/src/github.com/azbshiri/beers
RUN go install ./... \
    && apk add git coreutils \
    && GO111MODULE=off go get github.com/grpc-ecosystem/grpc-health-probe

FROM alpine
WORKDIR /bin/
RUN apk add --no-cache bash ca-certificates
COPY --from=builder /go/bin/beers-cli /go/bin/beers-server /go/bin/grpc-health-probe ./

ENV PORT=8000
EXPOSE 8000
CMD ["beers-server"]
