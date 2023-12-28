package configurator

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

type Data struct {
    Configurator map[string]map[string]string `json:"configurator"`
}

type Configurator struct {
    config *ConfiguratorConfig
    data *Data
}

func CreateConfigurator(config *ConfiguratorConfig, data *Data) *Configurator{
    return &Configurator{
        config: config,
        data: data,
    }
}

func (c *Configurator) GetValue(key string) (string, bool){
    curr := c.config.Pwd
    prev := ""
    out := ""
    found := false
    for curr != prev {
        if dir, ok := c.data.Configurator[curr]; ok {
            if value, ok := dir[key]; ok {
                out = value
                found = true
                break
            }
        }
        prev = curr
        curr = path.Dir(curr)
    }
    return out, found
}

func (c *Configurator) GetValueAll() map[string]string{
    out := map[string]string{}
    paths := []string{}
    curr := c.config.Pwd
    prev := ""
    for curr != prev {
        paths = append(paths, curr)
        prev = curr
        curr = path.Dir(curr)
    }

    for i := len(paths) - 1; i >= 0; i-- {
        if dir, ok := c.data.Configurator[paths[i]]; ok {
            for k, v := range dir {
                out[k] = v
            }
        }
    }
    return out
}

func (c *Configurator) SetValue(key, value string) {
    pwd := c.config.Pwd
    if _, ok := c.data.Configurator[pwd]; !ok {
        c.data.Configurator[pwd] = map[string]string{}
    }

    c.data.Configurator[pwd][key] = value
}

func (c *Configurator) RemoveValue(key string) {
    pwd := c.config.Pwd
    if _, ok := c.data.Configurator[pwd]; ok {
        delete(c.data.Configurator[pwd], key)
    }
}

func (c *Configurator) Save(){
    configFile := c.config.Config
    if _, err := os.Stat(configFile); err != nil {
        //Create directory if it doesn't exit
        err = os.MkdirAll(filepath.Dir(configFile), os.ModePerm)
        if err != nil {
            panic(err)
        }
        //Create config file
        file, err := os.Create(configFile)
        if err != nil {
            panic(fmt.Sprintln("couldn't create config file", err))
        }
        defer file.Close()
    }
    contents, err := json.Marshal(&c.data)
    if err != nil {
        panic("")
    }
    err = os.WriteFile(configFile, contents, 0600)
    if err != nil {
        panic("couldn't write config file")
    }
}

func defaultConfigurator(config *ConfiguratorConfig) *Configurator{
    return &Configurator{
        config: config,
        data: &Data{
            Configurator: map[string]map[string]string{},
        },
    }
}

func NewConfigurator(config *ConfiguratorConfig) *Configurator {
    if _, err := os.Stat(config.Config); err == nil {
        contents, err := os.ReadFile(config.Config)
        if err != nil {
            return defaultConfigurator(config)
        }

        var data Data
        err = json.Unmarshal(contents, &data)
        if err != nil {
            return defaultConfigurator(config)
        }

        return &Configurator{
            config: config,
            data: &data,
        }
    }
    return defaultConfigurator(config)
}
