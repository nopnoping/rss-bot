package parse

import (
	"github.com/beevik/etree"
	"log"
	"testing"
)

var rssV2_0_Data = `
<?xml version="1.0"?>
<rss version="2.0" 
    xmlns:content='http://purl.org/rss/1.0/modules/content/'>
    <channel>
        <title>rss_2.0.channel.title</title>
        <link>rss_2.0.channel.link</link>
        <description>rss_2.0.channel.description</description>
        <language>rss_2.0.channel.language</language>
        <rating>rss_2.0.channel.rating</rating>
        <copyright>rss_2.0.channel.copyright</copyright>
        <pubDate> Mon, 01 Jan 2001 00:00:00 GMT
        </pubDate>
        <lastBuildDate>Mon, 01 Jan 2001 01:00:00 GMT</lastBuildDate>
        <docs>rss_2.0.channel.docs</docs>
        <managingEditor>rss_2.0.channel.managingEditor</managingEditor>
        <webMaster>rss_2.0.channel.webMaster</webMaster>
        <category domain="rss_2.0.channel.category[0]^domain">rss_2.0.channel.category[0]</category>
        <category domain="rss_2.0.channel.category[1]^domain">rss_2.0.channel.category[1]</category>
        <generator>rss_2.0.channel.generator</generator>
        <ttl>100</ttl>

        <image>
            <title>rss_2.0.channel.image.title</title>
            <url>rss_2.0.channel.image.url</url>
            <link>rss_2.0.channel.image.link</link>
            <width>100</width>
            <height>200</height>
            <description>rss_2.0.channel.image.description</description>
        </image>

        <item>
            <title>rss_2.0.channel.item[0].title</title>
            <description>rss_2.0.channel.item[0].description</description>
            <link>rss_2.0.channel.item[0].link</link>
            <source url="rss_2.0.channel.item[0].source^url">rss_2.0.channel.item[0].source</source>
            <enclosure url="rss_2.0.channel.item[0].enclousure[0]^url" length="100"
                       type="rss_2.0.channel.item[0].enclousure[0]^type"/>
            <enclosure url="rss_2.0.channel.item[0].enclousure[1]^url" length="100"
                       type="rss_2.0.channel.item[0].enclousure[1]^type"/>
            <category domain="rss_2.0.channel.item[0].category[0]^domain">rss_2.0.channel.item[0].category[0]</category>
            <category domain="rss_2.0.channel.item[0].category[1]^domain">rss_2.0.channel.item[0].category[1]</category>
            <pubDate>Mon, 01 Jan 2001 00:00:00 GMT</pubDate>
            <expirationDate>Mon, 01 Jan 2001 01:00:00 GMT</expirationDate>
            <author>rss_2.0.channel.item[0].author</author>
            <comments>rss_2.0.channel.item[0].comments</comments>
            <guid isPermaLink="true">rss_2.0.channel.item[0].guid</guid>
            <content:encoded>rss_2.0.channel.item[0].content</content:encoded>
        </item>
        <item>
            <title>rss_2.0.channel.item[1].title</title>
            <description>rss_2.0.channel.item[1].description</description>
            <link>rss_2.0.channel.item[1].link</link>
            <source url="rss_2.0.channel.item[1].source^url">rss_2.0.channel.item[1].source</source>
            <enclosure url="rss_2.0.channel.item[1].enclousure[0]^url" length="100"
                       type="rss_2.0.channel.item[1].enclousure[0]^type"/>
            <enclosure url="rss_2.0.channel.item[1].enclousure[1]^url" length="100"
                       type="rss_2.0.channel.item[1].enclousure[1]^type"/>
            <category domain="rss_2.0.channel.item[1].category[0]^domain">rss_2.0.channel.item[1].category[0]</category>
            <category domain="rss_2.0.channel.item[1].category[1]^domain">rss_2.0.channel.item[1].category[1]</category>
            <pubDate>Mon, 02 Jan 2001 00:00:00 GMT</pubDate>
            <expirationDate>Mon, 01 Jan 2001 01:00:00 GMT</expirationDate>
            <author>rss_2.0.channel.item[1].author</author>
            <comments>rss_2.0.channel.item[1].comments</comments>
            <guid isPermaLink="false">rss_2.0.channel.item[1].guid</guid>
            <content:encoded>rss_2.0.channel.item[1].content</content:encoded>
        </item>

        <textInput>
            <title>rss_2.0.channel.textInput.title</title>
            <description>rss_2.0.channel.textInput.description</description>
            <name>rss_2.0.channel.textInput.name</name>
            <link>rss_2.0.channel.textInput.link</link>
        </textInput>
        <skipHours>
            <hours>0</hours>
            <hours>1</hours>
            <hours>2</hours>
            <hours>3</hours>
            <hours>4</hours>
            <hours>5</hours>
            <hours>6</hours>
            <hours>7</hours>
            <hours>8</hours>
            <hours>9</hours>
            <hours>10</hours>
            <hours>11</hours>
            <hours>12</hours>
            <hours>13</hours>
            <hours>14</hours>
            <hours>15</hours>
            <hours>16</hours>
            <hours>17</hours>
            <hours>18</hours>
            <hours>19</hours>
            <hours>20</hours>
            <hours>21</hours>
            <hours>23</hours>
        </skipHours>
        <skipdays>
          <day>Monday</day>
          <day>Tuesday</day>
          <day>Wednesday</day>
          <day>Thursday</day>
          <day>Friday</day>
          <day>Saturday</day>
          <day>Sunday</day>
        </skipdays>
        <cloud domain="rss_2.0.channel.cloud^domain" port="100" path="rss_2.0.channel.cloud^path"
                registerProcedure="rss_2.0.channel.cloud^registerProcedure"
                protocol="rss_2.0.channel.cloud^protocol"/>

    </channel>
</rss>
`

