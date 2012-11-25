// Copyright 2012 Daniel Walton - All rights reserved.
// Author: Daniel Walton <dan@RubyDeveloper.com>
//         https://github.com/baliw
//         @dan_gogh

//
// This package implements a rss 2.0 feed generator based on the spec
// located at: http://cyber.law.harvard.edu/rss/rss.html
//
// Test code for moverss.
//
// Use `go test` on the commandline to test this package
//
//

package moverss

import (
	"fmt"
	"testing"
	"time"
)

func Test1(*testing.T) {
	fmt.Printf("Testing Moverss...\n")

	c := ChannelFactory("Daniel's Channel", "http://RubyDeveloper.com/", "My Blog")
	c.SetPubDate(time.Now().UTC())

	c.AddItem(&Item{
		Title:       "Ruby Developer",
		Link:        "http://RubyDeveloper.com/",
		Description: "Ruby Developer",
		PubDate:     time.Now().UTC().Format(time.RFC1123),
	})
	c.AddItem(&Item{
		Title:       "Stack Overflow",
		Link:        "http://stackoverflow.com/users/1305696/daniel",
		Description: "Stack Overflow",
		PubDate:     time.Now().UTC().Format(time.RFC1123),
	})

	// Example: Using an item's SetPubDate method
	i := &Item{
		Title:       "LinkedIn",
		Link:        "http://www.linkedin.com/in/dangogh",
		Description: "My LinkedIn",
	}
	i.SetPubDate(time.Now().Unix())
	c.AddItem(i)

	fmt.Printf("%s\n\n", c.Publish())
	fmt.Printf("%s\n\n", c.PublishIndent())
}
