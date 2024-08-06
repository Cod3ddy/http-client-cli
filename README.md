# HTTP Client CLI

HTTP Client CLI is a command line interface application written in Go. It allows you to perform simple HTTP requests using the Cobra-CLI library.

## Getting Started

To run the application, use the `go run` command followed by the `main.go` file and the specific command you want to execute. For example:

```bash
go run main.go router post --url "http://localhost:4000/example/login" --data "{}"
```

## Supported Commands

Currently, the application supports the following commands:

- `post`: Sends a POST request to the specified URL with the provided data.
- `get`: Sends a GET request to the specified URL.

### POST Request

To send a POST request, use the `post` command followed by the `--url` flag to specify the URL and the `--data` flag to provide the data. The data should be in JSON format. For example:

```bash
go run main.go router post --url "http://localhost:4000/example/login" --data "{username: jondoe, password:123doe}"
```

You can also provide the data by dragging and dropping a JSON file onto the terminal or typing in the path to a `.json` file. For example:

```bash
go run main.go router post --data C:\User\example\Desktop\cred.json --url "http://localhost:4000/example/login"
```

### GET Request

To send a GET request, use the `get` command followed by the `--url` flag to specify the URL. For example:

```bash
go run main.go router get --url "http://localhost:4000/example/login"
```

## Future Improvements

The following features are planned for future releases:

- Support for GET method with additional parameters like headers.
- Implementation of DELETE request.
- Ability to send UPDATE request.