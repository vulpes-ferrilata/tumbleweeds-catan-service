package app_errors

var (
	ErrGameNotFound                                                        = NewNotFoundError("game")
	ErrStaleGame                                                           = NewStaleObjectError("game")
	ErrPlayerNotFound                                                      = NewNotFoundError("player")
	ErrTerrainNotFound                                                     = NewNotFoundError("terrain")
	ErrLandNotFound                                                        = NewNotFoundError("land")
	ErrPathNotFound                                                        = NewNotFoundError("path")
	ErrConstructionNotFound                                                = NewNotFoundError("construction")
	ErrResourceCardNotFound                                                = NewNotFoundError("resource-card")
	ErrDevelopmentCardNotFound                                             = NewNotFoundError("development-card")
	ErrAchievementCardNotFound                                             = NewNotFoundError("achievement-card")
	ErrGameHasFullPlayers                                                  = NewBusinessRuleError("game-has-full-players")
	ErrYouAlreadyJoinedTheGame                                             = NewBusinessRuleError("you-already-joined-the-game")
	ErrYouAreNotInTurn                                                     = NewBusinessRuleError("you-are-not-in-turn")
	ErrGameMustHaveAtLeastTwoPlayers                                       = NewBusinessRuleError("game-must-have-at-least-two-players")
	ErrGameHasNotStartedYet                                                = NewBusinessRuleError("game-has-not-started-yet")
	ErrGameAlreadyStarted                                                  = NewBusinessRuleError("game-already-started")
	ErrGameAlreadyFinished                                                 = NewBusinessRuleError("game-already-finished")
	ErrYouAreUnableToPerformThisActionInSetupPhase                         = NewBusinessRuleError("you-are-unable-to-perform-this-action-in-setup-phase")
	ErrYouAreUnableToPerformThisActionInResourceProductionPhase            = NewBusinessRuleError("you-are-unable-to-perform-this-action-in-resource-production-phase")
	ErrYouAreUnableToPerformThisActionInResourceDiscardPhase               = NewBusinessRuleError("you-are-unable-to-perform-this-action-in-resource-discard-phase")
	ErrYouAreUnableToPerformThisActionInRobbingPhase                       = NewBusinessRuleError("you-are-unable-to-perform-this-action-in-robbing-phase")
	ErrYouAreUnableToPerformThisActionInResourceConsumptionPhase           = NewBusinessRuleError("you-are-unable-to-perform-this-action-in-resource-consumption-phase")
	ErrSelectedLandMustBeAdjacentToYourRoad                                = NewBusinessRuleError("selected-land-must-be-adjacent-to-your-road")
	ErrSelectedPathMustBeAdjacentToYourConstructionOrRoad                  = NewBusinessRuleError("selected-path-must-be-adjacent-to-your-construction-or-road")
	ErrSelectedPathPassThroughConstructionOfOtherPlayer                    = NewBusinessRuleError("selected-path-pass-through-construction-of-other-player")
	ErrNearbyLandsMustBeVacant                                             = NewBusinessRuleError("nearby-lands-must-be-vacant")
	ErrSelectedLandAndPathMustBeAdjacent                                   = NewBusinessRuleError("selected-land-and-path-must-be-adjacent")
	ErrYouRunOutOfRoads                                                    = NewBusinessRuleError("you-run-out-of-roads")
	ErrYouRunOutOfSettlements                                              = NewBusinessRuleError("you-run-out-of-settlements")
	ErrYouRunOutOfCities                                                   = NewBusinessRuleError("you-run-out-of-cities")
	ErrYouHaveInsufficientResourceCards                                    = NewBusinessRuleError("you-have-insufficient-resource-cards")
	ErrRobberMustBeMovedToOtherTerrain                                     = NewBusinessRuleError("robber-must-be-moved-to-other-terrain")
	ErrSelectedPlayerMustHaveConstructionNextToRobber                      = NewBusinessRuleError("selected-player-must-have-construction-next-to-robber")
	ErrYouMustRobPlayerWhoHasConstructionNextToRobber                      = NewBusinessRuleError("you-must-rob-player-who-has-construction-next-to-robber")
	ErrSelectedConstructionAlreadyUpgraded                                 = NewBusinessRuleError("selected-construction-already-upgraded")
	ErrSelectedConstructionDoesNotBelongToAnyLand                          = NewBusinessRuleError("selected-construction-does-not-belong-to-any-land")
	ErrGameRunOutOfDevelopmentCards                                        = NewBusinessRuleError("game-run-out-of-development-cards")
	ErrGameHasInsufficientResourceCards                                    = NewBusinessRuleError("game-has-insufficient-resource-cards")
	ErrYouAlreadyOfferedThisPlayer                                         = NewBusinessRuleError("you-already-offered-this-player")
	ErrYouHaveNotReceivedAnyOffer                                          = NewBusinessRuleError("you-have-not-received-any-offer")
	ErrYouMustOfferAtLeastOneResourceCard                                  = NewBusinessRuleError("you-must-offer-at-least-one-resource-card")
	ErrSelectedPlayerMustOfferAtLeastOneResourceCard                       = NewBusinessRuleError("selected-player-must-offer-at-least-one-resource-card")
	ErrActivePlayerMustOfferAtLeastOneResourceCard                         = NewBusinessRuleError("active-player-must-offer-at-least-one-resource-card")
	ErrYouAlreadyDiscardedResourceCards                                    = NewBusinessRuleError("you-already-discarded-resource-cards")
	ErrYouHaveNoNeedToDiscardResourceCards                                 = NewBusinessRuleError("you-have-no-need-to-discard-resource-cards")
	ErrSelectedResourceCardsMustBeEqualsToHalfOfYourCurrentlyResourceCards = NewBusinessRuleError("selected-resource-cards-must-be-equals-to-half-of-your-currently-resource-cards")
	ErrSelectedDevelopmentCardIsUnavailableToUse                           = NewBusinessRuleError("selected-development-card-is-unavailable-to-use")
	ErrSelectedDevelopmentCardMustBeKnightCard                             = NewBusinessRuleError("selected-development-card-must-be-knight-card")
	ErrSelectedDevelopmentCardMustBeRoadBuildingCard                       = NewBusinessRuleError("selected-development-card-must-be-road-building-card")
	ErrSelectedDevelopmentCardMustBeYearOfPlentyCard                       = NewBusinessRuleError("selected-development-card-must-be-year-of-plenty-card")
	ErrSelectedDevelopmentCardMustBeMonopolyCard                           = NewBusinessRuleError("selected-development-card-must-be-monopoly-card")
	ErrSelectedDevelopmentCardMustBeVictoryPointCard                       = NewBusinessRuleError("selected-development-card-must-be-victory-point-card")
)
