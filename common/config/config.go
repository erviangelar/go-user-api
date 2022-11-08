package config

import (
	"strconv"

	"github.com/lib/pq"
	"github.com/spf13/viper"
)

// Configurations wraps all the config variables required by the auth service
type Configurations struct {
	Port                       string
	ServerAddress              string
	DBConn                     string
	AccessTokenPrivateKeyPath  string
	AccessTokenPublicKeyPath   string
	RefreshTokenPrivateKeyPath string
	RefreshTokenPublicKeyPath  string
	JwtExpiration              int
	TokenHash                  string
	Smpt                       *SmptConfig
	Cache                      CacheConfig
	Log                        LogConfig
}

type SmptConfig struct {
	Sender     string
	Host       string
	Port       int
	Username   string
	Password   string
	Encryption string
}

type CacheConfig struct {
	Host     string
	Username string
	Password string
}

type LogConfig struct {
	LogFilePath string `env:"Log_FILE_PATH"`
	LogFileName string `env:"LOG_FILE_NAME"`
}

// NewConfigurations returns a new Configuration object
func LoadConfig() *Configurations {

	viper.SetConfigFile("./env/.env")
	viper.ReadInConfig()

	dbURL := viper.GetString("DATABASE_URL")
	conn, _ := pq.ParseURL(dbURL)

	configs := &Configurations{
		Port:                       viper.GetString("PORT"),
		ServerAddress:              viper.GetString("SERVER_ADDRESS"),
		DBConn:                     conn,
		JwtExpiration:              viper.GetInt("JWT_EXPIRATION"),
		AccessTokenPrivateKeyPath:  viper.GetString("ACCESS_TOKEN_PRIVATE_KEY_PATH"),
		AccessTokenPublicKeyPath:   viper.GetString("ACCESS_TOKEN_PUBLIC_KEY_PATH"),
		RefreshTokenPrivateKeyPath: viper.GetString("REFRESH_TOKEN_PRIVATE_KEY_PATH"),
		RefreshTokenPublicKeyPath:  viper.GetString("REFRESH_TOKEN_PUBLIC_KEY_PATH"),
		TokenHash:                  viper.GetString("TokenHash"),
	}

	emailPort, err := strconv.Atoi(viper.GetString("PORT"))
	if err != nil {
		panic(err)
	}
	configs.Smpt = &SmptConfig{
		Sender:     viper.GetString("SENDER"),
		Host:       viper.GetString("HOST"),
		Port:       emailPort,
		Username:   viper.GetString("USERNAME"),
		Password:   viper.GetString("PASSWORD"),
		Encryption: viper.GetString("ENCRYPTION"),
	}

	configs.Cache = CacheConfig{
		Host:     viper.GetString("HOST"),
		Username: viper.GetString("USERNAME"),
		Password: viper.GetString("PASSWORD"),
	}
	// reading heroku provided port to handle deployment with heroku
	port := viper.GetString("PORT")
	if port != "" {
		configs.ServerAddress = "0.0.0.0:" + port
	}
	return configs
}
