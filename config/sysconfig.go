package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
)

var (
	// The path is relative to where execute binary(ETH-Agent)
	sysConfigPath string

	// SysConf - config of everything in agent
	SysConf SysConfig

	// EthURL - url for Eth node
	EthURL string

	//
	ToeknVerifyURI string

	// AWSALBCookieExpires - Setting how many time will the redis cache the AWSALB cookie for user.
	AWSALBCookieExpires = 576000

	// JSONRPCMethodCookieExipres - Setting how many time will the redis cache the JSON-RPC method which is cached typed.
	JSONRPCMethodCookieExipres = 30
)

// SysConfig - system config
type SysConfig struct {
	EthProxy EthProxyConf `yaml:"eth_proxy"`
	Eth      EthConf      `yaml:"eth"`
	SSO      SSOConf      `yaml:"sso"`
	Redis    RedisConf    `yaml:"redis"`
	Mongo    MongoConf    `yaml:"mongo"`
}

// EthProxyConf - config of eth_proxy service
type EthProxyConf struct {
	Protocol   string `yaml:"protocol"`
	Domain     string `yaml:"domain"`
	Port       string `yaml:"port"`
	LogFile    string `yaml:"log_file"`
	LogLevel   int8   `yaml:"log_level"`
	ProductBin string `yaml:"product_bin"`
}

// EthConf - config of ethereum service
type EthConf struct {
	Domain string `yaml:"domain"`
	Port   string `yaml:"port"`
}

// SSOConf - config of SSO service
type SSOConf struct {
	Domain string `yaml:"domain"`
	Port   string `yaml:"port"`
}

// RedisConf - config of redis service
type RedisConf struct {
	Domain   string `yaml:"domain"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}

// MongoConf - config of MongoDB service
type MongoConf struct {
	Domain   string `yaml:"domain"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"agent_db"`
	Username string `yaml:"agent_user"`
	Password string `yaml:"agent_pwd"`
}

// Parse content in sysConfigPath
func init() {
	app := cli.NewApp()
	app.Name = "eth-agent"
	app.Usage = "eth proxy"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Value: "sysconfig.yaml",
			Usage: "Configuration file of system setting.",
		},
	}
	app.Version = "0.1.0"
	app.Action = parseParams
	app.Run(os.Args)
}

func parseParams(c *cli.Context) error {
	sysConfigPath = c.GlobalString("config")
	if sysConfigPath == "" {
		fmt.Println("Empty config file path is not allowed")
	}
	absSysConfigPath, _ := filepath.Abs(sysConfigPath)

	var content []byte
	var ioErr error
	if content, ioErr = ioutil.ReadFile(absSysConfigPath); ioErr != nil {
		logrus.Fatalf("read service config file error: %v", ioErr)
		return nil
	}

	if ymlErr := yaml.Unmarshal(content, &SysConf); ymlErr != nil {
		logrus.Fatalf("error while unmarshal from db config: %v", ymlErr)
		return nil
	}

	EthURL = "http://" + SysConf.Eth.Domain + ":" + SysConf.Eth.Port
	ToeknVerifyURI = "http://" + SysConf.SSO.Domain + ":" + SysConf.SSO.Port + "/token/verify"
	fmt.Println(EthURL)
	fmt.Println(ToeknVerifyURI)

	return nil
}

// ToLogLevel - convert from int to Level type
func ToLogLevel(lv int8) logrus.Level {
	switch lv {
	case 0:
		return logrus.PanicLevel
	case 1:
		return logrus.FatalLevel
	case 2:
		return logrus.ErrorLevel
	case 3:
		return logrus.WarnLevel
	case 4:
		return logrus.InfoLevel
	case 5:
		return logrus.DebugLevel
	}

	return logrus.InfoLevel
}
