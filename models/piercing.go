package models

// Piercing is an artwork model of gogo.tattoo. it represents Piercing
type Piercing struct {
	Artwork
}

// NewPiercing returns a new henna model, requires id, the unique title of the new work
// link, also unique and final image ipfs hash
func NewPiercing(id, title, link, hash string) (h Piercing) {
	h.ID = id
	h.Link = link
	h.ImageIpfs = hash
	return
}