var rssV1_0_Data = `
<?xml version="1.0"?>
<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
         xmlns="http://purl.org/rss/1.0/"
         xmlns:dc="http://purl.org/dc/elements/1.1/">

  <channel rdf:about="http://www.example.com">
    <title>Example RSS 1.0 Feed</title>
    <link>http://www.example.com</link>
    <description>This is an example RSS 1.0 feed.</description>
    <items>
      <rdf:Seq>
        <rdf:li rdf:resource="http://www.example.com/item1"/>
        <rdf:li rdf:resource="http://www.example.com/item2"/>
      </rdf:Seq>
    </items>
  </channel>

  <item rdf:about="http://www.example.com/item1">
    <title>Item 1</title>
    <link>http://www.example.com/item1</link>
    <description>This is the first item in the feed.</description>
    <dc:date>2023-12-01T12:00:00Z</dc:date>
  </item>

  <item rdf:about="http://www.example.com/item2">
    <title>Item 2</title>
    <link>http://www.example.com/item2</link>
    <description>This is the second item in the feed.</description>
    <dc:date>2023-12-01T13:30:00Z</dc:date>
  </item>

</rdf:RDF>
`

func TestRssV2_0(t *testing.T) {
	data := ([]byte)(rssV2_0_Data)
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(data); err != nil {
		t.Fatalf("Doc Read bytes err:%v", err)
	}

	parser := rssV2_0{}
	feed := parser.Parse(doc.Root())
	log.Println(feed.channel.title, " ", feed.channel.link)
	for _, i := range feed.items {
		log.Println(i.title, " ", i.link, " ", i.pubDate)
	}
}

func TestRssV1_0(t *testing.T) {
	data := ([]byte)(rssV1_0_Data)
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(data); err != nil {
		t.Fatalf("Doc Read bytes err:%v", err)
	}

	parser := rssV1_0{}
	feed := parser.Parse(doc.Root())
	log.Println(feed.channel.title, " ", feed.channel.link)
	for _, i := range feed.items {
		log.Println(i.title, " ", i.link, " ", i.pubDate)
	}
}
