package search

import (
	"context"
	"fmt"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/filteredstream"
	"github.com/michimani/gotwi/tweet/filteredstream/types"
)

func SearchStream(client *gotwi.Client) {
	p := &types.CreateRulesInput{
		Add: []types.AddingRule{
			{Value: gotwi.String("Twitter API v2"), Tag: gotwi.String("example rule")},
		},
	}

	res, err := filteredstream.CreateRules(context.TODO(), client, p)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)

	// for _, r := range res.Data {
	// 	fmt.Printf("ID: %s, Value: %s, Tag: %s\n", gotwi.StringValue(r.ID), gotwi.StringValue(r.Value), gotwi.StringValue(r.Tag))
	// }
}
