package templa

import (
	"encoding/json"
	"os"
)

type Config struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Values      map[string]ValueConfig `json:"values"`
}

type ValueConfig struct {
	Required bool   `json:"required"`
	Default  string `json:"default"`
}

const DEFAULT_CONFIG_FILE = "templa.json"

func LoadConfig(name string) (Config, error) {
	config := Config{}

	b, err := os.ReadFile(name)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(b, &config)

	return config, err
}
