# Go Project

This is a simple Go project.

## Getting Started

Follow these instructions to set up and run the project.

### Prerequisites

- Go (version 1.16 or later)
- Apache ActiveMQ Artemis

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/yourusername/go-challenge.git
   cd go-challenge
   ```

2. Initialize the Go module:

   ```sh
   go mod init go-challenge
   ```

3. Install dependencies:
   ```sh
   go mod tidy
   ```

### Installing Apache ActiveMQ Artemis

1. Download the latest version of Apache ActiveMQ Artemis from the [official website](https://activemq.apache.org/components/artemis/download/).

2. Extract the downloaded archive:

   ```sh
   tar -xvzf apache-artemis-<version>-bin.tar.gz
   ```

3. Navigate to the extracted directory:

   ```sh
   cd apache-artemis-<version>
   ```

4. Create a new broker instance:

   ```sh
   ./bin/artemis create mybroker
   ```

5. Start the broker:

   ```sh
   ./mybroker/bin/artemis run
   ```

### Running the Project

To run the project, execute the following command:

```sh
go run index.go
```

## License

This project is licensed under the MIT License.
