
# **Important**
**Botzilla Server is no longer needed for Botzilla and is no longer supported.**

# Botzilla-Server

Botzilla-Server is a server-side application written in Go to handle core functions performed by Botzilla. It runs a TCP server for communication to register and manage components. The message routing in the Botzilla server allows the server to handle different operation based on the operation codes.

# Features

## TCP Server

The server listens TCP connections on a port and handles each connection in a separate go routine. The server checks the message integrity using a secretKey. An incoming message must contain a hash that is verified by the server to perform operations.

## Component Handling

The Component struct is used to assign an address, group and a name. New components use this structure to register themselves. Additionally, the server maintains a registry of the components for returning and assigning groups.

## Message routing

The router processes the incoming messages to the server and routes them to the handler based on the operation code in the message.

# Functions

## `tcp.go`

```
StartTCPServer(port int, secretKey string)
```

- `Description`: Starts the tcp server and listens on the port specified.

- `Parameters`:
  - `port`: The TCP port server is listening on.
  - `secretKey` : The secretKey for verifying the hash from an incoming message.
- `Returns`: Returns an error
