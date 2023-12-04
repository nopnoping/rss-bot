package parse

import (
	"github.com/beevik/etree"
	"log"
	"testing"
)

var atomV0_3_Data = `<?xml version="1.0" encoding="utf-8"?>
<feed version="0.3" 
      xmlns="http://purl.org/atom/ns#" >
  <Title>atom_0.3.feed.Title</Title>
  <Link rel="alternate" type="atom_0.3.feed.Link^type" href="atom_0.3.feed.Link^href"/>
  <author>
    <name>atom_0.3.feed.author.name</name>
    <url>atom_0.3.feed.author.url</url>
    <email>atom_0.3.feed.author.email</email>
  </author>
  <contributor>
    <name>atom_0.3.feed.contributor.name</name>
    <url>atom_0.3.feed.contributor.url</url>
    <email>atom_0.3.feed.contributor.email</email>
  </contributor>
  <tagline type="atom_0.3.feed.tagline^type" mode="xml">atom_0.3.feed.tagline</tagline>
  <id>atom_0.3.feed.id</id>
  <generator url="atom_0.3.feed.generator^url">atom_0.3.feed.generator</generator>
  <copyright>atom_0.3.feed.copyright</copyright>
  <info type="atom_0.3.feed.info^type" mode="xml">atom_0.3.feed.info</info>
  <modified>2000-01-01T00:00:00Z</modified>
  <entry>
    <Title>atom_0.3.feed.entry[0].Title</Title>
    <Link rel="alternate" type="atom_0.3.feed.entry[0].Link^type"
          href="atom_0.3.feed.entry[0].Link^href"/>
    <id>atom_0.3.feed.entry[0]^id</id>
    <author>
      <name>atom_0.3.feed.entry[0].author.name</name>
      <url>atom_0.3.feed.entry[0].author.url</url>
      <email>atom_0.3.feed.entry[0].author.email</email>
    </author>
    <contributor>
      <name>atom_0.3.feed.entry[0].contributor.name</name>
      <url>atom_0.3.feed.entry[0].contributor.url</url>
      <email>atom_0.3.feed.entry[0].contributor.email</email>
    </contributor>
    <modified>2000-01-01T00:00:00Z</modified>
    <issued>2000-01-01T01:00:00Z</issued>
    <created>2000-01-01T02:00:00Z</created>
    <summary type="atom_0.3.feed.entry[0].summary^type" mode="xml">atom_0.3.feed.entry[0].summary</summary>
    <content type="atom_0.3.feed.entry[0].content[0]^type" mode="xml">atom_0.3.feed.entry[0].content[0]</content>
    <content type="atom_0.3.feed.entry[0].content[1]^type" mode="xml">atom_0.3.feed.entry[0].content[1]</content>
  </entry>
    <entry>
      <Title>atom_0.3.feed.entry[1].Title</Title>
      <Link rel="alternate" type="atom_0.3.feed.entry[1].Link^type"
            href="atom_0.3.feed.entry[1].Link^href"/>
      <id>atom_0.3.feed.entry[1]^id</id>
      <author>
        <name>atom_0.3.feed.entry[1].author.name</name>
        <url>atom_0.3.feed.entry[1].author.url</url>
        <email>atom_0.3.feed.entry[1].author.email</email>
      </author>
      <contributor>
        <name>atom_0.3.feed.entry[1].contributor.name</name>
        <url>atom_0.3.feed.entry[1].contributor.url</url>
        <email>atom_0.3.feed.entry[1].contributor.email</email>
      </contributor>
      <modified>2000-02-01T00:00:00Z</modified>
      <issued>2000-02-01T01:00:00Z</issued>
      <created>2000-02-01T02:00:00Z</created>
      <summary type="atom_0.3.feed.entry[1].summary^type" mode="xml">atom_0.3.feed.entry[1].summary</summary>
      <content type="atom_0.3.feed.entry[1].content[0]^type" mode="xml">atom_0.3.feed.entry[1].content[0]</content>
      <content type="atom_0.3.feed.entry[1].content[1]^type" mode="xml">atom_0.3.feed.entry[1].content[1]</content>
    </entry>
</feed>
`

