package post

import (
	"context"
	"fmt"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)

func Tweet(client *gotwi.Client) {
	post := &types.CreateInput{
		Text: gotwi.String("This is a test post from robot\ntest link: https://www.tw-yk.com/"),
	}

	res, err := managetweet.Create(context.Background(), client, post)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("[%s] %s\n", gotwi.StringValue(res.Data.ID), gotwi.StringValue(res.Data.Text))
}
