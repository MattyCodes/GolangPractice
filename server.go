package main

import (
  "fmt"
  // "log"
  // "net"
  // "google.golang.org/grpc"
  pb "github.com/MattyCodes/GolangPractice/gen"
  "github.com/MattyCodes/GolangPractice/show_recommendation"
)

func main() {
  fmt.Println("Server running...")

  request := &pb.ShowRecommendationRequest{ Name: "Matty" }
  result := show_recommendation.RecommendShow(*request)

  fmt.Println(result)

  // s := &pb.ShowRecommendation{Name: "Matty", Show: "Peaky Blinders"}
  //
  // fmt.Println(s)
  // lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
  // if err != nil {
  //         log.Fatalf("failed to listen: %v", err)
  // }
  //
  // s := shoopy.Server{}
  //
  // grpcServer := grpc.NewServer()
  //
  // shoopy.RegisterShoopyServiceServer(grpcServer, &s)
  //
  // if err := grpcServer.Serve(lis); err != nil {
  //         log.Fatalf("failed to serve: %s", err)
  // }
}
