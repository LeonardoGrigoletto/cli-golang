
# Command Line Interface Application

This project is a Command Line Interface (CLI) application written in Go. It performs URL-based queries to retrieve IP addresses and name servers (NS) associated with a given host.

## Features

- **IP Lookup:** Retrieves IP addresses associated with a specified host.
- **Name Server Lookup:** Lists the name servers (NS) associated with a specified host.

## Prerequisites

- **Go**: Ensure that Go is installed on your system. You can download it from [Go's official website](https://go.dev/dl/).

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/repository-name.git
   ```
2. Navigate to the project directory:
   ```bash
   cd repository-name
   ```
3. Build the application:
   ```bash
   go build -o cli-app
   ```

## Usage

After building the application, you can execute it using the following commands:

### IP Lookup

Retrieve the IP addresses of a given host:
```bash
./cli-app ip --host=<HOST>
```

Example:
```bash
./cli-app ip --host=google.com
```

### Name Server Lookup

Retrieve the name servers of a given host:
```bash
./cli-app servers --host=<HOST>
```

Example:
```bash
./cli-app servers --host=google.com
```

## Example Output

### IP Lookup

Command:
```bash
./cli-app ip --host=google.com
```

Output:
```
64.233.160.100
64.233.160.101
64.233.160.102
```

### Name Server Lookup

Command:
```bash
./cli-app servers --host=google.com
```

Output:
```
ns1.google.com
ns2.google.com
ns3.google.com
```

## Dependencies

This application uses the following Go library:

- [urfave/cli](https://github.com/urfave/cli): A library for building command-line interfaces in Go.

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests to improve this project.

## License

This project is licensed under the [MIT License](LICENSE).