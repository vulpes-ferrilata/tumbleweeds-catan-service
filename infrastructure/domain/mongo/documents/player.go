package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
	Document         `bson:",inline"`
	UserID           primitive.ObjectID `bson:"user_id"`
	TurnOrder        int                `bson:"turn_order"`
	Color            string             `bson:"color"`
	IsOffered        bool               `bson:"is_offered"`
	Score            int                `bson:"score"`
	Achievements     []*Achievement     `bson:"achievements"`
	ResourceCards    []*ResourceCard    `bson:"resource_cards"`
	DevelopmentCards []*DevelopmentCard `bson:"development_cards"`
	Constructions    []*Construction    `bson:"constructions"`
	Roads            []*Road            `bson:"roads"`
}
