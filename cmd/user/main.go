package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/jchukwum/ai-fitted-user/pkg/cfg"
	"github.com/jchukwum/ai-fitted-user/pkg/handlers"
	"github.com/jchukwum/ai-fitted-user/pkg/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := cfg.Init()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := gorm.Open(mysql.Open(cfg.Spec.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	cfg.Db = db
	repository.Init()

	r := gin.Default()
	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte(cfg.Spec.Secret))
	if err != nil {
		panic(err)
	}
	store.Options(sessions.Options{
		MaxAge: 60 * 60 * cfg.Spec.ExpirationInHours, // expire in a day
		Path:   "/"},
	) // expire in a day
	r.Use(sessions.Sessions("ai-fitted-session", store))
	r.Use(cors.Default())
	r.GET("api/"+cfg.Spec.APIVersion+"/user/profile", handlers.GetProfile)
	r.POST("api/"+cfg.Spec.APIVersion+"/user/profile", handlers.SaveProfile)
	r.Run(":" + cfg.Spec.APIPort)
}
