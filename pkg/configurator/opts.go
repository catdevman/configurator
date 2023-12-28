package configurator

import (
    "github.com/hellflame/argparse"
)

type ConfiguratorOpts struct {
    Pwd string
    Config string
    Arguments []string
}

func GetOptions() (*ConfiguratorOpts, error) {
    parser := argparse.NewParser("configurator", "gets all the values", &argparse.ParserConfig{DisableDefaultShowHelp: true})
    args := parser.Strings("a", "args", &argparse.Option{
        Positional: true,
        Default: "",
        Required: false,
    })

    config := parser.String("c", "config", &argparse.Option{Required: false, Default: ""})
    pwd := parser.String("p", "pwd", &argparse.Option{Required: false, Default: ""})

    err := parser.Parse(nil)
    if err != nil {
        return nil, err
    }

    return &ConfiguratorOpts {
        Pwd: *pwd,
        Config: *config,
        Arguments: *args,
    }, nil
}
