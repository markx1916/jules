package main
import ("cost-manager/config"; "cost-manager/internal/api"; "log")
func main() { config.Init(); router := api.InitRouter(); log.Printf("Server starting on port %s. (Not running)", config.Config.Server.Port) }
