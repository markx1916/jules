package api
import ("github.com/gin-gonic/gin")
func InitRouter() *gin.Engine { r := gin.Default(); apiV1 := r.Group("/api/v1"); { apiV1.GET("/costs", GetCosts) }; return r }
func GetCosts(c *gin.Context) {}
