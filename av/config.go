package av

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// Config contains the unmarshalled config.json
type Config struct {
	APIKey string `json:"api_key"`
	URL    string `json:"url"`
}

// LoadConfig unmarshals config.json into a Config
func LoadConfig() Config {
	filepath, _ := filepath.Abs("/Users/cbrit/go/src/github.com/cbrit/coral/av/config.json")
	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}
	var c Config
	json.Unmarshal(raw, &c)
	return c
}
