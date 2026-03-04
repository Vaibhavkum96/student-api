package config

import (
	"flag"
	"log"
	"os"
)

type HTTPServer struct {
	Addr string
}

//env-default:"production"

type Config struct {
	Env string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer `yaml:"http_server"`
}

func MustLoad() {
	var configPath string 
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		// Checking if some configuration is passed on commandline
		flags := flag.String("config","", "path to configuration file")
		flag.Parse()
		configPath = *flags
	}

	// If still the Config Path is Empty then throw the error 
	
	if configPath == "" {
		// We don't want the code to go through beyond this!
		log.Fatal("Config path Not Set!")
	}

	// If the Config is Not Empty then check if the file storage path is correct or not!
	if _, err := os.Stat(configPath); os.IsNotExist(err){
		log.Fatalf("Config File Does Not Exist: %s", configPath)
	}

	
}