FROM golang:1.18-alpine

RUN mkdir /app

ADD . /app

WORKDIR  /app

RUN go build -o page_visit .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary
ENTRYPOINT ["./page_visit"]



