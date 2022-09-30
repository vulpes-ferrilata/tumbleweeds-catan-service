package commands

type BuildSettlement struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
	LandID string `validate:"required,objectid"`
}
