// Package gopod implements a RSS/Podcast 2.0 feed generator
//
//
// RSS (Rich Site Summary, or Really Simple Syndication) is a data format used
// to publish frequently updated works - such as blog entries, news headlines,
// audio and video - in a standardized format. Podcasts utilize this format 
// to deliver audio and video files and assiciated metadata.
//
// An RSS document (which is called a "feed", "web feed", or "channel") includes
// full or summarized text, plus metadata such as publishing dates and
// authorship.
//
// Spec located at: http://cyber.law.harvard.edu/rss/rss.html
//
// Example Usage
//
//		c := gopod.ChannelFavtory("Ruby Developer", "http://RubyDeveloper.com/", "Ruby Developer Blog")
//		c.AddItem(&gopod.Item{
//			Title:"Ruby Developer",
//			Link:"http://RubyDeveloper.com/",
//			Description:"My Blog",
//			PubDate:time.Now().Unix(),
//		})
//		c.AddItem(&gopod.Item{
//			Title:"Stack Overflow",
//			Link:"http://stackoverflow.com/users/1305696/daniel",
//			Description:"My Stack Overflow",
//			PubDate:time.Now().Unix(),
//		})
//		c.AddItem(&gopod.Item{
//			Title:"LinkedIn",
//			Link:"http://www.linkedin.com/in/dangogh",
//			Description:"My LinkedIn",
//			PubDate:time.Now().Unix(),
//		})
//		output := c.Publish()
//
//
//
package gopod

import (
	"bytes"
	"encoding/xml"
	"reflect"
	"time"
)

type Channel struct {
	// The name of the channel. It's how people refer to your service. If you have
	// an HTML website that contains the same information as your RSS file, the
	// title of your channel should be the same as the title of your website.
	Title string `xml:"title"`
	// The URL to the HTML website corresponding to the channel.
	Link string `xml:"link"`
	// Phrase or sentence describing the channel.
	Description string `xml:"description"`

	/////////////////////
	// Optional Fields //
	/////////////////////
	Language       	string `xml:"language,omitempty"`
	Copyright      	string `xml:"copyright,omitempty"`
	ManagingEditor 	string `xml:"managingEditor,omitempty"`
	WebMaster      	string `xml:"webMaster,omitempty"`
	PubDate        	string `xml:"pubDate,omitempty"`
	LastBuildDate  	string `xml:"lastBuildDate,omitempty"`
	Category       	string `xml:"category,omitempty"`
	Generator      	string `xml:"generator,omitempty"`
	Docs           	string `xml:"docs,omitempty"`
	TTL            	string `xml:"ttl,omitempty"`
	SkipHours      	string `xml:"skiphours,omitempty"`
	SkipDays       	string `xml:"skipdays,omitempty"`
	TunesAuthor    	string `xml:"itunes:author,omitempty"`
	TunesSubtitle  	string `xml:"itunes:subtitle,omitempty"`
	TunesSummary   	string `xml:"itunes:summary,omitempty"`
	TunesExplicit  	string `xml:"itunes:explicit,omitempty"`

	TunesOwner		[]*TunesOwner
	TunesImage		[]*TunesImage
	AtomLink		[]*AtomLink
	Items 			[]*Item

	// [Fields to be implemented]
	//  Cloud
	//  Image
	//  Rating
	//  textInput

	// Stub member just for the xml generator.
	XMLName xml.Name `xml:"channel"`
}

type TunesOwner struct {
	Name	string `xml:"itunes:name,omitempty"`
	Email	string `xml:"itunes:email,omitempty"`

	XMLName xml.Name `xml:"itunes:owner"`	
}

type TunesImage struct {
	Href		string `xml:"href,attr"`

	XMLName xml.Name `xml:"itunes:image"`	
}

type AtomLink struct {
	Href		string `xml:"href,attr"`
	Rel			string `xml:"rel,attr"`
	Type		string `xml:"type,attr"`

	XMLName xml.Name `xml:"atom:link"`	
}

type Item struct {
	//The URL of the item.
	Link string `xml:"link"`

	// The item synopsis.
	Description string `xml:"description"`

	// The title of the item.
	Title string `xml:"title,omitempty"`

	// Indicates when the item was published.
	// http://cyber.law.harvard.edu/rss/rss.html#ltpubdategtSubelementOfLtitemgt
	PubDate string `xml:"pubDate,omitempty"`

	// Email address of the author of the item.
	// http://cyber.law.harvard.edu/rss/rss.html#ltauthorgtSubelementOfLtitemgt
	Author string `xml:"author,omitempty"`

	// A string that uniquely identifies the item.
	// http://cyber.law.harvard.edu/rss/rss.html#ltguidgtSubelementOfLtitemgt
	Guid string `xml:"guid,omitempty"`

	// URL of a page for comments relating to the item.
	// http://cyber.law.harvard.edu/rss/rss.html#ltcommentsgtSubelementOfLtitemgt
	Comments string `xml:"comments,omitempty"`

	Creator 		string `xml:"dc:creator,omitempty"`
	TunesAuthor    	string `xml:"itunes:author,omitempty"`
	TunesSubtitle  	string `xml:"itunes:subtitle,omitempty"`
	TunesSummary   	string `xml:"itunes:summary,omitempty"`
	TunesExplicit  	string `xml:"itunes:explicit,omitempty"`
	TunesDuration  	string `xml:"itunes:duration,omitempty"`

	Enclosure		[]*Enclosure

	// Stub member just for the xml generator.
	XMLName xml.Name `xml:"item"`
}

