package models

type Tattoo struct {
	Artwork
}

// Address stores the location information where the work was made
// type Address struct {
// 	City    string `json:"city,omitempty"`
// 	Country string `json:"country,omitempty"`
// 	Shop    string `json:"shop,omitempty"`
// }

// NewTattoo returns a new tattoo , requires id, the unique title of the new work
// link, also unique and final image ipfs hash
func NewTattoo(id, title, link, hash string) (t Tattoo) {
	t.ID = id
	t.Link = link
	t.ImageIpfs = hash
	return
}