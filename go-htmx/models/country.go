package models

type Country struct {
	Name struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"name"`

	CCA2  string `json:"cca2"`
	Flags struct {
		PNG string `json:"png"`
		SVG string `json:"svg"`
	} `json:"flags"`
	Region  string   `json:"region"`
	Capital []string `json:"capital"`
}
