# ᚢ Heimdall ᚨ

A simple Go-based remote shell client that allows secure connections to remote servers.

## Description

Heimdall is a lightweight command-line tool written in Go that enables secure remote shell connections. It provides a simple way to establish an encrypted shell session with a remote server, preventing command interception and ensuring privacy.

## Features

- 🔒 TLS-encrypted connections
- 🔌 TCP connections
- ⚙️ Customizable host and port settings
- 💻 Interactive shell session
- 🔄 Environment variable preservation
- 🐚 Custom shell selection
- 🛡️ Protection against command interception

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

### Server Setup
First, generate TLS certificates:
```bash
# Generate private key
openssl genrsa -out server.key 2048

# Generate certificate signing request
openssl req -new -key server.key -out server.csr

# Generate self-signed certificate
openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt
```

Start the TLS server using socat:
```bash
socat OPENSSL-LISTEN:1234,cert=server.crt,key=server.key,verify=0,fork -
```

### Client Usage

Run the program with default settings (host=localhost, port=1234, shell=/bin/bash, insecure=true):
```bash
./Heimdall
```

Or specify custom host, port, shell, and insecure options:
```bash
./Heimdall --host example.com --port 8080 --shell /bin/zsh --insecure false
```

### Command Line Options

- `-H, --host`: Server host (default: "localhost")
- `-P, --port`: Server port (default: 1234)
- `-S, --shell`: User shell (default: "/bin/bash")
- `-i, --insecure`: Skip certificate verification (default: true, not recommended for production)

## Security Features

- 🔐 All communication is encrypted using TLS
- 🛡️ Commands and output cannot be intercepted by tools like tcpdump
- 🔑 Certificate-based authentication
- 🔒 End-to-end encryption

## Author

👨‍💻 Created by [thepolishdev](https://github.com/thepolishdev)

## Disclaimer

### Responsibility
The author of this repository is not responsible for any consequences arising from the use or misuse of this repository or the content provided by the third-party APIs and any damage or losses caused by users' actions.

### Educational Purposes Only
This repository and its content are provided strictly for educational purposes. By using the information and code provided, users acknowledge that they are using the APIs and models at their own risk and agree to comply with any applicable laws and regulations. 