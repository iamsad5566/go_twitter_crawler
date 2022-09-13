package main

import (
	"log"
	"os"
	"twitter-crawler/post"

	"github.com/joho/godotenv"
	"github.com/michimani/gotwi"
)

func main() {
	client, err := newClient()
	if err != nil {
		log.Println(err)
	}

	post.Tweet(client)
}

func newClient() (*gotwi.Client, error) {
	// Load the environment variables first
	loadEnv()
	os.Setenv("GOTWI_API_KEY", os.Getenv("key"))
	os.Setenv("GOTWI_API_KEY_SECRET", os.Getenv("secret"))

	// Login
	in := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           os.Getenv("accessToken"),
		OAuthTokenSecret:     os.Getenv("accessTokenSecret"),
	}

	return gotwi.NewClient(in)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
}
