package main

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/kardianos/service"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"modules/internal/logger"
	"modules/internal/program"
	"modules/internal/tools"
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
	configsDir := filepath.Join(execDir, "configs")

	cfg := viper.New()

	cfg.SetDefault("service.name", "passport")                                       // Имя службы
	cfg.SetDefault("service.displayName", "Passport")                                // Отображаемое имя службы
	cfg.SetDefault("service.description", "Authorization service")                   // Описание службы
	cfg.SetDefault("server.host", "localhost")                                       // Хост сервера
	cfg.SetDefault("server.port", 8000)                                              // Порт сервера
	cfg.SetDefault("writeLog", false)                                                // Вести log-файл?
	cfg.SetDefault("pathes.certFile", filepath.Join(configsDir, "idp.cert"))         // Путь до файла сертификата
	cfg.SetDefault("pathes.keyFile", filepath.Join(configsDir, "idp.key"))           // Путь до ключа
	cfg.SetDefault("pathes.logFile", filepath.Join(execDir, "logs", "passport.log")) // Путь до log-файла

	cfg.SetConfigName("passport")
	cfg.SetConfigType("yaml")

	cfg.AddConfigPath(filepath.Join(os.Getenv("PROGRAMDATA"), "Passport"))
	cfg.AddConfigPath(configsDir)

	if err := cfg.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	logger := logger.New(&lumberjack.Logger{
		Filename:   cfg.GetString("pathes.logFile"),
		MaxSize:    5,  // Максимальный размер в мегабайтах
		MaxAge:     30, // Количество дней для хранения старых логов
		MaxBackups: 10,
		LocalTime:  true,
		Compress:   true, // Сжимать в gzip-архивы
	})

	logger.SetEnabled(cfg.GetBool("writeLog"))
	log.SetOutput(logger)

	prg := program.New(cfg, logger, execDir)

	// Создаем службу
	options := make(service.KeyValue)
	options["Restart"] = "on-success"
	options["SuccessExitStatus"] = "1 2 8 SIGKILL"

	svcConfig := &service.Config{
		Name:        cfg.GetString("service.name"),
		DisplayName: cfg.GetString("service.displayName"),
		Description: cfg.GetString("service.description"),
		Option:      options,
	}

	svc, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	errs := make(chan error, 5)

	svcLogger, err := svc.Logger(errs)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			if err := <-errs; err != nil {
				log.Print(err)
			}
		}
	}()

	if len(*svcFlag) != 0 {
		if !tools.Contains(service.ControlAction[:], *svcFlag, true) {
			fmt.Fprintf(os.Stdout, "Valid actions: %q\n", service.ControlAction)
		} else if err := service.Control(svc, *svcFlag); err != nil {
			fmt.Fprintln(os.Stdout, err)
		}
		return
	}

	log.Printf("Used config file \"%s\"\n", cfg.ConfigFileUsed())

	cfg.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
		logger.SetEnabled(cfg.GetBool("writeLog"))
	})
	cfg.WatchConfig()

	if err := svc.Run(); err != nil {
		svcLogger.Error(err)
	}
}
