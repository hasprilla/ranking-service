package models

type ArtistRanking struct {
	ArtistID     uint   `json:"artistId"`
	Name         string `json:"name"`
	CurrentStars int    `json:"currentStars"`
	ImageURL     string `json:"imageUrl"`
}

type FanRanking struct {
	UserID     uint    `json:"userId"`
	Name       string  `json:"name"`
	TotalStars float64 `json:"totalStars"`
	ImageURL   string  `json:"imageUrl"`
}
