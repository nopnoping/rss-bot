package rsspull

type RssPull struct {
	client *rssClient
}

func NewRssPull() *RssPull {
	return &RssPull{}
}

func (r *RssPull) Pull(url string) {

}
