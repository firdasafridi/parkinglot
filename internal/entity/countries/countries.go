package countries

type ResponseCountry struct {
	Name         Name                   `json:"name"`
	Tld          []string               `json:"tld"`
	Cca2         string                 `json:"cca2"`
	Ccn3         string                 `json:"ccn3"`
	Cca3         string                 `json:"cca3"`
	Cioc         string                 `json:"cioc"`
	Independent  bool                   `json:"independent"`
	Status       string                 `json:"status"`
	Unmember     bool                   `json:"unMember"`
	Currencies   Currencies             `json:"currencies"`
	Idd          Idd                    `json:"idd"`
	Capital      []string               `json:"capital"`
	Altspellings []string               `json:"altSpellings"`
	Region       string                 `json:"region"`
	Subregion    string                 `json:"subregion"`
	Languages    Languages              `json:"languages"`
	Translations map[string]Translation `json:"translations"`
	Latlng       []float64              `json:"latlng"`
	Landlocked   bool                   `json:"landlocked"`
	Borders      []string               `json:"borders"`
	Area         float64                `json:"area"`
	Demonyms     Demonyms               `json:"demonyms"`
	Flag         string                 `json:"flag"`
	Maps         Maps                   `json:"maps"`
	Population   int                    `json:"population"`
	Gini         Gini                   `json:"gini"`
	Fifa         string                 `json:"fifa"`
	Car          Car                    `json:"car"`
	Timezones    []string               `json:"timezones"`
	Continents   []string               `json:"continents"`
	Flags        Flags                  `json:"flags"`
	Coatofarms   Coatofarms             `json:"coatOfArms"`
	Startofweek  string                 `json:"startOfWeek"`
	Capitalinfo  Capitalinfo            `json:"capitalInfo"`
	Postalcode   Postalcode             `json:"postalCode"`
}

type Nativename struct {
	Ind Translation `json:"ind"`
}

type Name struct {
	Common     string     `json:"common"`
	Official   string     `json:"official"`
	Nativename Nativename `json:"nativeName"`
}

type Idr struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type Currencies struct {
	Idr Idr `json:"IDR"`
}

type Idd struct {
	Root     string   `json:"root"`
	Suffixes []string `json:"suffixes"`
}

type Languages struct {
	Ind string `json:"ind"`
}

type Translation struct {
	Official string `json:"official"`
	Common   string `json:"common"`
}

type Eng struct {
	F string `json:"f"`
	M string `json:"m"`
}

type Fra struct {
	F string `json:"f"`
	M string `json:"m"`
}

type Demonyms struct {
	Eng Eng `json:"eng"`
	Fra Fra `json:"fra"`
}

type Maps struct {
	Googlemaps     string `json:"googleMaps"`
	Openstreetmaps string `json:"openStreetMaps"`
}

type Gini struct {
	Num2019 float64 `json:"2019"`
}

type Car struct {
	Signs []string `json:"signs"`
	Side  string   `json:"side"`
}

type Flags struct {
	Png string `json:"png"`
	Svg string `json:"svg"`
}

type Coatofarms struct {
	Png string `json:"png"`
	Svg string `json:"svg"`
}

type Capitalinfo struct {
	Latlng []float64 `json:"latlng"`
}

type Postalcode struct {
	Format string `json:"format"`
	Regex  string `json:"regex"`
}
