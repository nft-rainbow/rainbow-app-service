package services

import (
	"context"
	"errors"
	"sync"

	"github.com/nft-rainbow/rainbow-app-service/models"
)

type (
	SocialToolBot interface {
		SendChannelMessage(ctx context.Context, channedId string, msg string) error
		SendDirectMessage(ctx context.Context, userId string, msg string) error
	}
)

var (
	dodoBot           *DodoBot
	dodoBotCreateOnce sync.Once
)

func getSocialToolBot(socialToolType models.SocialToolType) (SocialToolBot, error) {
	switch socialToolType {
	case models.SOCIAL_TOOL_DODO:
		dodoBotCreateOnce.Do(func() {
			dodoBot = NewDodoBot()
		})
		return dodoBot, nil
	}
	return nil, errors.New("unsupported social tool")
}
