package config

import (
    "os"
    "encoding/json"
)

type Config struct {
    Database struct {
        Host        string  `json:"host"`
        Password    string  `json:"password"`
    }   `json:"database"`
}

func LoadConfig(fileName string) (Config, error) {
    var config Config
    configFile, err := os.Open(fileName)
    defer configFile.Close()
    if err != nil {
        return config, err
    }
    jsonParser := json.NewDecoder(configFile)
    err = jsonParser.Decode(&config)
    return config, err
}
