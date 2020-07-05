package model

type SearchResult struct {
	Hits int
	Start int
	Query string
	PrevForm int
	NextFrom int
	Items []interface{}
}
