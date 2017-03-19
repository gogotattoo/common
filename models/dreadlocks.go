package models

// Dreadlocks is an artwork model of gogo.tattoo. it represents Dreadlocks
type Dreadlocks struct {
	Artwork
}

// NewDreadlocks returns a new model, requires id, the unique title of the new work
// link, also unique and final image ipfs hash
func NewDreadlocks(id, title, link, hash string) (h Dreadlocks) {
	h.ID = id
	h.Link = link
	h.ImageIpfs = hash
	return
}
