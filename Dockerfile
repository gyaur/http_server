FROM golang:1.13.5-buster

COPY server.go .

EXPOSE 7890

RUN  go build server.go

CMD ["go","run","server.go"]