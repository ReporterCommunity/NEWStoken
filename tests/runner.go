package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func runscript(ch chan bool) {
	scriptName := viper.GetString("script")
	if err := loadScript(scriptName); err != nil {
		fmt.Println(err)
	}
	runScript()
	ch <- true
}
