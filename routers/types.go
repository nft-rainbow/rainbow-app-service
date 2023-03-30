package routers

type VerifySocialUserReq struct {
	UserSocialId string `json:"user_social_id" binding:"required"`
}
