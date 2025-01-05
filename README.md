# NetCat Project in Go

## Overview
This project implements a recreation of the NetCat (nc) command-line utility using Go. The utility operates in a **Server-Client Architecture**, allowing a server to listen for incoming connections and clients to connect, enabling communication across a network. The application mimics the functionality of the original NetCat command, creating a group chat system.

## Objectives
The main objectives of the project include:

1. Recreating NetCat functionality in a Server-Client Architecture.
2. Enabling a server to run on a specified port and listen for incoming client connections.
3. Allowing clients to connect to the server and communicate via TCP.
4. Creating a group chat environment with advanced functionality.

## Features
The project includes the following features:

### General Functionality
- **TCP Connection:** Implements a 1-to-many TCP connection between the server and multiple clients.
- **Name Requirement:** Clients must provide a name to join the chat.
- **Connection Control:** The server controls the number of simultaneous connections (maximum 10 clients).
- **Message Broadcasting:** Messages sent by a client are broadcasted to all connected clients, excluding empty messages.

### Message Formatting
- Messages are formatted with a timestamp and the sender's name, e.g.,
  ```
  [2025-01-01 12:00:00][Alice]: Hello, world!
  ```

### State Management
- **Message History:** New clients receive all previously sent messages upon joining the chat.
- **Join/Leave Notifications:**
  - When a client joins the chat, other clients are informed.
  - When a client leaves the chat, the rest of the clients are notified, but they remain connected.
- **Resilient Connections:** The server and remaining clients continue to operate normally if a client disconnects.

### Port Handling
- If no port is specified, the default port **8989** is used.
- If an invalid or missing port argument is provided, the program responds with a usage message:
  ```
  [USAGE]: ./TCPChat $port
  ```

### Implementation Details
- **Concurrency:** Uses Go-routines to handle simultaneous client connections and communication.
- **Synchronization:** Implements channels or Mutexes to synchronize data across clients and the server.
- **Error Handling:** Handles errors gracefully on both server and client sides.
- **Testing:** Includes unit tests for validating server and client functionality.

## Instructions

### Server Mode
To start the TCP server:
1. Compile the program using `go build`.
2. Run the server with a specified port or use the default port:
   ```
   ./TCPChat 8989
   ```

### Client Mode
To start a client:
1. Run the program and specify the server address and port:
   ```
   ./TCPChat <server_address>:<port>
   ```
2. Provide a name when prompted to join the chat.

### Development Notes
- The project is written in Go and adheres to Go best practices.
- Maximum connections are limited to 10 clients.
- The server must handle various errors, such as invalid ports, disconnected clients, and malformed messages.

### Testing
Unit tests are provided to:
1. Validate the server’s ability to handle connections and broadcasts.
2. Ensure clients can join, send, and receive messages properly.

## Example Usage

### Server Output
```
Server started on port 8989...
[2025-01-01 12:00:00] Alice joined the chat.
[2025-01-01 12:01:00] Bob joined the chat.
[2025-01-01 12:02:00][Alice]: Hello, Bob!
[2025-01-01 12:03:00][Bob]: Hi Alice!
[2025-01-01 12:04:00] Bob left the chat.
```

### Client Output
```
Welcome to the chat!
Messages:
[2025-01-01 12:00:00][Alice]: Hello, Bob!
[2025-01-01 12:01:00][Bob]: Hi Alice!

[2025-01-01 12:04:00] Bob left the chat.
```

## Limitations
- This project only supports TCP connections; UDP is not implemented.
- Connections are limited to 10 simultaneous clients.

## References
- [NetCat Manual](https://linux.die.net/man/1/nc)
- [Go Concurrency Patterns](https://go.dev/doc/effective_go#concurrency)

## License
This project is licensed under the MIT License. See the LICENSE file for details.

