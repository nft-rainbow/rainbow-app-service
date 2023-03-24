package services

import (
	"context"

	"github.com/nft-rainbow/rainbow-app-service/models"
)

type (
	SocialToolBot interface {
		SendChannelMessage(ctx context.Context, channedId string, msg string) error
		SendDirectMessage(ctx context.Context, userId string, msg string) error
	}
)

type VerifyUserResponse struct {
	Code string
}

type InsertProjectorReq struct {
	models.SocialToolProjecter
	AuthCode string `json:"auth_code"`
}
