package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type phase interface {
	buildSettlementAndRoad(userID primitive.ObjectID, landID primitive.ObjectID, pathID primitive.ObjectID) error
	rollDices(userID primitive.ObjectID) error
	discardResourceCards(userID primitive.ObjectID, resourceCardIDs []primitive.ObjectID) error
	moveRobber(userID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error
	endTurn(userID primitive.ObjectID) error
	buildSettlement(userID primitive.ObjectID, landID primitive.ObjectID) error
	buildRoad(userID primitive.ObjectID, pathID primitive.ObjectID) error
	upgradeCity(userID primitive.ObjectID, constructionID primitive.ObjectID) error
	buyDevelopmentCard(userID primitive.ObjectID) error
	toggleResourceCards(userID primitive.ObjectID, resourceCardIDs []primitive.ObjectID) error
	maritimeTrade(userID primitive.ObjectID, resourceCardType ResourceCardType, demandingResourceCardType ResourceCardType) error
	sendTradeOffer(userID primitive.ObjectID, playerID primitive.ObjectID) error
	confirmTradeOffer(userID primitive.ObjectID) error
	cancelTradeOffer(userID primitive.ObjectID) error
}
