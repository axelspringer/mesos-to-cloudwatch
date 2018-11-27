# BUILD
FROM golang:latest as build

LABEL maintainer="david.ullrich@spring-media.de"

WORKDIR /go/src/github.com/axelspringer/mesos-to-cloudwatch
COPY . .

RUN echo $GOPATH
RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mesos-to-cloudwatch main.go

# RUNTIME
FROM scratch

COPY --from=build /go/src/github.com/axelspringer/mesos-to-cloudwatch/mesos-to-cloudwatch /mesos-to-cloudwatch
RUN cp /go/src/github.com/axelspringer/mesos-to-cloudwatch/mesos-to-cloudwatch /mesos-to-cloudwatch
CMD [ "/mesos-to-cloudwatch" ]
