package models

// Tattoo represents a tattoo work, can keep a reference to the design it was based on
type Tattoo struct {
	Artwork

	Design *Design
}

// Tattoos represents array of tattoos
type Tattoos []Tattoo

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
