package internal

import (
	"github.com/spf13/viper"
	"errors"
	"strings"
	"strconv"
	"fmt"
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
	IP   string
	Port uint16 // port is 1 - 65535
}

type Logger struct {
	Path string `mapstructure:"output"`
}

func SetConfigFile()  {
	v := viper.New()
	c := Configuration{
		Server: &Server{},
		Logger: &Logger{},
	}

	v.SetConfigName(configName)
	v.AddConfigPath(configPath)

	//string to net.IP
	//net.ParseIP(ip)

	CheckErr(v.ReadInConfig(), "couldn't load config")
	//fmt.Println(v.GetStringMapString("server"))
	CheckErr(v.Unmarshal(&c), "couldn't unmarshal config")
	CheckErr(c.ValidConfig(), "couldn't validation config")

}

func (c *Configuration)ValidConfig() error {
	if err := validServer(c.Server); err != nil {
		return err
	}

	return nil
}

func validServer(s *Server) error {
	err := errors.New("Server.IP and Server.Port value is not valid")
	s.IP, s.Port = SplitAddress(s.Addr)
	fmt.Println(s)
	if s.Addr == "" {
		return err
	}

	ip := strings.Split(s.IP, ".")

	for _, v := range ip {
		octet, _ := strconv.ParseInt(v, 10, 16)
		if octet < 0 || octet > 255 {
			return err
		}
	}

	return nil
}