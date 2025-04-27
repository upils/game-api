package model

type Platform struct {
	Name  string  `json:",inline" gorm:"primaryKey"`
	Games []*Game `json:"-" gorm:"many2many:game_languages;"`
}

type Game struct {
	ID          uint   `json:"-" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"unique"`
	ReleaseDate string `json:"release_date"`
	Studio      string `json:"studio"`
	Ratings     uint
	Platforms []*Platform `json:"platforms" gorm:"many2many:game_platforms;"`
}

