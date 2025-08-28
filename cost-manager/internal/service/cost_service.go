package service
import ("fmt")
type AggregatedCost struct {
	GroupKey string  `json:"group_key"`
	TotalCost float64 `json:"total_cost"`
}
func GetAggregatedCosts(groupBy string) ([]AggregatedCost, error) {
	fmt.Printf("Service layer: Aggregating costs by %s\n", groupBy)
	mockData := []AggregatedCost{
		{GroupKey: "AmazonEC2", TotalCost: 150.75},
		{GroupKey: "AmazonS3", TotalCost: 75.20},
	}
	if groupBy == "cloud" {
		mockData = []AggregatedCost{
			{GroupKey: "AWS", TotalCost: 225.95},
			{GroupKey: "GCP", TotalCost: 295.60},
		}
	}
	return mockData, nil
}
