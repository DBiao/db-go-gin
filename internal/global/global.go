package global

import (
	config "db-go-gin/conf"
	"net/http"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CONFIG *config.Config
	VIPER  *viper.Viper
	LOG    *zap.Logger
	DB     *gorm.DB
	SERVER *http.Server
)
