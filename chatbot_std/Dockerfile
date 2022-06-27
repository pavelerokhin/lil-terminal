FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build

EXPOSE 7234

CMD ["./chatbot"]
