package configurator

import (
    "os"
    "path"
)

type Operation int

const (
    Print Operation = iota
    Add
    Remove
)

func (o Operation) String() string{
    return [...]string{"print", "add", "remove"}[o]
}

func (o Operation) Index() int {
    return int(o)
}

type ConfiguratorConfig struct {
    Pwd string
    Config string
    Operation Operation
    Arguments []string
}

func getConfig(config string) (string, error) {
    if config == "" {
        configDir, err := os.UserConfigDir()
        if err != nil {
            return "", err
        }

        return path.Join(configDir, "configurator", "configurator.json"), nil
    }

    return config, nil
}

func getPwd(pwd string) (string, error) {
    if pwd == "" {
        wd, err := os.Getwd()
        return wd, err
    }

    return pwd, nil
}

func isOperationCommand(op string) bool {
    return op == "add" || op == "remove";
}

func getArguments(commands []string) []string {

    if (len(commands) > 0 && isOperationCommand(commands[0])) {
        return commands[1:]
    }

    return commands
}

func getOperation(commands []string) Operation {
    if len(commands) == 0 {
        return Print
    }
    switch (commands[0]) {
        case "add": return Add
        case "remove": return Remove
    }
    return Print
}

func NewConfiguratorConfig(opts *ConfiguratorOpts) (*ConfiguratorConfig, error) {
    pwd, err := getPwd(opts.Pwd)
    if err != nil {
        return nil, err
    }

    config, err := getConfig(opts.Config)
    if err != nil {
        return nil, err
    }

    return &ConfiguratorConfig {
        Operation: getOperation(opts.Arguments),
        Arguments: getArguments(opts.Arguments),
        Pwd: pwd,
        Config: config,
    }, nil
}


