# Golang Web Assembly Calculator
A very simple calculator written in Golang and compiled as Web Assembly.

## Usage
Build the server binary and compile wasm
```bash
$ make
```
Run the server the relevant server binary for your platform from the `bin`
directory. Usage for the server is as follows:
```bash
Usage of ./bin/server-linux-amd64:
  -dir string
        The directory to serve files from (default ".")
  -port string
        The port to serve on (default "8080")
```
Finally open up your browser and navigate to http://localhost:8080/index.html