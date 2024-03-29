package models

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/app_errors"
	"github.com/vulpes-ferrilata/slices"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type resourceDiscardPhase struct {
	game *Game
}

func (r resourceDiscardPhase) buildSettlementAndRoad(userID primitive.ObjectID, landID primitive.ObjectID, pathID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceDiscardPhase)
}

func (r resourceDiscardPhase) rollDices(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceDiscardPhase)
}

func (r resourceDiscardPhase) discardResourceCards(userID primitive.ObjectID, resourceCardIDs []primitive.ObjectID) error {
	player, err := slices.Find(func(player *Player) (bool, error) {
		return player.userID == userID, nil
	}, r.game.getAllPlayers()...)
	if errors.Is(err, slices.ErrNoMatchFound) {
		return errors.WithStack(app_errors.ErrPlayerNotFound)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	if player.discardedResources {
		return errors.WithStack(app_errors.ErrYouAlreadyDiscardedResourceCards)
	}

	player.discardedResources = true

	if len(player.resourceCards) < 8 {
		return errors.WithStack(app_errors.ErrYouHaveNoNeedToDiscardResourceCards)
	}

	if len(resourceCardIDs) != len(player.resourceCards)/2 {
		return errors.WithStack(app_errors.ErrSelectedResourceCardsMustBeEqualsToHalfOfYourCurrentlyResourceCards)
	}

	selectedResourceCards, err := slices.Map(func(resourceCardID primitive.ObjectID) (*ResourceCard, error) {
		resourceCard, err := slices.Find(func(resourceCard *ResourceCard) (bool, error) {
			return resourceCard.id == resourceCardID, nil
		}, player.resourceCards...)
		if errors.Is(err, slices.ErrNoMatchFound) {
			return nil, errors.WithStack(app_errors.ErrResourceCardNotFound)
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}

		return resourceCard, nil
	}, resourceCardIDs...)
	if err != nil {
		return errors.WithStack(err)
	}

	player.resourceCards = slices.Remove(player.resourceCards, selectedResourceCards...)
	r.game.resourceCards = append(r.game.resourceCards, selectedResourceCards...)

	isAllPlayerDiscardedResources, err := slices.All(func(player *Player) (bool, error) {
		return player.discardedResources || len(player.resourceCards) < 8, nil
	}, r.game.getAllPlayers()...)
	if err != nil {
		return errors.WithStack(err)
	}
	if isAllPlayerDiscardedResources {
		r.game.phase = Robbing
	}

	return nil
}

func (r resourceDiscardPhase) moveRobber(userID primitive.ObjectID, terrainID primitive.ObjectID, playerID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceDiscardPhase)
}

func (r resourceDiscardPhase) endTurn(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceDiscardPhase)
}

func (r resourceDiscardPhase) buildSettlement(userID primitive.ObjectID, landID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceDiscardPhase)
}

func (r resourceDiscardPhase) buildRoad(userID primitive.ObjectID, pathID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceDiscardPhase)
}

func (r resourceDiscardPhase) upgradeCity(userID primitive.ObjectID, constructionID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceDiscardPhase)
}

func (r resourceDiscardPhase) buyDevelopmentCard(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceDiscardPhase)
}

func (r resourceDiscardPhase) toggleResourceCards(userID primitive.ObjectID, resourceCardID []primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceDiscardPhase)
}

func (r resourceDiscardPhase) maritimeTrade(userID primitive.ObjectID, resourceCardType ResourceCardType, demandingResourceCardType ResourceCardType) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceDiscardPhase)
}

func (r resourceDiscardPhase) sendTradeOffer(userID primitive.ObjectID, playerID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceDiscardPhase)
}

func (r resourceDiscardPhase) confirmTradeOffer(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceDiscardPhase)
}

func (r resourceDiscardPhase) cancelTradeOffer(userID primitive.ObjectID) error {
	return errors.WithStack(app_errors.ErrYouAreUnableToPerformThisActionInResourceDiscardPhase)
}
