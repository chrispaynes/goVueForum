package handlers

import (
	"net/http"
	"time"
)

// Post represents ...
type Post struct {
	Author    string `json:"author"`
	Body      string `json:"body"`
	Timestamp uint32 `json:"timestamp"`
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
			Author:    "Test Author",
			Body:      "Test Post Body",
			Timestamp: 1529886553,
		},
		{
			Author:    "Test Author2",
			Body:      "Test Post Body2",
			Timestamp: 1529886554,
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
