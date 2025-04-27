package api

import (
	"github.com/upils/game-api/model"
	"gorm.io/gorm"
)

type gameServer struct {
	db *gorm.DB
}

func NewGameServer(db *gorm.DB) gameServer {
	s := gameServer{db: db}
	s.db.AutoMigrate(&model.Game{}, &model.Platform{})

	platforms := []*model.Platform{
		{Name: "PC"},
		{Name: "PS4"},
		{Name: "PS5"},
		{Name: "Switch"},
		{Name: "One"},
		{Name: "WiiU"},
	}

	result := s.db.Create(platforms)
	if result.Error != nil {
		panic(result.Error)
	}
	return s
}
