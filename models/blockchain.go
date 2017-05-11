package models

// Blockchain is a struct to save references, where works are stored on the Blockchain
// TODO: Refactor, better naming convention, etc
type Blockchain struct {
	steem string `json:"steem,omitempty" toml:"steem"`
	golos string `json:"golos,omitempty" toml:"golos"`
}
