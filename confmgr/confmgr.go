package confmgr

import (
	"fmt"
	"io/ioutil"

	yaml "github.com/TR-SLimey/E/shim/yaml"
	sr "github.com/TR-SLimey/E/stringres"
)

// The rough layout of the E config file
type EConfigSkeleton struct {
	Matrix struct {
		AsId      string   `yaml:"asId"`
		Address   string   `yaml:"address"`
		BindAddrs []string `yaml:"bindAddrs"`
		BindPorts []int    `yaml:"bindPorts"`
		AsToken   string   `yaml:"asToken"`
		HsToken   string   `yaml:"hsToken"`
		Bot       struct {
			Username        string   `yaml:"username"`
			Displayname     string   `yaml:"displayname"`
			AvatarUrl       string   `yaml:"avatarUrl"`
			Sudoers         []string `yaml:"sudoers"`
			EnabledCommands []string `yaml:"enabledCommands"`
			NosudoCommands  []string `yaml:"nosudoCommands"`
		} `yaml:"bot"`
		Homeserver struct {
			Address      string `yaml:"address"`
			MxidSuffix   string `yaml:"mxidSuffix"`
			Provisioning struct {
				Path         string `yaml:"path"`
				SharedSecret string `yaml:"sharedSecret"`
			} `yaml:"provisioning"`
		} `yaml:"homeserver"`
		ManagedUsers struct {
			UsernameTemplate    string `yaml:"usernameTemplate"`
			DisplaynameTemplate string `yaml:"displaynameTemplate"`
		} `yaml:"managedUsers"`
	} `yaml:"matrix"`
	Esockets struct {
		ConfDir           string `yaml:"confDir"`
		FatalInitFailures bool   `yaml:"fatalInitFailures"`
	}
}

func GetEConfig(location string) (EConfigSkeleton, error) {
	var config EConfigSkeleton

	data, err := ioutil.ReadFile(location)
	if err != nil {
		return config, fmt.Errorf(sr.CONFIG_FILE_OPEN_ERR, location, err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf(sr.CONFIG_FILE_PARSE_ERR, location, err)
	}

	return config, nil
}

func GetEsocketConfig(location string, confVarPtr *struct{}) error {
	data, err := ioutil.ReadFile(location)
	if err != nil {
		return fmt.Errorf(sr.CONFIG_FILE_OPEN_ERR, location, err)
	}

	err = yaml.Unmarshal(data, confVarPtr)
	if err != nil {
		return fmt.Errorf(sr.CONFIG_FILE_PARSE_ERR, location, err)
	}

	return nil
}
