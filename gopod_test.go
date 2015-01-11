//
// This package implements a RSS/Podcast 2.0 feed generator
//
// Test code for gopod.
//
// Use `go test` on the commandline to test this package
//
//

package gopod

import (
	"fmt"
	"testing"
	"time"
)

func Test1(*testing.T) {
	fmt.Printf("Testing gopod...\n")

	c := ChannelFactory("Daniel's Channel", "http://RubyDeveloper.com/", "My Blog", "http://example.com/image.png")
	c.SetPubDate(time.Now().UTC())
	c.SetiTunesExplicit("No")

	c.AddItem(&Item{
		Title:       "Stack Overflow",
		Link:        "http://stackoverflow.com",
		Description: "Stack Overflow",
		PubDate:     time.Now().UTC().Format(time.RFC1123),
	})

	// Example: Using an item's methods
	t := "My title"
	l := "http://linkedin.com"
	i := &Item{
		Title:       	t,
		TunesSubtitle:	t,
		Link:        	l,
		Description: 	"My LinkedIn",
		TunesDuration:	"600",
		TunesSummary:	"I asked myself that question more than a decade ago and it changed my...",
		Guid:			l,
		Creator:		"Daniel's Channel",
	}
	i.SetEnclosure("http://example.com/sound.mp3", "600", "audio/mpeg")
	i.SetPubDate(time.Now().Unix())
	c.AddItem(i)

	fmt.Printf("%s\n\n", c.Publish())
	fmt.Printf("%s\n\n", c.PublishIndent())
}
