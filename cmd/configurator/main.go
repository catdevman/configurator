package main

import (
	"fmt"
	"os"

	"github.com/catdevman/configurator/pkg/configurator"
)

func main(){
    opts, err := configurator.GetOptions()
    if err != nil {
        fmt.Println("this is bad")
        os.Exit(1)
    }


    conf, err := configurator.NewConfiguratorConfig(opts)
    if err != nil {
        fmt.Println("more bad")
        os.Exit(1)
    }
    urator := configurator.NewConfigurator(conf)
    switch conf.Operation {
    case configurator.Print: 
        if len(conf.Arguments) == 0 {
            all := urator.GetValueAll()
            for k, v := range all {
                fmt.Print(fmt.Sprintf("%v=%v\n", k, v))
            }
            os.Exit(0)
        }
        val, ok := urator.GetValue(conf.Arguments[0])
        if !ok {
            fmt.Println("no key:", conf.Arguments[0])
            os.Exit(1)
        }
        fmt.Println(val)
    case configurator.Add:
        if len(conf.Arguments) != 2 {
            fmt.Println("this is a problem")
            os.Exit(1)
        }
        key, value := conf.Arguments[0], conf.Arguments[1]
        urator.SetValue(key, value)
        urator.Save()
    case configurator.Remove:
        if len(conf.Arguments) != 1 {
            panic("this is a problem")
        }
        urator.RemoveValue(conf.Arguments[0])
        urator.Save()
        os.Exit(0)
    default:
        fmt.Println("I am not suppose to be here")
    }
}
