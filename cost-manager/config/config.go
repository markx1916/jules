package config
import ("fmt";"github.com/spf13/viper")
var Config struct { Server struct { Port string `mapstructure:"port"` } `mapstructure:"server"` }
func Init() error { viper.SetConfigName("config"); viper.SetConfigType("yaml"); viper.AddConfigPath("./config"); if err := viper.ReadInConfig(); err != nil { return err }; return viper.Unmarshal(&Config) }
