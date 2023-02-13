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

	cfg.SetDefault("service.name", "passport")                                       // Имя службы
	cfg.SetDefault("service.displayName", "Passport")                                // Отображаемое имя службы
	cfg.SetDefault("service.description", "Authorization service")                   // Описание службы
	cfg.SetDefault("server.host", "localhost")                                       // Хост сервера
	cfg.SetDefault("server.port", 8000)                                              // Порт сервера
	cfg.SetDefault("writeLog", false)                                                // Вести log-файл?
	cfg.SetDefault("pathes.certFile", "/etc/passport/server.crt")                    // Путь до файла сертификата
	cfg.SetDefault("pathes.keyFile", "/etc/passport/server.key")                     // Путь до ключа
	cfg.SetDefault("pathes.logFile", filepath.Join(execDir, "logs", "passport.log")) // Путь до log-файла

	cfg.SetConfigName("passport")
	cfg.SetConfigType("yaml")

	cfg.AddConfigPath(filepath.Join(os.Getenv("PROGRAMDATA"), "Passport"))
	cfg.AddConfigPath(filepath.Join(execDir, "configs"))

	if err := cfg.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}
