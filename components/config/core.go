package config

import (
	"fmt"
	"strings"

	logger "github.com/geomyidia/zylog/logger"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Configuration related constants
const (
	AppName         string = "raster"
	ConfigDir       string = "configs"
	ConfigFile      string = "app"
	ConfigType      string = "yaml"
	ConfigReadError string = "Fatal error config file"
)

func init() {
	viper.AddConfigPath(ConfigDir)
	viper.SetConfigName(ConfigFile)
	viper.SetConfigType(ConfigType)
	viper.SetEnvPrefix(AppName)

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.Set("Verbose", true)
	viper.AutomaticEnv()
	viper.AddConfigPath("/")

	err := viper.ReadInConfig()
	if err != nil {
		// log.Panic is not used here, since logging depends ...
		log.Panicf("%s: %s", ConfigReadError, err)
	}
}

// ClientConfig ...
type ClientConfig struct {
	Logging *logger.ZyLogOptions
}

// RedisConfig ...
type RedisConfig struct {
	Name     string
	Host     string
	Port     int
	Password string
	DBIndex  int
}

// RedixConfig ...
type RedixConfig struct {
	Name string
	Host string
	Port int
}

// DBConfig ...
type DBConfig struct {
	Type  string
	Redis *RedisConfig
	Redix *RedixConfig
}

// GRPCDConfig ...
type GRPCDConfig struct {
	Host string
	Port int
}

// Config ...
type Config struct {
	Client  *ClientConfig
	DB      *DBConfig
	GRPCD   *GRPCDConfig
	Logging *logger.ZyLogOptions
}

// New is a constructor that creates the full coniguration data structure
// for use by our application(s) and client(s) as an in-memory copy of the
// config data (saving from having to make repeated and somewhat expensive
// calls to the viper library).
//
// Note that Viper does provide both the AllSettings() and Unmarshall()
// functions, but these require that you have a struct defined that will be
// used to dump the Viper config data into. We've already got that set up, so
// there's no real benefit to switching.
//
// Furthermore, in our case, we're utilizing structs from other libraries to
// be used when setting those up (see how we initialize the logging component
// in ./components/logging.go, Setup).
func New() *Config {
	return &Config{
		Client: &ClientConfig{
			Logging: &logger.ZyLogOptions{
				Colored:      viper.GetBool("raster.client.logging.colored"),
				Level:        viper.GetString("raster.client.logging.level"),
				Output:       viper.GetString("raster.client.logging.output"),
				ReportCaller: viper.GetBool("raster.client.logging.report-caller"),
			},
		},
		DB: &DBConfig{
			Type: viper.GetString("raster.db.type"),
			Redix: &RedixConfig{
				Host: viper.GetString("raster.db.redix.host"),
				Port: viper.GetInt("raster.db.redix.port"),
			},
		},
		GRPCD: &GRPCDConfig{
			Host: viper.GetString("raster.grpc.host"),
			Port: viper.GetInt("raster.grpc.port"),
		},
		Logging: &logger.ZyLogOptions{
			Colored:      viper.GetBool("raster.logging.colored"),
			Level:        viper.GetString("raster.logging.level"),
			Output:       viper.GetString("raster.logging.output"),
			ReportCaller: viper.GetBool("raster.logging.report-caller"),
		},
	}
}

// GRPCConnectionString ...
func (c *Config) GRPCConnectionString() string {
	return fmt.Sprintf("%s:%d", c.GRPCD.Host, c.GRPCD.Port)
}

// RedisConnectionString ...
func (c *Config) RedisConnectionString() string {
	return fmt.Sprintf("%s:%d", c.DB.Redis.Host, c.DB.Redis.Port)
}

// RedixConnectionString ...
func (c *Config) RedixConnectionString() string {
	return fmt.Sprintf("%s:%d", c.DB.Redix.Host, c.DB.Redix.Port)
}
