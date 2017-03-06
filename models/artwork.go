package models

// Artwork is a struct for storing a general information for a piece of art,
// as well as the processes of this artwork being done (images_ipfs)
type Artwork struct {
	ID    string `json:"id,omitempty"`
	Link  string `json:"link,omitempty"`
	Title string `json:"title,omitempty" toml:"title"`
	//MadeLocation Address  `json:"made_at"`
	DurationMin int    `json:"duration_min,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Extra       string `json:"extra,omitempty"`
	Article     string `json:"article,omitempty"`

	MadeDate        string   `json:"made_date,omitempty" toml:"made_date"`
	PublishDate     string   `json:"date,omitempty"`
	Tags            []string `json:"tags,omitempty"`
	BodyParts       []string `json:"bodypart,omitempty"`
	ImageIpfs       string   `json:"image_ipfs" toml:"image_ipfs"`
	ImagesIpfs      []string `json:"images_ipfs,omitempty" toml:"images_ipfs"`
	LocationCity    string   `json:"made_at_city" toml:"location_city"`
	LocationCountry string   `json:"made_at_country" toml:"location_country"`
	MadeAtShop      string   `json:"made_at_shop,omitempty" toml:"made_at_shop"`
}

// Artworks is plural for Artwork, supports sort.Sort interface
type Artworks []Artwork

func (slice Artworks) Len() int {
	return len(slice)
}

func (slice Artworks) Less(i, j int) bool {
	return slice[i].MadeDate < slice[j].MadeDate
}

func (slice Artworks) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
