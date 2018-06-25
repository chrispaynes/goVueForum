package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type forumCategory struct {
	ID    uint64
	Title string
}

// Conn ...
type Conn struct {
	host string
	url  string
	DB   *sql.DB
}

// NewConn ...
func NewConn(host string) *Conn {
	return &Conn{
		host: host,
		url:  "user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " host=" + host + " sslmode=disable",
	}
}

// Open ...
func (c *Conn) Open() error {
	db, err := sql.Open("postgres", c.url)

	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	c.DB = db

	return nil
}

// Close ...
func (c *Conn) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
}

// GetCategories ...
func (c *Conn) GetCategories() {
	rows, err := c.DB.Query("SELECT * from forum_category")
	defer rows.Close()
	if err != nil {
		log.Infof("failed to perform query: %s", err.Error())
	}

	categories := []forumCategory{}
	for rows.Next() {
		fc := &forumCategory{}

		err = rows.Scan(&fc.ID, &fc.Title)
		categories = append(categories, *fc)
	}

	err = rows.Err()

	if err != nil {
		log.Infof("failed to scan rows to destination: %s", err.Error())
	}

	log.Infof("rows response2 %+v\n", categories)

}
