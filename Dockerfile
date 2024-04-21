# golang installation! Layer 1
FROM golang:1.21.6

# set the working directory
WORKDIR /app

# install linux basics
RUN apt-get update \
    && apt-get install -y vim \
    && apt-get install -y nano \
    && apt-get install -y curl \
    && apt-get install -y htop \
    && apt-get install -y procps \
    && apt-get install -y findutils

# copy application code into current working directory
COPY . ./
# install application golang dependencies
RUN go mod download

# build the application binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /gokemon

# server accessible port
EXPOSE 8080

# execute application binary, the s flag signals a server start
CMD ["/gokemon", "s"]
