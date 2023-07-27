# Tcp Chat

## Description

This project consists on recreating the NetCat in a Server-Client Architecture that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.

## Usage

1. Command to running tcp server 

```
$ go run cmd/main.go
```

Server will start at **8989** port

-   Your also able to print your own port

Or, by audit case, you can run server by instruction below

```
$ go build -o tcp cmd/main.go
```

You will have a exe file that you can run through the command

```
$ ./tcp $PORT
```

2. Open another terminal, and connect to the chat

```
$ nc localhost $PORT
```

## Authors

@zsandibe
