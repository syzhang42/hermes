package internal

import (
	"encoding/base64"
	"net/http"

	//"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ApiManager struct {
	g *gin.Engine
}

func NewApiManager() *ApiManager {
	return &ApiManager{}
}
func (a *ApiManager) Init() {
	a.g = gin.New()
	a.g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "api_key"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	v1 := a.g.Group("/v1")
	v1.Use(a.verifyAPIKey)
	v1.POST("/sign_in", signIn())
	v1.POST("/log_in", logIn())
}
func (a *ApiManager) verifyAPIKey(c *gin.Context) {
	apiKey := c.GetHeader("api_key")
	// 在这里实现验证 API Key 的逻辑
	if !isValidAPIKey(apiKey) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid API Key"})
		return
	}
	c.Next() // 允许通过
}

func isValidAPIKey(apiKey string) bool {
	today := time.Now().Format("20060102")
	expectedKey := generateExpectedAPIKey(today)

	return apiKey == expectedKey
}
func generateExpectedAPIKey(date string) string {
	return base64.StdEncoding.EncodeToString([]byte("online," + date + "lzb/fxy"))
}
func (a *ApiManager) Run(host string) {
	if a.g == nil {
		panic("gin.engine is nil,you must use run() before init()")
	}
	err := a.g.Run(host)
	if err != nil {
		panic(err)
	}
}
