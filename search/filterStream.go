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
			{Value: gotwi.String("Elon mask"), Tag: gotwi.String("mask")},
		},
	}

	res, err := filteredstream.CreateRules(context.TODO(), client, p)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, r := range res.Data {
		fmt.Printf("ID: %s, Value: %s, Tag: %s\n", gotwi.StringValue(r.ID), gotwi.StringValue(r.Value), gotwi.StringValue(r.Tag))
	}
}
