// Executable file.
package main

// Import libraries.
import (
  "log"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  pb "github.com/MattyCodes/GolangPractice/src/application"
)

// Executable function.
func main() {
  // Instantiate the gRPC client connection.
  // Interesting note regarding the use of `&` and `*`:
  // - `&` denotes a "memory address", a representation of where a value or object
  //   is stored in memory, but it does _not_ give us the value itself.
  // - `*`, when placed before a memory address, will actually fetch the value or
  //   object from memory.
  var conn *grpc.ClientConn

  // Establish a connection to our server via gRPC.
  conn, err := grpc.Dial(":9000", grpc.WithInsecure())

  // If a connection could not be established, log the error.
  if err != nil {
    log.Fatalf("did not connect: %s", err)
  }

  // `defer` is funky. My understanding is that `defer` defines an action which
  // will be run no matter what, _after_ a function has finished doing everything
  // else. That is, when `main()` has finished running (including all of the lines
  // that come after this `defer` call), then `conn.Close()` will run.
  //
  // Also worth noting that mulitple `defer` calls can be made in a function, and
  // will be executed in "last in, first out" priority; a call to `defer` on line
  // 45 will be run _before_ a `defer` call on line 20.
  defer conn.Close()

  // Instantiate our service client - this client is automatically generated
  // from our Protobuf definition.
  c := pb.NewShowRecommendationServiceClient(conn)

  // Insantiate our "show recommendation request" message, which the `RecommendShow`
  // function takes as an argument.
  request := &pb.ShowRecommendationRequest{ Name: "Matty" }

  // Invoke the service's `RecommendShow` function, providing the request context
  // (the automatically generated client requires this), and our request message.
  response, err := c.RecommendShow(context.Background(), request)

  // Print out whatever the server returned to us.
  log.Printf("Response from server: %s", response)
}
