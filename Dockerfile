# BUILD
FROM golang:alpine as build

LABEL maintainer="david.ullrich@spring-media.de"

RUN apk update && apk add git && apk add ca-certificates

COPY . $GOPATH/src/github.com/axelspringer/mesos-to-cloudwatch/
WORKDIR $GOPATH/src/github.com/axelspringer/mesos-to-cloudwatch

RUN echo $GOPATH
RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mesos-to-cloudwatch main.go

# RUNTIME
FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/github.com/axelspringer/mesos-to-cloudwatch/mesos-to-cloudwatch /mesos-to-cloudwatch

CMD [ "/mesos-to-cloudwatch" ]
