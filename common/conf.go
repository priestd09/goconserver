package common

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

const (
	PERIODIC_INTERVAL      = time.Minute
	PIPELINE_SEND_INTERVAL = 300 * time.Millisecond
	RFC8601_SECOND         = "2006-01-02T15:04:05Z0700"
	RFC8601_MILLISECOND    = "2006-01-02T15:04:05Z0700.000"
	RFC8601_MICROSECOND    = "2006-01-02T15:04:05Z0700.000000"
	RFC8601_NANOSECOND     = "2006-01-02T15:04:05Z0700.000000000"
	CLIENT_CONGO_TYPE      = iota
	CLIENT_XCAT_TYPE
)

var (
	serverConfig *ServerConfig
	clientConfig *ClientConfig
	CONF_FILE    string

	CLIENT_TYPE_MAP = map[string]int{
		"congo": CLIENT_CONGO_TYPE,
		"xcat":  CLIENT_XCAT_TYPE,
	}
	FALSE_CONFIG = map[string]bool{
		"False": false,
		"FALSE": false,
		"false": false,
		"NO":    false,
		"no":    false,
		"No":    false,
		"n":     false,
		"f":     false,
	}
	TRUE_CONFIG = map[string]bool{
		"TRUE": true,
		"True": true,
		"true": true,
		"t":    true,
		"YES":  true,
		"Yes":  true,
		"yes":  true,
		"y":    true,
	}
	PRECISION_FORMAT = map[string]string{
		"second":      RFC8601_SECOND,
		"millisecond": RFC8601_MILLISECOND,
		"microsecond": RFC8601_MICROSECOND,
		"nanosecond":  RFC8601_NANOSECOND}
)

type FileCfg struct {
	Name   string `yaml:"name"`
	LogDir string `yaml:"logdir"`
}

type TCPCfg struct {
	Name          string `yaml:"name"`
	Host          string `yaml:"host"`
	Port          string `ymal:"port"`
	Timeout       int    `yaml:"timeout"`
	SSLKeyFile    string `yaml:"ssl_key_file"`
	SSLCertFile   string `yaml:"ssl_cert_file"`
	SSLCACertFile string `yaml:"ssl_ca_cert_file"`
	SSLInsecure   bool   `yaml:"ssl_insecure"`
}

type UDPCfg struct {
	Name    string `yaml:"name"`
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	Timeout int    `yaml:"timeout"`
}

type LoggerCfg struct {
	File []FileCfg `yaml:"file"`
	TCP  []TCPCfg  `yaml:"tcp"`
	UDP  []UDPCfg  `yaml:"udp"`
}

type ServerConfig struct {
	Global struct {
		Host          string `yaml:"host"`
		SSLKeyFile    string `yaml:"ssl_key_file"`
		SSLCertFile   string `yaml:"ssl_cert_file"`
		SSLCACertFile string `yaml:"ssl_ca_cert_file"`
		LogFile       string `yaml:"logfile"`
		LogLevel      string `yaml:"log_level"` // debug, info, warn, error, fatal, panic
		Worker        int    `yaml:"worker"`
		StorageType   string `yaml:"storage_type"`
	}
	API struct {
		Port        string `yaml:"port"`
		HttpTimeout int    `yaml:"http_timeout"`
	}
	Console struct {
		Port              string    `yaml:"port"`
		DataDir           string    `yaml:"datadir"`
		LogTimestamp      bool      `yaml:"log_timestamp"`
		TimePrecision     string    `yaml:"time_precision"`
		TimeFormat        string    `yaml:"-"`
		ReplayLines       int       `yaml:"replay_lines"`
		ClientTimeout     int       `yaml:"client_timeout"`
		TargetTimeout     int       `yaml:"target_timeout"`
		ReconnectInterval int       `yaml:"reconnect_interval"`
		RPCPort           string    `yaml:"rpcport"`
		Loggers           LoggerCfg `yaml:"logger"`
	}
	Etcd struct {
		DailTimeout     int    `yaml:"dail_timeout"`
		RequestTimeout  int    `yaml:"request_timeout"`
		Endpoints       string `yaml:"endpoints"`
		ServerHeartbeat int64  `yaml:"server_heartbeat"`
	}
}

