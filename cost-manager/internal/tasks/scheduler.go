package tasks
import ("github.com/robfig/cron/v3")
func InitScheduler() { c := cron.New(); c.AddFunc("0 0 * * *", CheckBudgetsTask); c.Start() }
