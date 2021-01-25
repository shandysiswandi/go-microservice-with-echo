package jsonplaceholder

import "errors"

// The URL from jsonplaceholder
const (
	BaseURL   string = "https://jsonplaceholder.typicode.com"
	PostsPath string = BaseURL + "/posts"
	UsersPath string = BaseURL + "/users"
)

// The common error from post
var (
	ErrPostFetch  = errors.New("Failed Fetch From " + PostsPath)
	ErrPostGet    = errors.New("Failed Get From " + PostsPath)
	ErrPostCreate = errors.New("Failed Create From " + PostsPath)
)
