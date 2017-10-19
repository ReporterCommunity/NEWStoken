package main

import (
	"fmt"
	"log"
	"math/big"

	"os"

	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

func exit(str string) {
	fmt.Println(str)
	log.Println(str)
	os.Exit(1)
}
func main() {
	viper.SetConfigName("config")         // name of config file (without extension)
	viper.AddConfigPath("/etc/icotest/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.icotest") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	testing := viper.GetBool("testing")
	//listing := viper.GetBool("listing")
	logName := viper.GetString("log")
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/" + logName,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	})
	blkNo = big.NewInt(0)
	if !areContractsLaunched() {
		exit("waiting for contract stability")
	}
	if !areContractsLinked() {
		exit("waiting for contract linkages")
	}
	if !areContractsSetup() {
		exit("still crossing the i's and dotting the t's")
	}
	sendEthereum(roleKey("banker"), customerServiceAddress(), big.NewInt(1000000000000000000), big.NewInt(21000))
	if testing {
		done := make(chan bool)
		go runscript(done)
		<-done
	}

}