type Enclosure struct	{
	Url 		string `xml:"url,attr"`
	Length 		string `xml:"length,attr"`
	Type 		string `xml:"type,attr"`

	XMLName xml.Name `xml:"enclosure"`
}

// Use this to create the base channel of the rss feed.
// The Title is the name of the channel. It's how people refer to your
// service. If you have an HTML website that contains the same information
// as your RSS file, the title of your channel should be the same as the
// title of your website.
// The Link is the URL to the HTML website corresponding to the channel.
// The Description is the phrase or sentence describing the channel.
// The Image is the image to be displayed with the podcast.
func ChannelFactory(Title string, Link string, Description string, Image string) *Channel {
	c := &Channel{Title: Title, Link: Link, Description: Description}
	c.Generator = "gopod - http://github.com/jbckmn/gopod"
	c.SetTunesImage(&TunesImage{
			Href:	Image,
		})
	c.AtomLink = append(c.AtomLink, &AtomLink{
			Href:	Link,
			Rel:	"self",
			Type:	"application/rss+xml",
		})
	return c
}

// Add an image struct to the feed
func (c *Channel) SetTunesImage(i *TunesImage) {
	c.TunesImage = append(c.TunesImage, i)
}

// Add an item to the feed list
func (c *Channel) AddItem(i *Item) {
	c.Items = append(c.Items, i)
}

// Publish the prepared rss feed without indentation
func (c *Channel) Publish() []byte {
	output, err := xml.Marshal(c)
	if err != nil {
		panic(err)
	}
	return bytes.Join([][]byte{
		[]byte(`<?xml version="1.0" encoding="UTF-8"?>`),
		[]byte(`<rss xmlns:atom="http://www.w3.org/2005/Atom" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" version="2.0">`),
		output,
		[]byte(`</rss>`),
	}, []byte(""))
}

// Publish the prepared rss feed with indentation
func (c *Channel) PublishIndent() []byte {
	output, err := xml.MarshalIndent(c, "  ", "    ")
	if err != nil {
		panic(err)
	}
	return bytes.Join([][]byte{
		[]byte(`<?xml version="1.0" encoding="UTF-8"?>`),
		[]byte(`<rss xmlns:atom="http://www.w3.org/2005/Atom" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" version="2.0">`),
		output,
		[]byte(`</rss>`),
	}, []byte("\n"))
}

// (optional) Set the channel's language field to the language the channel
// is written in. This allows aggregators to group all
// Italian language sites, for example, on a single page. A list of allowable
// values for this element, as provided by Netscape (
// http://cyber.law.harvard.edu/rss/languages.html ).
// You may also use values defined by the W3C (
// http://www.w3.org/TR/REC-html40/struct/dirlang.html#langcodes )
func (c *Channel) SetLanguage(Language string) {
	c.Language = Language
}

// (optional) Set the channel's Copyright notice for content.
func (c *Channel) SetCopyright(Copyright string) {
	c.Copyright = Copyright
}

// (optional) Set the channel's ManagingEditor field.
// The ManagingEditor field should contain the email address for the person
// responsible for editorial content.
func (c *Channel) SetManagingEditor(ManagingEditor string) {
	c.ManagingEditor = ManagingEditor
}

// (optional) Set the channel's webMaster field.
// The webmaster field should contain the email address of person responsible
// for technical issues relating to channel
func (c *Channel) SetWebMaster(WebMaster string) {
	c.WebMaster = WebMaster
}

// (optional) Set the channel's pubDate field.  This method can take an instance
// of time.Time, a unix timestamp (int64) or a raw string to use for the date.
// The date should represent the publication date for the content in the
// channel.  For example, the New York Times publishes on a daily basis, the
// publication date flips once every 24 hours. That's when the pubDate of
// the channel changes. All date-times in RSS conform to the Date and Time
// Specification of RFC 822, with the exception that the year may be expressed
// with two characters or four characters (four preferred).
func (c *Channel) SetPubDate(t interface{}) {
	if reflect.TypeOf(t) == reflect.TypeOf(time.Now()) {
		c.PubDate = t.(time.Time).UTC().Format(time.RFC1123)
		return
	} else if reflect.TypeOf(t) == reflect.TypeOf("") {
		c.PubDate = t.(string)
		return
	} else if reflect.TypeOf(t) == reflect.TypeOf(int64(1)) {
		c.PubDate = time.Unix(t.(int64), 0).UTC().Format(time.RFC1123)
		return
	}
	panic("gopod.Channel.SetPubDate() error: Invalid date type")
}

