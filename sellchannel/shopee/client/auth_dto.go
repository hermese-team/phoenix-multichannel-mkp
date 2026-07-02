package client

type GetAccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpireIn     int64  `json:"expire_in"`
	ShopID       int64  `json:"shop_id"`
	Error        string `json:"error"`
	Message      string `json:"message"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpireIn     int64  `json:"expire_in"`
	ShopID       int64  `json:"shop_id"`
	Error        string `json:"error"`
	Message      string `json:"message"`
}
