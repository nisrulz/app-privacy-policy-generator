package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ThirdPartyService struct {
	Name    string            `yaml:"name" json:"name"`
	Enabled bool              `yaml:"enabled" json:"enabled"`
	Logo    string            `yaml:"logo" json:"logo"`
	Link    map[string]string `yaml:"link" json:"link,omitempty"`
}

func buildThirdPartyJS() error {
	if err := ensureDir("public/tmp"); err != nil {
		return fmt.Errorf("create tmp dir: %w", err)
	}

	services, err := loadThirdPartyServices()
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(services)
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}

	js := fmt.Sprintf("var thirdPartyServicesJsonArray = %s", jsonData)
	if err := os.WriteFile("public/tmp/thirdpartyservices.js", []byte(js), 0644); err != nil {
		return fmt.Errorf("write js: %w", err)
	}

	return nil
}

func loadThirdPartyServices() ([]ThirdPartyService, error) {
	data, err := os.ReadFile("src/includes/yaml/thirdpartyservices.yml")
	if err != nil {
		return nil, err
	}

	var services []ThirdPartyService
	if err := yaml.Unmarshal(data, &services); err != nil {
		return nil, err
	}

	return services, nil
}
