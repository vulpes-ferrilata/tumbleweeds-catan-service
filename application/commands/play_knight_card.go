package commands

type PlayKnightCard struct {
	GameID    string `validate:"required,objectid"`
	UserID    string `validate:"required,objectid"`
	TerrainID string `validate:"required,objectid"`
	PlayerID  string `validate:"omitempty,objectid"`
}