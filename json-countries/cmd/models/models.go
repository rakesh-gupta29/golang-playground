package models

type Country struct {
	Name         Name                       `json:"name"`
	Independent  bool                       `json:"independent"`
	Status       string                     `json:"status"`
	UnMember     bool                       `json:"unMember"`
	Currencies   map[string]CurrencyType    `json:"currencies,omitempty"`
	Capital      []string                   `json:"capital,omitempty"`
	Cca2         string                     `json:"cca2"`
	Ccn3         string                     `json:"ccn3"`
	Cca3         string                     `json:"cca3"`
	Cioc         string                     `json:"cioc"`
	AltSpellings []string                   `json:"altSpellings,omitempty"`
	Region       string                     `json:"region,omitempty"`
	Subregion    string                     `json:"subregion,omitempty"`
	Languages    map[string]string          `json:"languages,omitempty"`
	Translations map[string]TranslationType `json:"translations,omitempty"`
	LatLng       []uint8                    `json:"latlng,omitempty"`
	Landlocked   bool                       `json:"landlocked"`
	Borders      []string                   `json:"borders,omitempty"`
	Area         float64                    `json:"area"`
	Demonyms     map[string]Demonym         `json:"demonyms,omitempty"`
	Flag         Flags                      `json:"flag,omitempty"`
	Maps         map[string]string          `json:"maps,omitempty"`
	Population   int                        `json:"population"`
	Fifa         string                     `json:"fifa"`
	Car          Car                        `json:"car"`
	Timezones    []string                   `json:"timezones,omitempty"`
	Continents   []string                   `json:"continents,omitempty"`
	Flags        map[string]string          `json:"flags,omitempty"`
	CoatOfArms   CoatOfArms                 `json:"coatOfArms,omitempty"`
	StartOfWeek  string                     `json:"startOfWeek,omitempty"`
	CapitalInfo  map[string][]float64       `json:"capitalInfo,omitempty"`
	PostalCode   PostalCode                 `json:"postalCode,omitempty"`
}

type CurrencyType struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type TranslationType struct {
	Common   string `json:"common"`
	Official string `json:"official"`
}

type Name struct {
	Common     string               `json:"common"`
	Official   string               `json:"official"`
	NativeName map[string]NameTypes `json:"nativeName"`
}

type NameTypes struct {
	Official string `json:"official"`
	Common   string `json:"common"`
}

type Demonym struct {
	F string `json:"f"`
	M string `json:"m"`
}

type Maps struct {
	GoogleMaps     string `json:"googleMaps"`
	OpenStreetMaps string `json:"openStreetMaps"`
}

type Flags struct {
	PNG string `json:"png"`
	SVG string `json:"svg"`
	Alt string `json:"alt"`
}

type CoatOfArms struct {
	PNG string `json:"png"`
	SVG string `json:"svg"`
}

type PostalCode struct {
	Format string `json:"format"`
	Regex  string `json:"regex"`
}

type Car struct {
	Signs []string `json:"signs"`
	Side  string   `json:"side"`
}

type ValidationError struct {
	IsOK    bool   `json:"isOK"`
	Message string `json:"message"`
}
