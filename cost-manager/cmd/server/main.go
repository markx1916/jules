package main
import ("cost-manager/config"; "cost-manager/internal/api"; "cost-manager/internal/tasks"; "log")
func main() { config.Init(); tasks.InitScheduler(); router := api.InitRouter(); log.Printf("Server starting on port %s...", config.Config.Server.Port); router.Run(":" + config.Config.Server.Port) }
