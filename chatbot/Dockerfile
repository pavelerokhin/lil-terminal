FROM golang:latest

LABEL maintainer = "Pavel Erokhin <pavel.v.erokhin@gmail.com>"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build

EXPOSE 7234

CMD ["./chatbot"]
