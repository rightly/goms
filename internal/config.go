package internal

import (
	"github.com/spf13/viper"
	"errors"
	"strings"
	"strconv"
	"os"
)

const (
	configName = "goms"
	configPath = "./"
)

type Configuration struct {
	*Server `mapstructure:"server"`
	*Logger `mapstructure:"log"`
}

type Server struct {
	Addr string	`mapstructure:"address"`
	Role string `mapstructure:"role"`
	IP   string
	Port uint16 // port is 1 - 65535
}

type Logger struct {
	Path string `mapstructure:"output"`
}

func SetConfigFile() *Configuration {
	v := viper.New()
	c := &Configuration{
		Server: &Server{},
		Logger: &Logger{},
	}

	v.SetConfigName(configName)
	v.AddConfigPath(configPath)

	//string to net.IP
	//net.ParseIP(ip)

	CheckErr(v.ReadInConfig(), "couldn't load config")
	CheckErr(v.Unmarshal(&c), "couldn't unmarshal config")
	if CheckErr(c.ValidConfig(), "couldn't validation config") {
		os.Exit(ConfigureError)
	}
	return c
}

func (c *Configuration)ValidConfig() error {
	if err := c.Server.valid(); err != nil {
		return err
	}

	return nil
}

func (s *Server)valid() error {
	// Address 가 없을 시 err
	s.IP, s.Port = SplitAddress(s.Addr)
	if s.Addr == "" {
		err := errors.New("[server] address configuration value is not valid")
		return err
	}

	// IP의 각 octet 이 0 - 255 가 아닐 시 err
	ip := strings.Split(s.IP, ".")
	for _, v := range ip {
		octet, _ := strconv.ParseInt(v, 10, 16)
		if octet < 0 || octet > 255 {

			err := errors.New("[server] address configuration value is not valid")
			return err
		}
	}

	// role 이 manager or collector 가 아닐 시 err
	if !(s.Role == "manager" || s.Role == "collector") {
		err := errors.New("[server] role configuration value is not valid")
		return err
	}

	return nil
}