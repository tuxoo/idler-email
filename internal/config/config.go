package config

import (
	"github.com/spf13/viper"
	. "github/eugene-krivtsov/idler-email/pkg/mail"
	"strings"
)

type (
	Config struct {
		Mongo MongoConfig
		Mail  SenderConfig
	}

	MongoConfig struct {
		Host     string
		Port     string
		User     string
		Password string
		DB       string `mapstructure:"db"`
	}
)

func Init(path string) (*Config, error) {
	viper.AutomaticEnv()

	if err := parseConfigFile(path); err != nil {
		return nil, err
	}

	if err := parseEnv(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshalConfig(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func parseConfigFile(filepath string) error {
	path := strings.Split(filepath, "/")

	viper.AddConfigPath(path[0]) // folder
	viper.SetConfigName(path[1]) // config file name

	return viper.ReadInConfig()
}

func parseEnv() error {
	if err := parseMongoEnv(); err != nil {
		return err
	}

	if err := parseMailEnv(); err != nil {
		return err
	}

	return nil
}

func parseMongoEnv() error {
	if err := viper.BindEnv("mongo.host", "MONGO_HOST"); err != nil {
		return err
	}

	if err := viper.BindEnv("mongo.port", "MONGO_PORT"); err != nil {
		return err
	}

	if err := viper.BindEnv("mongo.db", "MONGO_DB"); err != nil {
		return err
	}

	if err := viper.BindEnv("mongo.user", "MONGO_INITDB_ROOT_USERNAME"); err != nil {
		return err
	}

	return viper.BindEnv("mongo.password", "MONGO_INITDB_ROOT_PASSWORD")
}

func parseMailEnv() error {
	if err := viper.BindEnv("mail.server", "MAIL_SERVER_NAME"); err != nil {
		return err
	}

	if err := viper.BindEnv("mail.user", "MAIL_USERNAME"); err != nil {
		return err
	}

	if err := viper.BindEnv("mail.password", "MAIL_PASSWORD"); err != nil {
		return err
	}

	if err := viper.BindEnv("mail.sender.name", "MAIL_SENDER_NAME"); err != nil {
		return err
	}

	if err := viper.BindEnv("mail.sender.address", "MAIL_SENDER_ADDRESS"); err != nil {
		return err
	}

	return nil
}

func unmarshalConfig(cfg *Config) error {
	if err := viper.UnmarshalKey("mongo", &cfg.Mongo); err != nil {
		return err
	}

	return nil
}

func setFromEnv(cfg *Config) {
	cfg.Mail.ServerName = viper.GetString("mail.server")
	cfg.Mail.Username = viper.GetString("mail.user")
	cfg.Mail.Password = viper.GetString("mail.password")
	cfg.Mail.SenderName = viper.GetString("mail.sender.name")
	cfg.Mail.SenderAddress = viper.GetString("mail.sender.address")

	cfg.Mongo.Host = viper.GetString("mongo.host")
	cfg.Mongo.Port = viper.GetString("mongo.port")
	cfg.Mongo.User = viper.GetString("mongo.user")
	cfg.Mongo.Password = viper.GetString("mongo.password")
	cfg.Mongo.DB = viper.GetString("mongo.db")
}