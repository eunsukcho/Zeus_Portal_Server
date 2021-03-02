package keycloak

func GetAccessToken() string {
	h_accessToken := InitAccessTokenInfo()
	r_accessToken := InitResponseAccessTokenInfo()

	res := h_accessToken.request_AccessToken()
	access_token := r_accessToken.parsing_ResponseBodyJson(res)
	return access_token
}
