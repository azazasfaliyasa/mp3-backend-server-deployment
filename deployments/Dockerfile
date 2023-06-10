FROM golang:alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o mini-project-3-azaz ./cmd

EXPOSE 8080

CMD ["./mini-project-3-azaz"]
