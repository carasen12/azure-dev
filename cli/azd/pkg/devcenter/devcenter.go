package devcenter

import (
	"encoding/json"
	"fmt"
)

// Merges supplemental configuration into the base config only if the key/value doesn't already exist in the base config
// Example: If the base config is a fully configured object, then any supplemental configuration will be ignored
func MergeConfigs(configs ...*Config) *Config {
	if len(configs) == 0 {
		panic("no configs provided")
	}

	destConfig := configs[0]

	mergedConfig := &Config{
		Name:                  destConfig.Name,
		Catalog:               destConfig.Catalog,
		Project:               destConfig.Project,
		EnvironmentType:       destConfig.EnvironmentType,
		EnvironmentDefinition: destConfig.EnvironmentDefinition,
	}

	for _, config := range configs[1:] {
		if config == nil {
			continue
		}

		if config.Name != "" && mergedConfig.Name == "" {
			mergedConfig.Name = config.Name
		}

		if config.Catalog != "" && mergedConfig.Catalog == "" {
			mergedConfig.Catalog = config.Catalog
		}

		if config.Project != "" && mergedConfig.Project == "" {
			mergedConfig.Project = config.Project
		}

		if config.EnvironmentType != "" && mergedConfig.EnvironmentType == "" {
			mergedConfig.EnvironmentType = config.EnvironmentType
		}

		if config.EnvironmentDefinition != "" && mergedConfig.EnvironmentDefinition == "" {
			mergedConfig.EnvironmentDefinition = config.EnvironmentDefinition
		}

		if config.User != "" && mergedConfig.User == "" {
			mergedConfig.User = config.User
		}
	}

	return mergedConfig
}

// ParseConfig attempts to parse a partial JSON configuration into a devcenter configuration
func ParseConfig(partialConfig any) (*Config, error) {
	var config *Config

	jsonBytes, err := json.Marshal(partialConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal dev center configuration: %w", err)
	}

	if err := json.Unmarshal(jsonBytes, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal dev center configuration: %w", err)
	}

	return config, nil
}

// MarshalConfig attempts to marshal a devcenter configuration into a map[string]any
func MarshalConfig(config *Config) (map[string]any, error) {
	jsonBytes, err := json.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal dev center configuration: %w", err)
	}

	var configMap map[string]any
	if err := json.Unmarshal(jsonBytes, &configMap); err != nil {
		return nil, fmt.Errorf("failed to unmarshal dev center configuration: %w", err)
	}

	return configMap, nil
}
