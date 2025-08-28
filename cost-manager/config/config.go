package config
import ("fmt";"github.com/spf13/viper")
var Config struct {
	Server struct { Port string `mapstructure:"port"` } `mapstructure:"server"`
	Database struct { Host, Port, User, Password, DBName string } `mapstructure:"database"`
}
func Init() error { viper.SetConfigName("config"); viper.SetConfigType("yaml"); viper.AddConfigPath("./config"); if err := viper.ReadInConfig(); err != nil { return err }; return viper.Unmarshal(&Config) }