var atomV1_0_Data = `<?xml version="1.0" encoding="utf-8"?>
<feed xmlns='http://www.w3.org/2005/Atom' 
      xml:lang='en-us'>
  <Title type="html">atom_1.0.feed.Title</Title>
  <Link rel="self" type="text/html" href="http://example.com/blog/atom_1.0.xml"/>
  <Link rel="alternate" type="text/html" href="http://example.com/blog"/>
  <Link rel="alternate" type="text/plain" href="http://example.com/blog_plain"/>
  <author>
    <name>atom_1.0.feed.author.name</name>
    <uri>http://example.com</uri>
    <email>author0@example.com</email>
  </author>
  <contributor>
    <name>atom_1.0.feed.contributor.name</name>
    <uri>http://example.com</uri>
    <email>author1@example.com</email>
  </contributor>
  <subtitle type="html">atom_1.0.feed.tagline</subtitle>
  <id>http://example.com/blog/atom_1.0.xml</id>
  <generator uri="http://example.com/test">atom_1.0.feed.generator</generator>
<rights>atom_1.0.feed.copyright</rights>
  <updated>2000-01-01T00:00:00Z</updated>
  <entry>
    <Title type="text">atom_1.0.feed.entry[0].Title</Title>
    <Link rel="alternate" type="text/html"
          href="http://example.com/blog/entry1"/>
    <Link rel="alternate" type="text/plain"
          href="http://example.com/blog/entry1_plain"/>
	<Link rel="enclosure" type="image/gif"
          href="http://example.com/blog/enclosure1.gif"/>
	<Link rel="test" type="image/gif"
          href="tag:example.com,2005:Atom-Tests:xml-base:Test0"/>
    <id>atom_1.0.feed.entry[0]^id</id>
    <author>
      <name>atom_1.0.feed.entry[0].author.name</name>
      <uri>http://example.com</uri>
      <email>author0@example.com</email>
    </author>
    <contributor>
      <name>atom_1.0.feed.entry[0].contributor.name</name>
      <uri>http://example.com</uri>
      <email>author1@example.com</email>
    </contributor>
    <updated>2000-01-01T00:00:00Z</updated>
    <published>2000-01-01T01:00:00Z</published>
    <summary type="html">atom_1.0.feed.entry[0].summary</summary>
    <content type="html">atom_1.0.feed.entry[0].content[0]</content>
	<rights>atom_1.0.feed.entry[0].rights</rights>
  </entry>
    <entry>
      <Title type="text">atom_1.0.feed.entry[1].Title</Title>
      <Link rel="alternate" type="text/html"
            href="http://example.com/blog/entry2"/>
      <Link rel="enclosure" type="image/gif"
	        href="http://example.com/blog/enclosure2.gif"/>
	<Link rel="test" type="image/gif"
          href="tag:example.com,2005:Atom-Tests:xml-base:Test1"/>
      <id>atom_1.0.feed.entry[1]^id</id>
      <author>
        <name>atom_1.0.feed.entry[1].author.name</name>
        <uri>http://example.com</uri>
        <email>author0@example.com</email>
      </author>
      <contributor>
        <name>atom_1.0.feed.entry[1].contributor.name</name>
        <uri>http://example.com</uri>
        <email>author1@example.com</email>
      </contributor>
      <updated>2000-02-01T00:00:00Z</updated>
      <published>2000-02-01T01:00:00Z</published>
      <summary type="html">atom_1.0.feed.entry[1].summary</summary>
      <content type="html">atom_1.0.feed.entry[1].content[0]</content>
    </entry>
    </feed>
`

func TestAtomV0_3(t *testing.T) {
	data := ([]byte)(atomV0_3_Data)
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(data); err != nil {
		t.Fatalf("Doc Read bytes err:%v", err)
	}

	parser := atomV0_3{}
	feed := parser.Parse(doc.Root())
	log.Println(feed.Channel.Title, " ", feed.Channel.Link)
	for _, i := range feed.Items {
		log.Println(i.Title, " ", i.Link, " ", i.PubDate)
	}
}

func TestAtomV1_0(t *testing.T) {
	data := ([]byte)(atomV1_0_Data)
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(data); err != nil {
		t.Fatalf("Doc Read bytes err:%v", err)
	}

	parser := atomV1_0{}
	feed := parser.Parse(doc.Root())
	log.Println(feed.Channel.Title, " ", feed.Channel.Link)
	for _, i := range feed.Items {
		log.Println(i.Title, " ", i.Link, " ", i.PubDate)
	}
}
