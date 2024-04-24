FROM golang:latest

EXPOSE 9000

RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o main .

CMD ["./main"]
