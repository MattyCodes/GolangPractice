package show_recommendation

import (
  "fmt"
  pb "github.com/MattyCodes/GolangPractice/gen"
)

func RecommendShow(recommendation_request pb.ShowRecommendationRequest) (pb.ShowRecommendation) {
  fmt.Println(recommendation_request)

  return *&pb.ShowRecommendation{ Name: "Matty", Show: "Peaky Blinders" }
}
