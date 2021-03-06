// Our server file - followed along with a few different tutorials while working
// through this:
// 1. https://docs.dapr.io/developing-applications/building-blocks/service-invocation/howto-invoke-services-grpc/
// 2. https://tutorialedge.net/golang/go-grpc-beginners-tutorial/
// 3. https://grpc.io/docs/languages/go/basics/

// Executable file.
package main

// Import libraries.
import (
  "fmt"
  "net"
  "google.golang.org/grpc"
  "golang.org/x/net/context"
  "github.com/MattyCodes/GolangPractice/shows"
  pb "github.com/MattyCodes/GolangPractice/src/application"
)

// This is a `struct` which implements the `UnimplementedShowRecommendationServiceServer`
// interface - that interface is automatically generated with Protobuf.
// This server is leveraged specifically to serve up our `ShowRecommendationService`.
type server struct {
  pb.UnimplementedShowRecommendationServiceServer
}

// This function is leveraged by our service (see usage in the `application.proto` file).
// The context here is required by default by our Protobuf-generated client. Additionally,
// the function will take a `RecommendationRequest` argument and provide a `ShowRecommendation`
// in return. The client also expects that an `error` can be returned here, though I have not
// added any actual error handling.
func (s *server) RecommendShow(ctx context.Context, recommendationRequest *pb.ShowRecommendationRequest) (*pb.ShowRecommendation, error) {
  // Log that the request was received, and is being handled by this function as expected.
  fmt.Println("Recommendation request received.")

  // Fetch the "name" attribute from the recommendation request. The `GetName` function
  // is automatically generated from our Protobuf definition.
  name := recommendationRequest.GetName()

  // Leverage our "shows" package to randomly select a show name.
  show := shows.SelectRandom()

  // Instantiate a `ShowRecommendation` message and return it; the `nil` returned
  // with it is our `error` object, since we are not doing any actual error handling here.
  return &pb.ShowRecommendation{ Name: name, Show: show }, nil
}

// Executable function.
func main() {
  fmt.Println("Server running...")

  // Set up a listener for TCP connections on port 9000.
  lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))

  // If an error occurred with our listener, log the error.
  if err != nil {
    fmt.Println("failed to listen: %v", err)
  }

  // Instantiate a gRPC server.
  grpcServer := grpc.NewServer()

  // Register our service, so that we can actually serve up show recommendations
  // using the previously defined TCP connection.
  pb.RegisterShowRecommendationServiceServer(grpcServer, &server{})

  // If an error occurs when serving up the TCP connection, log the error.
  if err := grpcServer.Serve(lis); err != nil {
    fmt.Println("failed to serve: %s", err)
  }
}
