package entity

type Crime struct {
	CrimeID     int    `json:"ID"`
	HeroID      int    `json:"HeroID"`
	VillainID   int    `json:"VillainID"`
	Description string `json:"Description"`
	CrimeTime   string `json:"CrimeTime"`
}

type ErrorMessage struct {
	Message string
}
