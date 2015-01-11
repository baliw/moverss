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
