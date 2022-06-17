package dyflexis

type Offices []Office

type Office struct {
}

type Informationstream struct {
	Day struct {
		Value int `json:"value"`
		Hours []struct {
			Hour  int `json:"hour"`
			Value int `json:"value"`
		} `json:"hours"`
	} `json:"day"`
}
