FROM golang:1.17

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY main.go ./

RUN go build .

EXPOSE 5000

CMD [ "./go-simple" ]
