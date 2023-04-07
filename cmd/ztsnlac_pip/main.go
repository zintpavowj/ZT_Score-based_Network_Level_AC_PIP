package main

import (
	"flag"
	"log"

	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/apiserver"
)

var (
	confFilePath string
	useDBCache   bool
)

// The function makes some initialization before the main routine
func init() {

	// Definition of the command line arguments
	flag.StringVar(&confFilePath, "c", "./config/config.yml", "path to config file")
	flag.BoolVar(&useDBCache, "with-cache", false, "is given, enables caching all attributes from the DB")
}

// The main function
func main() {

	// Process the given command line arguments
	flag.Parse()

	// Create a default config und update its fields with the values from config file
	config := apiserver.NewConfig()
	if err := apiserver.UpdateConfigFromFile(confFilePath, config); err != nil {
		log.Fatal(err)
	}

	// Update the config value with the value from the program arguments
	config.UseDBCache = useDBCache

	// Check the config for necessary field to be initialized
	if err := apiserver.CheckConfig(config); err != nil {
		log.Fatal(err)
	}

	// Create and start the API server
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
