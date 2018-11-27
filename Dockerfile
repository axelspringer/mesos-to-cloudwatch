# BUILD
#FROM golang:latest as build
FROM golang:latest

LABEL maintainer="david.ullrich@spring-media.de"

WORKDIR /go/src/github.com/axelspringer/mesos-to-cloudwatch
COPY . .

RUN echo $GOPATH
RUN go get -d -v ./...

#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o swerve -v -ldflags "-extldflags '-static'" -a -installsuffix cgo main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mesos-to-cloudwatch main.go

# RUNTIME
#FROM scratch

#MAINTAINER Jan Michalowsky <sejamich@googlemail.com>

#COPY --from=build /go/src/github.com/axelspringer/swerve/swerve /swerve
RUN cp /go/src/github.com/axelspringer/mesos-to-cloudwatch/mesos-to-cloudwatch /mesos-to-cloudwatch
CMD [ "/mesos-to-cloudwatch" ]
