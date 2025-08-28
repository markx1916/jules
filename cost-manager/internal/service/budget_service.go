package service
import ("cost-manager/internal/db"; "cost-manager/internal/model"; "fmt")
func GetBudgets() ([]model.Budget, error) {
	if db.DB == nil { return nil, fmt.Errorf("db not initialized") }
	var budgets []model.Budget
	err := db.DB.Find(&budgets).Error
	return budgets, err
}
func CreateBudget(budget *model.Budget) error {
	if db.DB == nil { return fmt.Errorf("db not initialized") }
	return db.DB.Create(budget).Error
}
