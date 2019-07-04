package main

import (
  "database/sql"
  "fmt"
  "log"
  "net/http"
  "os"
  
  // Google's UUID type
  "github.com/pborman/uuid"
  // A HTTP Web framework
  "github.com/gin-gonic/gin"
  _ "github.com/heroku/x/hmetrics/onload"  
  _ "github.com/lib/pq"
)

var (
  host		= os.Getenv("PG_HOST")
  port		= os.Getenv("PG_PORT")
  user		= os.Getenv("PG_USER")
  password	= os.Getenv("PG_PASSWORD")
  dbname	= os.Getenv("PG_DBNAME")
)

// Represents a house model in the database
type House struct {
  ID			uuid.UUID
  Address, City, State	string

  CreatedAt string `db:"created_at"`
  UpdatedAt string `db:"updated_at"`
}

func main() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

  port := os.Getenv("PORT")

  if port == "" {
    log.Fatal("$PORT must be set")
  }

  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  router := gin.New()
  router.Use(gin.Logger())
  router.LoadHTMLGlob("templates/*.tmpl.html")
  router.Static("/static", "static")

  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl.html", nil)
  })

  router.Run(":" + port)
}
