package models

// Artist is a struct for storing a general information of the artists,
// as well as some necessary info, it'll expand
type Artist struct {
	Link     string   `json:"link" toml:"link"`
	Name     string   `json:"name" toml:"name"`
	Services []string `json:"services,omitempty" toml:"services"`

	JoinDate      string `json:"join_date" toml:"join_date"`
	BirthDate     string `json:"birth_date" toml:"birth_date"`
	Experience    string `json:"experience" toml:"experience"`
	About         string `json:"about" toml:"about"`
	Origin        string `json:"origin" toml:"origin"`
	LocationNow   string `json:"location_now" toml:"location_now"`
	AvatarIpfs    string `json:"avatar_ipfs" toml:"avatar_ipfs"`
	CurrentStudio string `json:"current_studio" toml:"current_studio"`
}

// Artists is plural for Artist, supports sort.Sort interface
type Artists []Artist

func (slice Artists) Len() int {
	return len(slice)
}

func (slice Artists) Less(i, j int) bool {
	return slice[i].JoinDate < slice[j].JoinDate
}

func (slice Artists) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
