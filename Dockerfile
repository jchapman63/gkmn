# Explain installing dependencies!
FROM golang:1.21.6

WORKDIR /app


RUN apt-get update \
    && apt-get install -y vim \
    && apt-get install -y nano \
    && apt-get install -y curl \
    && apt-get install -y htop \
    && apt-get install -y procps \
    && apt-get install -y findutils


COPY . ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /gokemon

# TODO: Go app pulls this port in with OS library?
EXPOSE 8080

# s flag starts the server
CMD ["/gokemon", "s"]
