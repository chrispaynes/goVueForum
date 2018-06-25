package handlers

import (
	"net/http"
	"time"
)

// Post represents ...
type Post struct {
	Title       string `json:"title"`
	Author      Author `json:"author"`
	Body        string `json:"body"`
	ID          uint32 `json:"id"`
	Timestamp   uint32 `json:"timestamp"`
	LastUpdated uint32 `json:"lastUpdatedAt"`
}

// Author represents...
type Author struct {
	LastLogin uint32 `json:"lastLogin"`
	ID        uint32 `json:"id"`
	Username  string `json:"username"`
}

// GetHealth serves as a simple server health check
func GetHealth(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

// GetIndex ...
func GetIndex(w http.ResponseWriter, req *http.Request) {
	var res JSONresponse

	start := time.Now()

	// get index error
	// res = JSONresponse{
	// 	Error: err.Error(),
	// }

	// if err != nil {
	// 	writeJSONresponse(w, req.Header, start, err, res)
	// 	return
	// }

	posts := []Post{
		{
			Title: "Test Title",
			Author: Author{
				Username:  "Test Author",
				ID:        123,
				LastLogin: 1529886553,
			},
			Body:        "Test Post Body",
			ID:          1,
			Timestamp:   1529886553,
			LastUpdated: 1529886553,
		},
		{
			Title: "Test Title2",
			Author: Author{
				Username:  "Test Author2",
				ID:        456,
				LastLogin: 1529886553,
			},
			Body:        "Test Post Body2",
			ID:          2,
			Timestamp:   1529886554,
			LastUpdated: 1529886554,
		},
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
