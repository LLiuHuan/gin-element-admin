package response

import "gin-element-admin/model"

type LoginResponse struct {
	User         model.SysUser `json:"user"`
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	ExpiresAt    int64         `json:"expiresAt"`
}
