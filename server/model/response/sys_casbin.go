package response

import "gin-element-admin/model/request"

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
