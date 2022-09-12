package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/fields"

	// "github.com/michimani/gotwi/tweet/managetweet/types"
	"github.com/michimani/gotwi/user/userlookup"
	"github.com/michimani/gotwi/user/userlookup/types"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	in := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           os.Getenv("accessToken"),
		OAuthTokenSecret:     os.Getenv("accessTokenSecret"),
	}

	c, err := gotwi.NewClient(in)
	if err != nil {
		fmt.Println(err)
		return
	}

	// p := &types.CreateInput{
	// 	Text: gotwi.String("測試 post 123123"),
	// 	Poll: &types.CreateInputPoll{
	// 		DurationMinutes: gotwi.Int(5),
	// 		Options: []string{
	// 			"1",
	// 			"2",
	// 			"3",
	// 			"4",
	// 		},
	// 	},
	// }

	// res, err := managetweet.Create(context.Background(), c, p)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Printf("[%s] %s\n", gotwi.StringValue(res.Data.ID), gotwi.StringValue(res.Data.Text))

	p := &types.GetByUsernameInput{
		Username: "elonmusk",
		Expansions: fields.ExpansionList{
			fields.ExpansionPinnedTweetID,
		},
		UserFields: fields.UserFieldList{
			fields.UserFieldCreatedAt,
		},
		TweetFields: fields.TweetFieldList{
			fields.TweetFieldCreatedAt,
		},
	}

	u, err := userlookup.GetByUsername(context.Background(), c, p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ID:          ", gotwi.StringValue(u.Data.ID))
	fmt.Println("Name:        ", gotwi.StringValue(u.Data.Name))
	fmt.Println("Username:    ", gotwi.StringValue(u.Data.Username))
	fmt.Println("CreatedAt:   ", u.Data.CreatedAt)
	if u.Includes.Tweets != nil {
		for _, t := range u.Includes.Tweets {
			fmt.Println("PinnedTweet: ", gotwi.StringValue(t.Text))
		}
	}
}
