package dbconfig

import (
	"database/sql"
	"fmt"
	"github.com/gomodule/redigo/redis"
	_ "github.com/lib/pq"
	"os"
	"strings"
)

var (
	host     = strings.TrimSpace(os.Getenv("DATABASE_HOST_SWAPI")) //"localhost"
	port     = 5432
	user     = strings.TrimSpace(os.Getenv("DATABASE_USER_SWAPI")) //"postgres"
	password = strings.TrimSpace(os.Getenv("DATABASE_PWD_SWAPI"))  //"Welcome2000!"
	dbname   = strings.TrimSpace(os.Getenv("DATABASE_NAME_SWAPI")) //"bushadb"
)

// DB set up
func DBConn() (myDBConn *sql.DB, myErr error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return db, err
	}

	err = db.Ping()

	return db, err
}

func RedisConn() (redis.Conn, error) {
	c, err := redis.DialURL(os.Getenv("REDIS_URL"), redis.DialTLSSkipVerify(true))
	return c, err
}
