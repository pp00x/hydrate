package config

import (
    "log"
    "time"

    "github.com/spf13/viper"
)

type Config struct {
    Server   ServerConfig
    Database DatabaseConfig
    JWT      JWTConfig
}

type ServerConfig struct {
    Port         string
    ReadTimeout  time.Duration
    WriteTimeout time.Duration
}

type DatabaseConfig struct {
    Host     string
    Port     string
    User     string
    DBName   string
    Password string
    SSLMode  string
    TimeZone string
}

type JWTConfig struct {
    SecretKey string
}

var AppConfig *Config

func InitConfig() {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        log.Printf("Error reading config file, %s", err)
    }

    AppConfig = &Config{
        Server: ServerConfig{
            Port:         viper.GetString("server.port"),
            ReadTimeout:  viper.GetDuration("server.read_timeout"),
            WriteTimeout: viper.GetDuration("server.write_timeout"),
        },
        Database: DatabaseConfig{
            Host:     viper.GetString("database.host"),
            Port:     viper.GetString("database.port"),
            User:     viper.GetString("database.user"),
            DBName:   viper.GetString("database.dbname"),
            Password: viper.GetString("database.password"),
            SSLMode:  viper.GetString("database.sslmode"),
            TimeZone: viper.GetString("database.timezone"),
        },
        JWT: JWTConfig{
            SecretKey: viper.GetString("jwt.secret_key"),
        },
    }
}

func init() {
    viper.SetDefault("server.port", "8080")
    viper.SetDefault("server.read_timeout", 5*time.Second)
    viper.SetDefault("server.write_timeout", 5*time.Second)

    viper.SetDefault("database.host", "localhost")
    viper.SetDefault("database.port", "5432")
    viper.SetDefault("database.user", "postgres")
    viper.SetDefault("database.dbname", "water_tracker")
    viper.SetDefault("database.password", "password")
    viper.SetDefault("database.sslmode", "disable")
    viper.SetDefault("database.timezone", "UTC")

    viper.SetDefault("jwt.secret_key", "your_secret_key")

    // Override with environment variables if they exist
    viper.AutomaticEnv()

    // Load the configuration
    InitConfig()
}