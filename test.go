// This is just a dummy file to test that my Go installation works.

// Installing Protobuf and gRPC Go libraries:
// $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
// $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

// Executable file.
package main

// Shorthand for "format", allows printing to console.
import "fmt"

// `main()` is automatically run for executable Go files.
// Test with `go run test.go`.
func main() {
  fmt.Println("Hello, World!")
}
