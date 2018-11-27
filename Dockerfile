# BUILD
FROM golang:latest as build

LABEL maintainer="david.ullrich@spring-media.de"

COPY . $GOPATH/src/github.com/axelspringer/mesos-to-cloudwatch/
WORKDIR $GOPATH/src/github.com/axelspringer/mesos-to-cloudwatch

RUN echo $GOPATH
RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mesos-to-cloudwatch main.go

# RUNTIME
FROM scratch

COPY --from=build /go/src/github.com/axelspringer/mesos-to-cloudwatch/mesos-to-cloudwatch /mesos-to-cloudwatch

CMD [ "/mesos-to-cloudwatch" ]
