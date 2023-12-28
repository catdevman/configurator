package configurator_test

import (
	"testing"

	"github.com/catdevman/configurator/pkg/configurator"
)

func getData() *configurator.Data {
    return &configurator.Data{
        Configurator: map[string]map[string]string{
            "/" : {
                "foo": "bar1",
                "fem": "is_great",
            },
            "/foo": {
                "middle": "stuck_here_with_you",
                "foo": "bar2",
            },
            "/foo/bar": {
                "foo": "bar3",
            },
        },
    }
}

func getConfigurator(pwd string, data *configurator.Data) *configurator.Configurator{
    return configurator.CreateConfigurator(
        &configurator.ConfiguratorConfig{
            Arguments: []string{},
            Operation: configurator.Print,
            Pwd: pwd,
            Config: "config in test",
        },
        data,
    )
}

func TestGetValue(t *testing.T){
    data := getData()
    conf := getConfigurator("/foo/bar", data)

    tests := map[string]struct{
        key string
        expect string
    }{
        "get value foo": {
            "foo",
            "bar3",
        },
        "get value middle": {
            "middle",
            "stuck_here_with_you",
        },
        "get value fem": {
            "fem",
            "is_great",
        },
    }


    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            t.Parallel()
            val, ok := conf.GetValue(test.key)
            if !ok {
                t.Error("was not able to get value")
            }

            if val != test.expect {
                t.Errorf("expected %v but got %v", test.expect, val)
            }
        })
    }

}
