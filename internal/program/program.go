package program

import (
	"github.com/spf13/viper"
	"modules/internal/logger"
	"modules/internal/samlidp"
	"net/http"
)

type Program struct {
	exit  chan struct{}
	cfg   *viper.Viper
	srv   *http.Server
	store *samlidp.SqliteStore
}

func New(cfg *viper.Viper, logger *logger.LoggerEx, workDir string) *Program {
	p := new(Program)
}
