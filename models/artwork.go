package models

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
