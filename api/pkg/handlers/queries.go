package handlers

import (
	"fmt"
	"net/http"
	"time"

	"goVueForum/api/pkg/postgres"

	"database/sql"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Post represents ...
type Post struct {
	Author      Author `json:"author"`
	Body        string `json:"body"`
	CreatedAt   string `json:"createdAt"`
	ID          uint32 `json:"id"`
	LastUpdated string `json:"lastUpdatedAt"`
	Title       string `json:"title"`
}

// Author represents...
type Author struct {
	ID        uint32 `json:"id"`
	LastLogin uint32 `json:"lastLogin,omitempty"`
	Username  string `json:"username"`
}

// User represents...
type User struct {
	AvatarURL sql.NullString `json:"avatarUrl"`
	Email     string         `json:"email"`
	FirstName string         `json:"firstName"`
	ID        uint32         `json:"id"`
	LastLogin sql.NullString `json:"lastLogin"`
	LastName  string         `json:"lastName"`
	Location  sql.NullString `json:"location"`
	PostCount string         `json:"postCount"`
	Username  string         `json:"username"`
}

type forumCategory struct {
	ID    uint64
	Title string
}

// Thread represents...
type Thread struct {
	Author      Author `json:"author"`
	CreatedAt   string `json:"createdAt"`
	ForumID     uint32 `json:"forumID"`
	ID          uint32 `json:"id"`
	LastReplyAt string `json:"lastReplyAt"`
	Posts       []Post `json:"posts"`
	Title       string `json:"title"`
}

// GetHealth serves as a simple server health check
func GetHealth(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

// GetPosts ...
func GetPosts(c *postgres.Conn) (*[]Post, error) {
	rows, err := c.DB.Query("SELECT * FROM posts_v")

	if err != nil {
		return nil, fmt.Errorf("failed to perform query: %s", err.Error())
	}

	defer rows.Close()

	posts := &[]Post{}
	for rows.Next() {
		p := &Post{}

		err = rows.Scan(&p.ID, &p.Author.Username, &p.Author.ID, &p.Title, &p.CreatedAt, &p.LastUpdated, &p.Body)
		*posts = append(*posts, *p)
	}

	err = rows.Err()

	if err != nil {
		return nil, fmt.Errorf("failed to scan rows to destination: %s", err.Error())
	}

	return posts, nil
}

func fetchUser(ID string, conn *postgres.Conn) (*User, error) {
	user := &User{}
	fmt.Println("ID: ", ID)

	row := conn.DB.QueryRow(`SELECT
                            avatar_url,
                            email,
                            first_name,
                            user_account_id AS id,
                            last_login,
                            last_name,
                            location,
                            post_count,
                            username
                            FROM user_account
                            WHERE user_account_id = $1 
                          `, ID)

	err := row.Scan(&user.AvatarURL, &user.Email, &user.FirstName, &user.ID, &user.LastLogin, &user.LastName, &user.Location, &user.PostCount, &user.Username)

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return user, nil

}

// fetchThread ...
func fetchThread(ID string, conn *postgres.Conn) (*Thread, error) {
	thread := &Thread{}

	c1 := make(chan Thread, 1)
	go func(c chan Thread) {
		t := &Thread{}

		row := &sql.Row{}

		// fetch thread by thread ID
		row = conn.DB.QueryRow(`
        SELECT
          t.author_id,
          u.username,
          t.created_at,
          t.forum_id,
          t.thread_id,
          t.last_reply_at,
          t.title
        FROM thread t
        LEFT JOIN post p
          ON t.thread_id = p.thread_id
        LEFT JOIN user_account u
          ON t.author_id = u.user_account_id
        WHERE t.thread_id = $1;
        `, ID)
		row.Scan(&t.Author.ID, &t.Author.Username, &t.CreatedAt, &t.ForumID, &t.ID, &t.LastReplyAt, &t.Title)

		c <- *t
	}(c1)

	*thread = <-c1

	c2 := make(chan []Post)

	// fetch posts by thread ID
	go func(c chan []Post) {
		posts := &[]Post{}
		rows := &sql.Rows{}

		rows, _ = conn.DB.Query(`
        SELECT
          u.user_account_id,
          u.username,
          pb.body,
          p.created_at,
          p.post_id,
          p.last_updated_at,
          p.title
        FROM post p
        LEFT JOIN post_body pb
          ON pb.post_body_id = p.post_body_id
        LEFT JOIN user_account u
          ON p.author_id = u.user_account_id
        WHERE p.thread_id = $1;
        `, ID)
		defer rows.Close()

		for rows.Next() {
			p := Post{}
			err := rows.Scan(&p.Author.ID, &p.Author.Username, &p.Body, &p.CreatedAt, &p.ID, &p.LastUpdated, &p.Title)

			*posts = append(*posts, p)
			if err != nil {
				log.Fatalln(err)
			}
		}

		c2 <- *posts
	}(c2)

	(*thread).Posts = <-c2
	// row.Scan(&thread.Author.ID, &thread.Author.Username, &thread.CreatedAt, &thread.ForumID, &thread.ID, &thread.LastReplyAt, &thread.Title)

	// if err != nil {
	// 	return nil, fmt.Errorf("failed to perform query: %s", err.Error())
	// }

	// for rows.Next() {
	// 	p := Post{}
	// 	err := rows.Scan(&p.Author.ID, &p.Author.Username, &p.Body, &p.CreatedAt, &p.ID, &p.LastUpdated, &p.Title)

	// 	// err := rows.Scan(&thread.Author.ID, &thread.Author.Username, &thread.CreatedAt, &thread.ForumID, &thread.ID, &thread.LastReplyAt, &thread.Title, &p.Author.ID, &p.Author.Username, &p.Body, &p.CreatedAt, &p.ID, &p.LastUpdated, &p.Title)

	// 	thread.Posts = append(thread.Posts, p)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }

	// if err != nil {
	// 	return nil, fmt.Errorf("failed to scan rows to destination: %s", err.Error())
	// }

	return thread, nil
}

// GetThread ...
func GetThread(w http.ResponseWriter, req *http.Request) {
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

	vars := mux.Vars(req)

	thread, err := fetchThread(vars["id"], conn)

	if err != nil {
		log.Errorf("error getting posts: %v", err.Error())
	}

	res = JSONresponse{
		Result: Result{
			Data: map[string]interface{}{
				"thread": thread,
			},
		},
	}

	writeJSONresponse(w, req.Header, start, nil, res)
}

// GetUser ...
func GetUser(w http.ResponseWriter, req *http.Request) {
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

	vars := mux.Vars(req)

	user, err := fetchUser(vars["id"], conn)

	if err != nil {
		log.Errorf("error fetch user: %v", err.Error())
	}

	res = JSONresponse{
		Result: Result{
			Data: map[string]interface{}{
				"user": user,
			},
		},
	}

	writeJSONresponse(w, req.Header, start, nil, res)
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
