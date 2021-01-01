package config

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"strings"
)
import _ "gopkg.in/yaml.v2"

var defaultFileMode os.FileMode = 0644

func getConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrapf(err, "error getting config dir location")
	}
	return path.Join(home, ".config", "workspaces"), nil
}

func getConfigFile() (string, error) {
	confDir, err := getConfigDir()
	if err != nil {
		return "", errors.Wrapf(err, "error getting config file location")
	}
	return path.Join(confDir, "config.yml"), nil
}

func getRepoCacheFile() (string, error) {
	confDir, err := getConfigDir()
	if err != nil {
		return "", errors.Wrapf(err, "error getting repo cache file location")
	}
	return path.Join(confDir, ".repos.yml"), nil
}

type Organization struct {
	url url.URL
}

func (o Organization) GetApiUrl() url.URL {
	return url.URL{
		Scheme: o.url.Scheme,
		Opaque: o.url.Opaque,
		User:   o.url.User,
		Host:   "api." + o.url.Host,
		Path:   "/",
	}
}

func (o Organization) GetOrgName() string {
	return strings.TrimPrefix(o.url.Path, "/")
}

func (o Organization) String() string {
	return o.url.String()
}

func (o *Organization) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	parsed, err := url.Parse(s)
	if err == nil {
		o.url = *parsed

	}
	return err
}

func (o Organization) MarshalYAML() (interface{}, error) {
	return o.url.String(), nil
}

type Workspace struct {
	Name         string       `yaml:"name"`
	Directory    string       `yaml:"directory"`
	Organization Organization `yaml:"organization"`
}

type Config struct {
	Workspaces []Workspace `yaml:"workspaces"`
}

type RepoCache map[Organization][]string

func emptyConfig() Config {
	return Config{Workspaces: []Workspace{}}
}

func ReadConfig() (Config, error) {
	var conf Config

	confFile, err := getConfigFile()
	if err != nil {
		return conf, err
	}

	err = readYamlFile(confFile, &conf)
	if os.IsNotExist(err) {
		return emptyConfig(), nil
	}

	return conf, err
}

func WriteConfig(config Config) error {
	confDir, err := getConfigDir()
	if err != nil {
		return err
	}
	confFile, err := getConfigFile()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(confDir, defaultFileMode); err != nil {
		return errors.Wrapf(err, "error creating config dir: %s", confDir)
	}
	return writeYamlFile(confFile, &config)
}

func ReadRepoCache() (RepoCache, error) {
	repos := RepoCache{}

	repoCacheFile, err := getRepoCacheFile()
	if err != nil {
		return nil, err
	}

	err = readYamlFile(repoCacheFile, &repos)
	if os.IsNotExist(err) {
		return nil, errors.Errorf("repo cache file does not exist: %s", repoCacheFile)
	}

	return repos, err
}

func WriteRepoCacheFile(repos RepoCache) error {
	confDir, err := getConfigDir()
	if err != nil {
		return nil
	}

	repoCacheFile, err := getRepoCacheFile()
	if err != nil {
		return nil
	}

	if err := os.MkdirAll(confDir, defaultFileMode); err != nil {
		return errors.Wrapf(err, "error creating config dir: %s", confDir)
	}
	return writeYamlFile(repoCacheFile, &repos)
}

func writeYamlFile(path string, out interface{}) error {
	yamlConfig, err := yaml.Marshal(out)
	if err != nil {
		return errors.Wrapf(err, "error serializing config for file: %s", path)
	}
	if err := ioutil.WriteFile(path, yamlConfig, defaultFileMode); err != nil {
		return errors.Wrapf(err, "error writing file: %s", path)
	}
	return nil
}

func readYamlFile(path string, out interface{}) error {
	yamlConfig, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.Wrapf(err, "error reading file: %s", path)
	}
	err = yaml.Unmarshal(yamlConfig, out)
	if err != nil {
		return errors.Wrapf(err, "error parsing file: %s", path)
	}
	return nil
}
