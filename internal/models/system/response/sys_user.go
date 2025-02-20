package response

import "github.com/unicrm/server/internal/models/system"

type SysUserResponse struct {
	User system.SysUser `json:"user"`
}

type LoginResponse struct {
	User  system.SysUser `json:"user"`
	Token string         `json:"token"`
}
