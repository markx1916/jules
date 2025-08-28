package api
import ("github.com/gin-gonic/gin"; "net/http"; "cost-manager/internal/service"; "cost-manager/internal/model")
func InitRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/costs", func(c *gin.Context){ costs, _ := service.GetAggregatedCosts(c.Query("groupBy")); c.JSON(http.StatusOK, costs) })
		v1.GET("/budgets", func(c *gin.Context){ budgets, _ := service.GetBudgets(); c.JSON(http.StatusOK, budgets) })
		v1.POST("/budgets", func(c *gin.Context){ var b model.Budget; c.BindJSON(&b); service.CreateBudget(&b); c.JSON(http.StatusOK, b) })
	}
	return r
}
