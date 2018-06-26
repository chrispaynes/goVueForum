package handlers

import (
	"fmt"
	"net/http"
	"time"

	"goVueForum/api/pkg/postgres"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Post represents ...
type Post struct {
	ID          uint32 `json:"id"`
	Author      Author `json:"author"`
	Title       string `json:"title"`
	CreatedAt   string `json:"createdAt"`
	LastUpdated string `json:"lastUpdatedAt"`
	Body        string `json:"body"`
}

// Author represents...
type Author struct {
	LastLogin uint32 `json:"lastLogin"`
	ID        uint32 `json:"id"`
	Username  string `json:"username"`
}

type forumCategory struct {
	ID    uint64
	Title string
}

// GetHealth serves as a simple server health check
func GetHealth(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

// GetPosts ...
func GetPosts(c *postgres.Conn) (*[]Post, error) {
	rows, err := c.DB.Query("SELECT * FROM posts_v")
	defer rows.Close()

	if err != nil {
		return nil, fmt.Errorf("failed to perform query: %s", err.Error())
	}

	posts := &[]Post{}
	for rows.Next() {
		p := &Post{}

		err = rows.Scan(&p.ID, &p.Author.Username, &p.Title, &p.CreatedAt, &p.LastUpdated, &p.Body)
		*posts = append(*posts, *p)
	}

	err = rows.Err()

	if err != nil {
		return nil, fmt.Errorf("failed to scan rows to destination: %s", err.Error())
	}

	return posts, nil
}

// GetIndex ...
func GetIndex(w http.ResponseWriter, req *http.Request) {
	conn := postgres.NewConn(viper.GetString("dev.bridge_IP"))
	err := conn.Open()
	defer conn.Close()
	var res JSONresponse

	start := time.Now()

	if err != nil {
		res = JSONresponse{
			Error: err.Error(),
		}

		writeJSONresponse(w, req.Header, start, err, res)
		return
	}

	posts, err := GetPosts(conn)

	if err != nil {
		log.Errorf("error getting posts: %v", err.Error())
	}

	res = JSONresponse{
		Result: Result{
			Data: map[string]interface{}{
				"posts": posts,
			},
		},
	}

	writeJSONresponse(w, req.Header, start, nil, res)
}

// GetCategories ...
func GetCategories(c *postgres.Conn) {
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
}
