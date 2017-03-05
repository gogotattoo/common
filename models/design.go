package models

// Design is an artwork model of gogo.tattoo. it represents design
type Design struct {
	Artwork
}

// NewDesign returns a new design model, requires id, the unique title of the new work
// link, also unique and final image ipfs hash
func NewDesign(id, title, link, hash string) (h Design) {
	h.ID = id
	h.Link = link
	h.ImageIpfs = hash
	return
}
