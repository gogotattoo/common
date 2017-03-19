package models

// Artist is a struct for storing a general information of the artists,
// as well as some necessary info, it'll expand
type Artist struct {
	Link     string   `json:"link,omitempty" toml:"link"`
	Name     string   `json:"name,omitempty" toml:"name"`
	Services []string `json:"services,omitempty" toml:"services"`

	JoinDate  string `json:"join_date,omitempty" toml:"join_date"`
	BirthDate string `json:"birth_date,omitempty" toml:"birth_date"`
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
