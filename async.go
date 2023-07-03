package main

import (
	"encoding/json"
	"net/http"
)

type PostResponse struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func Async() []PostResponse {
	// Make the http request
	url := "https://jsonplaceholder.typicode.com/posts"

	resp, err := http.Get(url)
	if err != nil {
		print(err)
	}

	// Close the body
	defer resp.Body.Close()

	var posts []PostResponse

	// Decode the JSON response
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		print(err)
	}
	// Print the result on the console
	// fmt.Printf("UserId: %v\n", post.UserId)
	// fmt.Printf("Id: %v\n", post.Id)
	// fmt.Printf("Title: %v\n", post.Title)
	// fmt.Println(post)
	// for i := 0; i < len(posts); i++ {
	// 	post := posts[i]
	// 	fmt.Println(post.Body)

	// }
	return posts
}
