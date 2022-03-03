// Executable file.
package main

// Import libraries.
import (
  "fmt"
  "net"
  "google.golang.org/grpc"
  "golang.org/x/net/context"
  pb "github.com/MattyCodes/GolangPractice/src/application"
)

type server struct {
  pb.UnimplementedShowRecommendationServiceServer
}

func (s *server) RecommendShow(ctx context.Context, recommendation_request *pb.ShowRecommendationRequest) (*pb.ShowRecommendation, error) {
  fmt.Println("Recommendation request received...")
  fmt.Println(recommendation_request)

  return &pb.ShowRecommendation{ Name: "Matty", Show: "Peaky Blinders" }, nil
}

func main() {
  fmt.Println("Server running...")

  lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))

  if err != nil {
    fmt.Println("failed to listen: %v", err)
  }

  grpcServer := grpc.NewServer()

  pb.RegisterShowRecommendationServiceServer(grpcServer, &server{})

  if err := grpcServer.Serve(lis); err != nil {
    fmt.Println("failed to serve: %s", err)
  }
}
