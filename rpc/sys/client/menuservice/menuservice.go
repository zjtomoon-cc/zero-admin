// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package menuservice

import (
	"context"

	"zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ConfigAddReq          = sysclient.ConfigAddReq
	ConfigAddResp         = sysclient.ConfigAddResp
	ConfigDeleteReq       = sysclient.ConfigDeleteReq
	ConfigDeleteResp      = sysclient.ConfigDeleteResp
	ConfigListData        = sysclient.ConfigListData
	ConfigListReq         = sysclient.ConfigListReq
	ConfigListResp        = sysclient.ConfigListResp
	ConfigUpdateReq       = sysclient.ConfigUpdateReq
	ConfigUpdateResp      = sysclient.ConfigUpdateResp
	DeptAddReq            = sysclient.DeptAddReq
	DeptAddResp           = sysclient.DeptAddResp
	DeptDeleteReq         = sysclient.DeptDeleteReq
	DeptDeleteResp        = sysclient.DeptDeleteResp
	DeptListData          = sysclient.DeptListData
	DeptListReq           = sysclient.DeptListReq
	DeptListResp          = sysclient.DeptListResp
	DeptUpdateReq         = sysclient.DeptUpdateReq
	DeptUpdateResp        = sysclient.DeptUpdateResp
	DictAddReq            = sysclient.DictAddReq
	DictAddResp           = sysclient.DictAddResp
	DictDeleteReq         = sysclient.DictDeleteReq
	DictDeleteResp        = sysclient.DictDeleteResp
	DictListData          = sysclient.DictListData
	DictListReq           = sysclient.DictListReq
	DictListResp          = sysclient.DictListResp
	DictUpdateReq         = sysclient.DictUpdateReq
	DictUpdateResp        = sysclient.DictUpdateResp
	InfoReq               = sysclient.InfoReq
	InfoResp              = sysclient.InfoResp
	JobAddReq             = sysclient.JobAddReq
	JobAddResp            = sysclient.JobAddResp
	JobDeleteReq          = sysclient.JobDeleteReq
	JobDeleteResp         = sysclient.JobDeleteResp
	JobListData           = sysclient.JobListData
	JobListReq            = sysclient.JobListReq
	JobListResp           = sysclient.JobListResp
	JobUpdateReq          = sysclient.JobUpdateReq
	JobUpdateResp         = sysclient.JobUpdateResp
	LoginLogAddReq        = sysclient.LoginLogAddReq
	LoginLogAddResp       = sysclient.LoginLogAddResp
	LoginLogDeleteReq     = sysclient.LoginLogDeleteReq
	LoginLogDeleteResp    = sysclient.LoginLogDeleteResp
	LoginLogListData      = sysclient.LoginLogListData
	LoginLogListReq       = sysclient.LoginLogListReq
	LoginLogListResp      = sysclient.LoginLogListResp
	LoginReq              = sysclient.LoginReq
	LoginResp             = sysclient.LoginResp
	MenuAddReq            = sysclient.MenuAddReq
	MenuAddResp           = sysclient.MenuAddResp
	MenuDeleteReq         = sysclient.MenuDeleteReq
	MenuDeleteResp        = sysclient.MenuDeleteResp
	MenuListData          = sysclient.MenuListData
	MenuListReq           = sysclient.MenuListReq
	MenuListResp          = sysclient.MenuListResp
	MenuListTree          = sysclient.MenuListTree
	MenuUpdateReq         = sysclient.MenuUpdateReq
	MenuUpdateResp        = sysclient.MenuUpdateResp
	QueryMenuByRoleIdReq  = sysclient.QueryMenuByRoleIdReq
	QueryMenuByRoleIdResp = sysclient.QueryMenuByRoleIdResp
	ReSetPasswordReq      = sysclient.ReSetPasswordReq
	ReSetPasswordResp     = sysclient.ReSetPasswordResp
	RoleAddReq            = sysclient.RoleAddReq
	RoleAddResp           = sysclient.RoleAddResp
	RoleDeleteReq         = sysclient.RoleDeleteReq
	RoleDeleteResp        = sysclient.RoleDeleteResp
	RoleListData          = sysclient.RoleListData
	RoleListReq           = sysclient.RoleListReq
	RoleListResp          = sysclient.RoleListResp
	RoleUpdateReq         = sysclient.RoleUpdateReq
	RoleUpdateResp        = sysclient.RoleUpdateResp
	SysLogAddReq          = sysclient.SysLogAddReq
	SysLogAddResp         = sysclient.SysLogAddResp
	SysLogDeleteReq       = sysclient.SysLogDeleteReq
	SysLogDeleteResp      = sysclient.SysLogDeleteResp
	SysLogListData        = sysclient.SysLogListData
	SysLogListReq         = sysclient.SysLogListReq
	SysLogListResp        = sysclient.SysLogListResp
	UpdateConfigRoleReq   = sysclient.UpdateConfigRoleReq
	UpdateConfigRoleResp  = sysclient.UpdateConfigRoleResp
	UpdateMenuRoleReq     = sysclient.UpdateMenuRoleReq
	UpdateMenuRoleResp    = sysclient.UpdateMenuRoleResp
	UserAddReq            = sysclient.UserAddReq
	UserAddResp           = sysclient.UserAddResp
	UserDeleteReq         = sysclient.UserDeleteReq
	UserDeleteResp        = sysclient.UserDeleteResp
	UserListData          = sysclient.UserListData
	UserListReq           = sysclient.UserListReq
	UserListResp          = sysclient.UserListResp
	UserStatusReq         = sysclient.UserStatusReq
	UserStatusResp        = sysclient.UserStatusResp
	UserUpdateReq         = sysclient.UserUpdateReq
	UserUpdateResp        = sysclient.UserUpdateResp

	MenuService interface {
		MenuAdd(ctx context.Context, in *MenuAddReq, opts ...grpc.CallOption) (*MenuAddResp, error)
		MenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error)
		MenuUpdate(ctx context.Context, in *MenuUpdateReq, opts ...grpc.CallOption) (*MenuUpdateResp, error)
		MenuDelete(ctx context.Context, in *MenuDeleteReq, opts ...grpc.CallOption) (*MenuDeleteResp, error)
	}

	defaultMenuService struct {
		cli zrpc.Client
	}
)

func NewMenuService(cli zrpc.Client) MenuService {
	return &defaultMenuService{
		cli: cli,
	}
}

func (m *defaultMenuService) MenuAdd(ctx context.Context, in *MenuAddReq, opts ...grpc.CallOption) (*MenuAddResp, error) {
	client := sysclient.NewMenuServiceClient(m.cli.Conn())
	return client.MenuAdd(ctx, in, opts...)
}

func (m *defaultMenuService) MenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error) {
	client := sysclient.NewMenuServiceClient(m.cli.Conn())
	return client.MenuList(ctx, in, opts...)
}

func (m *defaultMenuService) MenuUpdate(ctx context.Context, in *MenuUpdateReq, opts ...grpc.CallOption) (*MenuUpdateResp, error) {
	client := sysclient.NewMenuServiceClient(m.cli.Conn())
	return client.MenuUpdate(ctx, in, opts...)
}

func (m *defaultMenuService) MenuDelete(ctx context.Context, in *MenuDeleteReq, opts ...grpc.CallOption) (*MenuDeleteResp, error) {
	client := sysclient.NewMenuServiceClient(m.cli.Conn())
	return client.MenuDelete(ctx, in, opts...)
}
