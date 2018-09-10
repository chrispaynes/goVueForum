package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"goVueForum/api/pkg/postgres"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx/reflectx"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Post represents ...
type Post struct {
	Author      `json:"author"`
	Body        string `json:"body" db:"body"`
	CreatedAt   string `json:"createdAt" db:"created_at"`
	ID          uint32 `json:"id" db:"post_id"`
	LastUpdated string `json:"lastUpdatedAt" db:"last_updated_at"`
	Title       string `json:"title" db:"title"`
}

// Author represents...
type Author struct {
	Username  string `json:"username" db:"author_username"`
	ID        uint32 `json:"id" db:"author_id"`
	LastLogin uint32 `json:"lastLogin,omitempty"`
}

// User represents...
type User struct {
	AvatarURL *string `json:"avatarUrl"`
	Email     string  `json:"email"`
	FirstName string  `json:"firstName"`
	ID        uint32  `json:"id"`
	LastLogin *string `json:"lastLogin"`
	LastName  string  `json:"lastName"`
	Location  *string `json:"location"`
	PostCount string  `json:"postCount"`
	Username  string  `json:"username"`
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
	stmt, err := c.DB.Preparex("select * from titan.posts_v")
	defer c.DB.Close()

	if err != nil {
		return nil, fmt.Errorf("failed to prepare select statement : %s", err.Error())
	}

	rows, err := stmt.Queryx()

	posts := &[]Post{}
	c.DB.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)

	for rows.Next() {
		var p Post

		err = rows.StructScan(&p)
		*posts = append(*posts, p)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to scan posts response to destination: %s", err.Error())
	}

	return posts, nil
}

func fetchUser(ID string, conn *postgres.Conn) (*User, error) {
	user := &User{}

	userID, err := strconv.Atoi(ID)

	if err != nil {
		return nil, fmt.Errorf("failed to parse ID %v to an integer: %s", ID, err)
	}

	row := conn.DB.QueryRow("SELECT * FROM get_user($1)", userID)

	err = row.Scan(&user.AvatarURL, &user.Email, &user.FirstName, &user.ID, &user.LastLogin, &user.LastName, &user.Location, &user.PostCount, &user.Username)

	if err != nil {
		return nil, fmt.Errorf("failed to scan user response to destination: %s", err)
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
        FROM titan.thread t
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

		rows, err := conn.DB.Query(`
        SELECT
          u.user_account_id,
          u.username,
          pb.body,
          p.created_at,
          p.post_id,
          p.last_updated_at,
          p.title
        FROM titan.post p
        LEFT JOIN post_body pb
          ON pb.post_body_id = p.post_body_id
        LEFT JOIN titan.users_v u
          ON p.author_id = u.user_account_id
        WHERE p.thread_id = $1;
		`, ID)

		if err != nil {
			log.Fatalln(err)
		}

		defer rows.Close()

		for rows.Next() {
			p := Post{}
			err := rows.Scan(&p)

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
	conn := postgres.NewConn(viper.GetString("dev.BRIDGE_IP"), "READONLY")
	err := conn.Open()
	var res JSONresponse
	start := time.Now()

	if err != nil {
		res = JSONresponse{
			Error: err.Error(),
		}

		writeJSONresponse(w, req.Header, start, err, res)
		return
	}

	defer conn.Close()

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
	conn := postgres.NewConn(viper.GetString("dev.BRIDGE_IP"), "READONLY")
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
		log.Errorf("failed to fetch user: %v", err.Error())
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

// GetUsers ...
func GetUsers(w http.ResponseWriter, req *http.Request) {
	conn := postgres.NewConn(viper.GetString("dev.BRIDGE_IP"), "READONLY")
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

	users, err := fetchUsers(conn)

	if err != nil {
		log.Errorf("failed to fetch users: %v", err.Error())
	}

	res = JSONresponse{
		Result: Result{
			Data: map[string]interface{}{
				"users": users,
			},
		},
	}

	writeJSONresponse(w, req.Header, start, nil, res)
}

func fetchUsers(conn *postgres.Conn) (*[]User, error) {
	rows, err := conn.DB.Query("SELECT * FROM users_v")
	defer rows.Close()

	if err != nil {
		log.Infof("failed to fetch users: %s", err.Error())
	}

	users := []User{}

	for rows.Next() {
		u := &User{}

		err = rows.Scan(&u.AvatarURL, &u.Email, &u.FirstName, &u.ID, &u.LastLogin, &u.LastName, &u.Location, &u.PostCount, &u.Username)
		users = append(users, *u)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to scan user response to destination: %s", err)
	}

	return &users, nil
}

// GetIndex ...
func GetIndex(w http.ResponseWriter, req *http.Request) {
	conn := postgres.NewConn(viper.GetString("dev.BRIDGE_IP"), "READONLY")

	err := conn.Open()

	var res JSONresponse
	start := time.Now()

	if err != nil {
		res = JSONresponse{
			Error: err.Error(),
		}

		writeJSONresponse(w, req.Header, start, err, res)
		return
	}

	defer conn.Close()

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
