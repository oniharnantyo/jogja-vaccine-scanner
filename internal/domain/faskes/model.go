package faskes

type Faskes struct {
	ID     int    `json:"id"`
	Name   string `json:"faskes"`
	Region string `json:"wilayah"`
	Year   int    `json:"tahun"`
}

type Faskeses []Faskes
