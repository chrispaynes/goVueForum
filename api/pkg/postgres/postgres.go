package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var log = logrus.New()

type forumCategory struct {
	ID    uint64 `db:"forum_category_id"`
	Title string `db:"title"`
}

// Conn ...
type Conn struct {
	Host string
	URL  string
	DB   *sqlx.DB
}

// Role ...
type Role struct {
	username string
	password string
}

// NewRole ...
func NewRole(rolename string) *Role {
	viper.SetConfigName("app")
	viper.AddConfigPath("../../")

	return &Role{
		username: viper.GetString("dev.POSTGRES_" + rolename + "_USERNAME"),
		password: viper.GetString("dev.POSTGRES_" + rolename + "_PASSWORD"),
	}
}

func (r *Role) name() string {
	return r.username
}

func (r *Role) passwd() string {
	return r.password
}

// NewConn ...
func NewConn(host string, rolename string) *Conn {
	role := NewRole(rolename)

	return &Conn{
		Host: host,
		URL:  "user=" + role.name() + " password=" + role.passwd() + " dbname=" + "vueforum" + " host=" + host + " sslmode=disable" + " connect_timeout=5",
	}
}

// Open ...
func (c *Conn) Open() error {
	var err error

	c.DB, err = sqlx.Connect("postgres", c.URL)

	if err != nil {
		return fmt.Errorf("failed to open connection to postgres database %s", err.Error())
	}

	return nil
}

// Close ...
func (c *Conn) Close() error {
	if c.DB != nil {
		err := c.DB.Close()

		if err != nil {
			return fmt.Errorf("%s", err.Error())
		}
	}

	return nil
}

// GetCategories ...
func (c *Conn) GetCategories() {
	rows, err := c.DB.Query("SELECT * from titan.forum_category")
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
