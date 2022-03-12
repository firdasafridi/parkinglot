package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const (
	devConfigPath = "files/etc"
	DEV           = "development"
	PROD          = "production"
	STAG          = "staging"
)

var (
	fmtSprintf   = fmt.Sprintf
	osGetenv     = os.Getenv
	filepathJoin = filepath.Join
	osOpen       = os.Open
	osGetwd      = os.Getwd
	envGet       = os.Getenv("ENV")
)

// Config is main configuration file. It containts server configuration
type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	NewRelic NewRelic `yaml:"new_relic"`
}

type NewRelic struct {
	AppName string `yaml:"app_name"`
	Secret  string `yaml:"secret"`
}

// Server manage handler HTTP and Open Parking lot server
type Server struct {
	Name string `json:"name" yaml:"name"`
	HTTP HTTP   `json:"http" yaml:"http"`
}

// HTTP contains address and http namagement
type HTTP struct {
	Address        string `yaml:"address"`
	WriteTimeout   int    `yaml:"write_timeout"`
	ReadTimeout    int    `yaml:"read_timeout"`
	MaxHeaderBytes int    `yaml:"max_header_bytes"`
	Enable         bool   `yaml:"enable"`
}

// Database used for manage database configuration
type Database struct {
	Testing           bool   `yaml:"testing"`
	DSN               string `yaml:"dsn"`
	MaxConnection     int    `yaml:"max_conns"`
	MaxIdleConnection int    `yaml:"max_idle_conns"`
	MaxRetry          int    `yaml:"max_retry"`
}

// New is create initial and read configuration file.
func New(repoName string) (*Config, error) {
	path := getConfigFile(repoName)
	f, err := osOpen(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// getConfigFile get  config file name
func getConfigFile(repoName string) string {
	var (
		envConfig = envGet
		isLocal   = osGetenv("ISLOCAL")
		filename  = fmt.Sprintf("%s.%s.yaml", repoName, envConfig)
	)
	if isLocal != "1" {
		// for non dev env, use config in /etc
		if envConfig != DEV {
			return fmtSprintf("/etc/%s/%s", repoName, filename)
		}
	}
	envConfig = "development"
	filename = fmt.Sprintf("%s.%s.yaml", repoName, envConfig)

	/* #nosec */
	dir, _ := osGetwd()

	// use local files in dev
	return filepathJoin(dir, devConfigPath, repoName, filename)
}
