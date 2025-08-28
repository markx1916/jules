package service
import ("cost-manager/internal/db"; "cost-manager/internal/model"; "fmt")
type AggregatedCost struct { GroupKey string `json:"group_key"`; TotalCost float64 `json:"total_cost"` }
func GetAggregatedCosts(groupBy string) ([]AggregatedCost, error) {
	if db.DB == nil { return nil, fmt.Errorf("db not initialized") }
	allowed := map[string]string{"service": "service_name", "cloud": "cloud_provider", "account": "account_id"}
	dbColumn, ok := allowed[groupBy]; if !ok { return nil, fmt.Errorf("invalid groupBy: %s", groupBy) }
	var results []AggregatedCost
	err := db.DB.Model(&model.CostData{}).Select(dbColumn + " as group_key, sum(cost_actual) as total_cost").Group(dbColumn).Find(&results).Error
	return results, err
}
