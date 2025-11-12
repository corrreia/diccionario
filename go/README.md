# Diccionario

A simple English word list control plane server.

## Golang Instructions

From the root of this repo run:

```sh
docker build -f go/Dockerfile -t diccionario .
docker run -it -p 8080:8080 -v ./go:/usr/src/diccionario diccionario
```

The server will be available at http://localhost:8080.
It will automatically reload when you make changes to the source code
(it's using [air](https://github.com/air-verse/air)).

To stop the server, press Ctrl+C in the terminal where it's running.

To access the running container, run:

```sh
docker exec -it `docker ps | grep diccionario | awk '{print $1}'` bash
```

To exit the terminal session, press Ctrl+D in the terminal where it's running.
