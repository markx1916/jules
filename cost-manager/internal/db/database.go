package db
import ("cost-manager/config";"cost-manager/internal/model";"fmt";"gorm.io/driver/mysql";"gorm.io/gorm")
var DB *gorm.DB
func Init() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config.Database.User, config.Cofig.Database.Password, config.Config.Database.Host, config.Config.Database.Port, config.Config.Database.DBName)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); if err != nil { return err }
	return DB.AutoMigrate(&model.CostData{}, &model.Budget{}, &model.Alert{})
}
