package server

type OAuthCallbackRequest struct {
	Code   string `form:"code" binding:"required"`
	ShopID int64  `form:"shop_id" binding:"required"`
}
