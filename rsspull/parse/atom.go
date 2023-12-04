package parse

import (
	"github.com/beevik/etree"
)

type atomV0_3 struct {
}

var _ Parser = (*atomV0_3)(nil)

func (a atomV0_3) Parse(root *etree.Element) *FeedInfo {
	panic("Implement it")
}

func init() {
	ParserMap["atom-0.3"] = &atomV0_3{}
}
