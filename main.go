package personalassistant

import (
  "database/sql"
  "fmt"
  "log"
  "net/http"
  "os"
  
  "github.com/gin-gonic/gin"
  _ "github.com/heroku/x/hmetrics/onload"  
  _ "github.com/lib/pq"
)

const (
  host		= os.Getenv("PG_HOST")
  port		= os.Getenv("PG_PORT")
  user		= os.Getenv("PG_USER")
  password	= os.Getenv("PG_PASSWORD")
  dbname	= os.Getenv("PG_DBNAME")
)

func main() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

  port := os.Getenv("PORT")

  if port == "" {
    log.Fatal("$PORT must be set")
  }

  router := gin.New()
  router.Use(gin.Logger())
  router.LoadHTMLGlob("templates/*.tmpl.html")
  router.Static("/static", "static")

  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl.html", nil)
  })

  router.Run(":" + port)
}