func InitServerConfig(confFile string) (*ServerConfig, error) {
	serverConfig.Global.Host = "127.0.0.1"
	serverConfig.Global.LogFile = ""
	serverConfig.Global.Worker = 1
	serverConfig.Global.LogLevel = "info"
	serverConfig.Global.StorageType = "file"
	serverConfig.API.Port = "12429"
	serverConfig.API.HttpTimeout = 10
	serverConfig.Console.Port = "12430"
	serverConfig.Console.DataDir = "/var/lib/goconserver/"
	serverConfig.Console.LogTimestamp = true
	serverConfig.Console.TimePrecision = "microsecond"
	serverConfig.Console.TimeFormat = RFC8601_MICROSECOND
	serverConfig.Console.ReplayLines = 30
	serverConfig.Console.ClientTimeout = 30
	serverConfig.Console.TargetTimeout = 30
	serverConfig.Console.ReconnectInterval = 5
	serverConfig.Console.RPCPort = "12431" // only for async storage type

	serverConfig.Etcd.DailTimeout = 5
	serverConfig.Etcd.RequestTimeout = 2
	serverConfig.Etcd.Endpoints = "127.0.0.1:2379"
	serverConfig.Etcd.ServerHeartbeat = 5
	data, err := ioutil.ReadFile(confFile)
	if err != nil {
		return serverConfig, nil
	}
	err = yaml.Unmarshal(data, &serverConfig)
	if err != nil {
		return serverConfig, err
	}
	i := 0
	for i = 0; i < len(serverConfig.Console.Loggers.TCP); i++ {
		if serverConfig.Console.Loggers.TCP[i].Timeout == 0 {
			serverConfig.Console.Loggers.TCP[i].Timeout = 3
		}
	}
	for i = 0; i < len(serverConfig.Console.Loggers.UDP); i++ {
		if serverConfig.Console.Loggers.UDP[i].Timeout == 0 {
			serverConfig.Console.Loggers.UDP[i].Timeout = 3
		}
	}
	if format, ok := PRECISION_FORMAT[serverConfig.Console.TimePrecision]; ok {
		serverConfig.Console.TimeFormat = format
	} else {
		plog.Error("TimePrecision configuration is not valied, use microsecond by default")
	}
	CONF_FILE = confFile
	return serverConfig, nil
}

func GetServerConfig() *ServerConfig {
	return serverConfig
}

type ClientConfig struct {
	SSLKeyFile     string
	SSLCertFile    string
	SSLCACertFile  string
	HTTPUrl        string
	ConsolePort    string
	ConsoleTimeout int
	ServerHost     string
	ClientType     int
	Insecure       bool
}

func NewClientConfig() (*ClientConfig, error) {
	var err error
	var ok bool
	clientConfig = new(ClientConfig)
	clientConfig.HTTPUrl = "http://127.0.0.1:12429"
	clientConfig.ServerHost = "127.0.0.1"
	clientConfig.ConsolePort = "12430"
	clientConfig.ConsoleTimeout = 30
	clientConfig.Insecure = false
	if os.Getenv("CONGO_URL") != "" {
		clientConfig.HTTPUrl = os.Getenv("CONGO_URL")
	}
	if os.Getenv("CONGO_SERVER_HOST") != "" {
		clientConfig.ServerHost = os.Getenv("CONGO_SERVER_HOST")
	}
	if os.Getenv("CONGO_PORT") != "" {
		clientConfig.ConsolePort = os.Getenv("CONGO_PORT")
	}
	if os.Getenv("CONGO_CONSOLE_TIMEOUT") != "" {
		clientConfig.ConsoleTimeout, err = strconv.Atoi(os.Getenv("CONGO_CONSOLE_TIMEOUT"))
		if err != nil {
			return nil, err
		}
	}
	if os.Getenv("CONGO_SSL_KEY") != "" {
		clientConfig.SSLKeyFile = os.Getenv("CONGO_SSL_KEY")
	}
	if os.Getenv("CONGO_SSL_CERT") != "" {
		clientConfig.SSLCertFile = os.Getenv("CONGO_SSL_CERT")
	}
	if os.Getenv("CONGO_SSL_CA_CERT") != "" {
		clientConfig.SSLCACertFile = os.Getenv("CONGO_SSL_CA_CERT")
	}
	clientConfig.ClientType = CLIENT_CONGO_TYPE
	if os.Getenv("CONGO_CLIENT_TYPE") != "" {
		var clientType int
		temp := os.Getenv("CONGO_CLIENT_TYPE")
		if clientType, ok = CLIENT_TYPE_MAP[temp]; ok {
			clientConfig.ClientType = clientType
		}
	}
	if _, ok = TRUE_CONFIG[os.Getenv("CONGO_SSL_INSECURE")]; ok {
		clientConfig.Insecure = true
	}
	return clientConfig, nil
}

func GetClientConfig() *ClientConfig {
	return clientConfig
}