// (optional) Set the lastBuildDate field. The lastBuildDate is The last time
// the content of the channel changed.  This method can take an instance of
// time.Time, a Unix timestamp (int64) or a raw string.
func (c *Channel) SetLastBuildDate(t interface{}) {
	if reflect.TypeOf(t) == reflect.TypeOf(time.Now()) {
		c.LastBuildDate = t.(time.Time).UTC().Format(time.RFC1123)
		return
	} else if reflect.TypeOf(t) == reflect.TypeOf("") {
		c.LastBuildDate = t.(string)
		return
	} else if reflect.TypeOf(t) == reflect.TypeOf(int64(1)) {
		c.LastBuildDate = time.Unix(t.(int64), 0).UTC().Format(time.RFC1123)
		return
	}
	panic("gopod.Channel.SetPubDate() error: Invalid date type")
}

// (optional) Set the category field.  The category field specifies one or more
// categories that the channel belongs to. It follows the same rules as the
// <item>-level category element.
// ( http://cyber.law.harvard.edu/rss/rss.html#syndic8 )
func (c *Channel) SetCategory(cat string) {
	c.Category = cat
}

// (optional) Set generator field.  The generator field is a string indicating
// the program used to generate the rss feed.
func (c *Channel) SetGenerator(gen string) {
	c.Generator = gen
}

// (optional) Set the docs field.  The docs field is a URL that points to the
// documentation for the format used in the RSS file.  It's usually a
// pointer to "http://cyber.law.harvard.edu/rss/rss.html".
// It's for people who might stumble across an RSS file on a Web server
// 25 years from now and wonder what it is.
func (c *Channel) SetDocs(docs string) {
	c.Docs = docs
}

// (optional) Set the TTL field.  TTL stands for time to live.  It's a number
// of minutes that indicate how long a channel can be cached before refreshing
// from the source.  More info here:
// http://cyber.law.harvard.edu/rss/rss.html#ltttlgtSubelementOfLtchannelgt
func (c *Channel) SetTTL(ttl string) {
	c.TTL = ttl
}

// (optional) Set the channel's skipDays field.  The skipDays field is a hint
// for aggregators telling them which days they can skip.  More info:
// http://cyber.law.harvard.edu/rss/skipHoursDays.html#skipdays
func (c *Channel) SetSkipHours(skiphours string) {
	c.SkipHours = skiphours
}

// (optional) Set the channel's skipHous field.  The skipHours field is a hint
// for aggregators telling them which hours they can skip.  More info:
// http://cyber.law.harvard.edu/rss/skipHoursDays.html#skiphours
func (c *Channel) SetSkipDays(skipdays string) {
	c.SkipDays = skipdays
}

// (optional) Set the channel's iTunes Explicit field. 
func (c *Channel) SetiTunesExplicit(explicit string) {
	c.TunesExplicit = explicit
}

// (optional) Set the channel's iTunes Author field. 
func (c *Channel) SetiTunesAuthor(author string) {
	c.TunesAuthor = author
}

// (optional) Set the channel's iTunes Subtitle field. 
func (c *Channel) SetiTunesSubtitle(subtitle string) {
	c.TunesSubtitle = subtitle
}

// (optional) Set the channel's iTunes Summary field. 
func (c *Channel) SetiTunesSummary(summary string) {
	c.TunesSummary = summary
}

// (optional) Set the channel's iTunes Owner struct. 
func (c *Channel) SetiTunesOwner(name string, email string) {
	c.TunesOwner = append(c.TunesOwner, &TunesOwner{
			Name:	name,
			Email: 	email,
		})
}

// (optional) Set the item's pubDate field.  This method can take an instance
// of time.Time, a unix timestamp (int64) or a raw string to use for the date.
// The date should represent the publication date for the content in the
// channel.  For example, the New York Times publishes on a daily basis, the
// publication date flips once every 24 hours. That's when the pubDate of
// the channel changes. All date-times in RSS conform to the Date and Time
// Specification of RFC 822, with the exception that the year may be expressed
// with two characters or four characters (four preferred).
func (i *Item) SetPubDate(t interface{}) {
	if reflect.TypeOf(t) == reflect.TypeOf(time.Now()) {
		i.PubDate = t.(time.Time).UTC().Format(time.RFC1123)
		return
	} else if reflect.TypeOf(t) == reflect.TypeOf("") {
		i.PubDate = t.(string)
		return
	} else if reflect.TypeOf(t) == reflect.TypeOf(int64(1)) {
		i.PubDate = time.Unix(t.(int64), 0).UTC().Format(time.RFC1123)
		return
	}
	panic("gopod.Item.SetPubDate() error: Invalid date type")
}

// (optional) Set the item's iTunes Enclosure struct. 
func (i *Item) SetEnclosure(url string, length string, eType string) {
	i.Enclosure = append(i.Enclosure, &Enclosure{
			Url: 	url,
			Length:	length,
			Type:	eType,
		})
}