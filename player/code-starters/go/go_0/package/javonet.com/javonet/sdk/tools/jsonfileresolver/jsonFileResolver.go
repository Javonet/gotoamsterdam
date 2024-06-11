package jsonfileresolver

import (
	"encoding/json"
	"errors"
	"os"
)

type JsonFileResolver struct {
	Path       string
	JsonObject map[string]interface{}
	Runtimes   map[string]interface{}
}

func NewJsonFileResolver(path string) (*JsonFileResolver, error) {
	resolver := &JsonFileResolver{Path: path}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, errors.New("Configuration file " + path + " not found. Please check your configuration file.")
	}

	jsonText, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonText, &resolver.JsonObject)
	if err != nil {
		return nil, err
	}

	return resolver, nil
}

func (resolver *JsonFileResolver) GetLicenseKey() (string, error) {
	if licenseKey, ok := resolver.JsonObject["licenseKey"].(string); ok {
		return licenseKey, nil
	}
	return "", errors.New("License key not found in configuration file. Please check your configuration file.")
}

func (resolver *JsonFileResolver) GetRuntimes() {
	resolver.Runtimes = resolver.JsonObject["runtimes"].(map[string]interface{})
}

func (resolver *JsonFileResolver) GetRuntime(runtimeName string, configName string) (map[string]interface{}, error) {
	runtimes := resolver.JsonObject["runtimes"].(map[string]interface{})
	for key, value := range runtimes {
		if key != runtimeName {
			continue
		}
		switch v := value.(type) {
		case []interface{}:
			for _, runtime := range v {
				runtimeMap := runtime.(map[string]interface{})
				if runtimeMap["name"] == configName {
					return runtimeMap, nil
				}
			}
		case map[string]interface{}:
			if v["name"] == configName {
				return v, nil
			}
		}
	}
	return nil, errors.New("Runtime config " + configName + " not found in configuration file for runtime " + runtimeName + ". Please check your configuration file.")
}

func (resolver *JsonFileResolver) GetChannel(runtimeName string, configName string) (map[string]interface{}, error) {
	runtime, err := resolver.GetRuntime(runtimeName, configName)
	if err != nil {
		return nil, err
	}
	if channel, ok := runtime["channel"].(map[string]interface{}); ok {
		return channel, nil
	}
	return nil, errors.New("Channel key not found in configuration file for config " + configName + ". Please check your configuration file.")
}

func (resolver *JsonFileResolver) GetChannelType(runtimeName string, configName string) (string, error) {
	channel, err := resolver.GetChannel(runtimeName, configName)
	if err != nil {
		return "", err
	}
	if channelType, ok := channel["type"].(string); ok {
		return channelType, nil
	}
	return "", errors.New("Channel type not found in configuration file for config " + configName + ". Please check your configuration file.")
}

func (resolver *JsonFileResolver) GetChannelHost(runtimeName string, configName string) (string, error) {
	channel, err := resolver.GetChannel(runtimeName, configName)
	if err != nil {
		return "", err
	}
	if channelHost, ok := channel["host"].(string); ok {
		return channelHost, nil
	}
	return "", errors.New("Channel host not found in configuration file for config " + configName + ". Please check your configuration file.")
}

func (resolver *JsonFileResolver) GetChannelPort(runtimeName string, configName string) (int, error) {
	channel, err := resolver.GetChannel(runtimeName, configName)
	if err != nil {
		return 0, err
	}
	if channelPort, ok := channel["port"].(float64); ok {
		return int(channelPort), nil
	}
	return 0, errors.New("Channel port not found in configuration file for config " + configName + ". Please check your configuration file.")
}
