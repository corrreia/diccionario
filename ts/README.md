# Diccionario

A simple English word list control plane server.

## Typescript Instructions

From the root of this repo run:

```sh
docker build -f ts/Dockerfile -t diccionario .
docker run -it -p 5000:5000 -v ./ts:/usr/src/diccionario diccionario
```

The server will be available at http://localhost:5000.
It will automatically reload when you make changes to the source code.

To stop the server, press Ctrl+C in the terminal where it's running.

To access the running container, run:

```sh
docker exec -it `docker ps | grep diccionario | awk '{print $1}'` bash
```

To exit the terminal session, press Ctrl+D in the terminal where it's running.
