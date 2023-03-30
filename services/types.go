package services

type VerifyUserResponse struct {
	AuthCode string `json:"auth_code"`
}

type InsertProjectorReq struct {
	UserSocialId string  `json:"user_social_id" binding:"required"`
	AuthCode     *string `json:"auth_code"`
}
