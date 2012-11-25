About
=====

This is a library implementing an RSS feed generator for the Go programming
language (http://golang.org/).

RSS (Rich Site Summary, or Really Simple Syndication) is a data format used to
publish frequently updated works - such as blog entries, news headlines, audio
and video - in a standardized format.

An RSS document (which is called a "feed", "web feed", or "channel") includes
full or summarized text, plus metadata such as publishing dates and authorship.


Installing
==========

    $ go get github.com/baliw/moverss
	
Example
=======

```go
package main

import (
	"github.com/baliw/moverss"
	"fmt"
)

func main() {
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
```

Full documentation
==================

Read it [online](http://go.pkgdoc.org/github.com/baliw/moverss) or run 

    $ go doc github.com/baliw/moverss

Other Details
=====================

A few of the more obscure data points aren't implemented yet.
I haven't seen feeds that implement them yet and I just want to make sure my implementation is good.

If you find this package useful and want to give back, I'd greatly appreciate a link to my [Ruby developer](http://RubyDeveloper.com/) site.






