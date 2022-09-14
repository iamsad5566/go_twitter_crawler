package getusr

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/tweet/searchtweet"
	sttypes "github.com/michimani/gotwi/tweet/searchtweet/types"
	"github.com/michimani/gotwi/user/userlookup"
	"github.com/michimani/gotwi/user/userlookup/types"
)

type twitterUser struct {
	ID        string
	Name      string
	UserName  string
	CreatedAt *time.Time
}

func (twiusr twitterUser) displayName() string {
	return fmt.Sprintf("%s@%s", twiusr.Name, twiusr.UserName)
}

func GetUserInfo(client *gotwi.Client, userName string) {
	p := &types.GetByUsernameInput{
		Username: userName,
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

	u, err := userlookup.GetByUsername(context.Background(), client, p)
	if err != nil {
		fmt.Println(err)
		return
	}

	twiusr := twitterUser{
		ID:        gotwi.StringValue(u.Data.ID),
		Name:      gotwi.StringValue(u.Data.Name),
		UserName:  gotwi.StringValue(u.Data.Username),
		CreatedAt: u.Data.CreatedAt,
	}

	fmt.Printf("ID\t\t:%s\n", twiusr.ID)
	fmt.Printf("Name\t\t:%s\n", twiusr.Name)
	fmt.Printf("UserName\t:%s\n", twiusr.UserName)
	fmt.Printf("CreatedAt\t:%v\n", twiusr.CreatedAt)
	fmt.Println("")

	getTweets(client, &twiusr)
}

func getTweets(client *gotwi.Client, twiusr *twitterUser) {
	p := &sttypes.ListRecentInput{
		MaxResults:  10,
		Query:       "from:" + twiusr.UserName + " -is:retweet -is:reply",
		TweetFields: fields.TweetFieldList{fields.TweetFieldCreatedAt},
	}

	res, err := searchtweet.ListRecent(context.Background(), client, p)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("----- %s's recent Tweets -----\n", twiusr.displayName())
	for k, t := range res.Data {
		if k >= 5 {
			break
		}
		fmt.Printf("[%s] %s\n", t.CreatedAt, gotwi.StringValue(t.Text))
	}
}
