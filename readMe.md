# GOKEMON

Welcome to Gokemon, a CLI client-server application for having
pokemon battles with a friend in the terminal. Here in -v0.0.1 , there
is very little battle functionality. But! This project is a proof of concept for
creating a local deployment of a multi-container application, the first one I have ever created.

## Getting Started

Docker Desktop and an image of postgres is required (along with a build of the game image).

**Images**
From project root

- `docker build -t gokemon-build .`
- `docker pull postgres`

**Running the Application**
1.) Pull down or download the application code
2.) Run `go run .` from project root
3.) Select `host`

Now a container of the server will spin up, along with a postgres container. The
two should be connected successfully as this happened with `docker compose up -d` from within the
application code.

Finally, you can open a second terminal window:

1.) `go run .`
2.) `connect`

Enjoy!
