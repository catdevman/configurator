package configurator_test

import (
    "testing"
    "github.com/catdevman/configurator/pkg/configurator"
)

func TestConfigPrint(t *testing.T) {
    opts := configurator.ConfiguratorOpts {
        Config: "",
        Pwd: "",
        Arguments: []string{},
    }

    config, err := configurator.NewConfiguratorConfig(&opts)

    if err != nil {
        t.Errorf("error returned from configurator config %v", err)
    }

    if config.Operation != configurator.Print {
        t.Errorf("operation expected was print but got %v", config.Operation)
    }
}

func TestConfigAdd(t *testing.T) {
    opts := configurator.ConfiguratorOpts {
        Config: "",
        Pwd: "",
        Arguments: []string{"add", "foo", "bar"},
    }

    config, err := configurator.NewConfiguratorConfig(&opts)

    if err != nil {
        t.Errorf("error returned from configurator config %v", err)
    }

    if config.Operation != configurator.Add {
        t.Errorf("operation expected was add but got %v", config.Operation)
    }

    if config.Arguments[0] != "foo" || config.Arguments[1] != "bar" {
        t.Errorf("expected arguments to equal {'foo', 'bar'} but got %+v", config.Arguments)
    }
}


func TestConfigRemove(t *testing.T) {
    opts := configurator.ConfiguratorOpts {
        Config: "",
        Pwd: "",
        Arguments: []string{"remove", "foo"},
    }

    config, err := configurator.NewConfiguratorConfig(&opts)

    if err != nil {
        t.Errorf("error returned from configurator config %v", err)
    }

    if config.Operation != configurator.Remove {
        t.Errorf("operation expected was add but got %v", config.Operation)
    }

    if config.Arguments[0] != "foo" {
        t.Errorf("expected arguments to equal 'foo' but got %+v", config.Arguments[0])
    }
}
