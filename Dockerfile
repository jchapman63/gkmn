FROM golang:1.21.6

WORKDIR /app

COPY . ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /gokemon

# TODO: Go app pulls this port in with OS library?
EXPOSE 8080

# s flag starts the server
CMD ["/gokemon", "s"]
