# HTTPSniffer - Go HTTP Logger

HTTPSniffer is a Go package that provides a middleware for logging HTTP requests and responses. It captures important information such as request method, URL. It offers enhanced readability and enables quick identification of different request methods.

## Installation

To use HTTPSniffer in your Go project, you need to have Go installed and set up. Once you have Go installed, you can install the package using the following command:

```bash
go get github.com/majid-2002/go-http-sniffer
```

## Usage

Follow the steps below to use the HTTPSniffer package in your Go application:

1. Import the necessary packages:

   ```go
   import (
   	"log"
   	"net/http"
   	"os"
   
   	"github.com/majid-2002/go-http-sniffer"
   )
   ```

2. Create a logger instance using `NewHTTPLogger` and pass your desired HTTP handler:

   ```go
   logger := httpsniffer.NewHTTPLogger(yourHTTPHandler)
   ```

3. Use the `logger` instance as middleware in your HTTP server:

   ```go
   http.ListenAndServe(":8080", logger)
   ```

   Replace `yourHTTPHandler` with your existing HTTP handler or router.

4. Start your server, and the HTTPSniffer middleware will automatically log each incoming request and its corresponding response.

## Example

For a detailed example of how to use the HTTPSniffer package, please refer to the [example folder](https://github.com/majid-2002/go-http-sniffer/tree/main/example). 

The example folder contains a sample Go file that demonstrates the usage of the HTTPSniffer package in a simple HTTP server implementation.


## Acknowledgments

HTTPSniffer is inspired by various Go HTTP logging middleware implementations and aims to provide a simple and customizable solution for logging HTTP requests and responses.

## Contributing

Contributions to HTTPSniffer are welcome! If you encounter any issues or have suggestions for improvements, please open an issue on the GitHub repository: [https://github.com/majid-2002/go-http-sniffer/issues](https://github.com/majid-2002/go-http-sniffer/issues)
