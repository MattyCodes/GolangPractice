// This is just a dummy file to test that my Go installation works.
//
// Installing Go (on Linux Debian/Ubuntu):
// $ wget https://dl.google.com/go/go1.13.5.linux-amd64.tar.gz
// $ sudo tar -C /usr/local/ -xzf go1.13.5.linux-amd64.tar.gz
// 
// Add the following to the `~/.bashrc`:
// export PATH=$PATH:/usr/local/go/bin
// export GO111MODULE=on
//
// Installing Protobuf and gRPC Go libraries:
// $ go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
// $ go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

// Executable file.
package main

// Shorthand for "format", allows printing to console.
import "fmt"

// `main()` is automatically run for executable Go files.
// Test with `go run test.go`.
func main() {
  fmt.Println("Hello, World!")
}
