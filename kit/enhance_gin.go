package kit

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/open-kingfisher/king-utils/common"
	"github.com/open-kingfisher/king-utils/common/log"
	"time"
)

func EnhanceGin(g *gin.Engine) *gin.Engine {
	// 使用zap处理gin自身日志
	g.Use(log.GinZap(log.NewZap()))
	g.Use(log.RecoveryWithZap(log.NewZap(), true))
	// 跨域访问
	g.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", common.HeaderSigning},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))
	return g
}
