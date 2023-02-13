package main

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

func main() {
	svcFlag := flag.String("service", "", "Control the system service.")
	flag.Parse()

	execPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	execDir, _ := filepath.Split(execPath)
	execDir = filepath.Clean(execDir)

	cfg := viper.New()

	cfg.SetDefault("service.name", "passport")
	
}
