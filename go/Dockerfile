FROM golang:1.25.4

EXPOSE 8080

WORKDIR /usr/src/diccionario

COPY go/go.mod go/go.sum /usr/src/diccionario/
COPY words.txt /words.txt
RUN go mod download

RUN go install github.com/air-verse/air@latest

CMD ["air", "-c", ".air.toml"]
