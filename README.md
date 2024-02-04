# Bitcoin CLI Transaction Decoder

This repository contains a Go program that decodes Bitcoin raw transactions using the Bitcoin Core CLI's `decoderawtransaction` command. It provides a convenient way to extract detailed information about transactions, including inputs, outputs, values, and more.

## Features

- **Transaction Decoding**: Decode raw Bitcoin transactions into a human-readable format.
- **Unit Tests**: Includes comprehensive unit tests to ensure accurate decoding.
- **Flexible**: The program allows you to input any valid Bitcoin raw transaction hex for decoding.

## Getting Started

To use the Bitcoin CLI Transaction Decoder, follow these steps:

1. Clone the repository to your local machine.
2. Run the `go run main.go` command, providing the Bitcoin raw transaction hex as an argument.
3. Run `go get -u github.com/btcsuite/btcd/wire` to install the `btcd` package.
4. The program will decode the transaction and display the results in the terminal.

Example:
```bash
go run main.go
