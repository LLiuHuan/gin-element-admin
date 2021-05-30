package response

import "gin-element-admin/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
