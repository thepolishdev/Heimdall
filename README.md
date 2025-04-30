# Heimdall 🔒

A simple Go-based remote shell client that allows connections to remote servers.

## Description

Heimdall is a lightweight command-line tool written in Go that enables remote shell connections. It provides a simple way to establish a shell session with a remote server.

## Features

- 🔌 TCP connections
- ⚙️ Customizable host and port settings
- 💻 Interactive shell session
- 🔄 Environment variable preservation

## Installation

1. Make sure you have Go installed (version 1.16 or higher)
2. Clone the repository:
   ```bash
   git clone https://github.com/thepolishdev/Heimdall.git
   cd Heimdall
   ```
3. Build the project:
   ```bash
   go build
   ```

## Usage

Run the program with default settings (localhost:1234):
```bash
./Heimdall
```

Or specify custom host and port:
```bash
./Heimdall --host example.com --port 8080
```

### Command Line Options

- `-H, --host`: Server host (default: "localhost")
- `-P, --port`: Server port (default: 1234)

## Author

👨‍💻 Created by [thepolishdev](https://github.com/thepolishdev) 