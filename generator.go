package personalassistant

import (
  "database/sql"
  "fmt"
  
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
}
