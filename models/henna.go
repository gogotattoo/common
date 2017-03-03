package models

// Henna is an artwork model of gogo.tattoo. it represents henna
type Henna struct {
	Artwork
}

// NewHenna returns a new henna model, requires id, the unique title of the new work
// link, also unique and final image ipfs hash
func NewHenna(id, title, link, hash string) (h Henna) {
	h.ID = id
	h.Link = link
	h.ImageIpfs = hash
	return
}
