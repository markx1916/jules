package tasks
import ("cost-manager/internal/db"; "cost-manager/internal/model"; "log"; "time"; "fmt")
func CheckBudgetsTask() {
	log.Println("Running real budget check task...")
	if db.DB == nil { log.Println("DB not initialized, skipping."); return }
	var budgets []model.Budget
	db.DB.Find(&budgets)
	for _, budget := range budgets {
		var totalCost float64
		firstOfMonth := time.Now().UTC().Truncate(24 * time.Hour).AddDate(0, 0, -time.Now().Day() + 1)
		q := db.DB.Model(&model.CostData{}).Select("sum(cost_actual)").Where("billing_date >= ?", firstOfMonth)
		switch budget.ScopeType {
		case "account": q = q.Where("account_id = ?", budget.ScopeValue)
		case "service": q = q.Where("service_name = ?", budget.ScopeValue)
		default: continue
		}
		if err := q.Row().Scan(&totalCost); err != nil { continue }
		if totalCost > budget.Amount {
			alert := model.Alert{ AlertType: "budget_overrun", Description: fmt.Sprintf("Budget '%s' exceeded.", budget.Name), Severity: "warning" }
			db.DB.Create(&alert)
			log.Printf("ALERT: Budget '%s' exceeded!", budget.Name)
		}
	}
}
