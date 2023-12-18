// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package userservice

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

	UserService interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		UserInfo(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error)
		UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddResp, error)
		UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error)
		UserUpdate(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*UserUpdateResp, error)
		UserDelete(ctx context.Context, in *UserDeleteReq, opts ...grpc.CallOption) (*UserDeleteResp, error)
		ReSetPassword(ctx context.Context, in *ReSetPasswordReq, opts ...grpc.CallOption) (*ReSetPasswordResp, error)
		UpdateUserStatus(ctx context.Context, in *UserStatusReq, opts ...grpc.CallOption) (*UserStatusResp, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

func (m *defaultUserService) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := sysclient.NewUserServiceClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUserService) UserInfo(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error) {
	client := sysclient.NewUserServiceClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}

func (m *defaultUserService) UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddResp, error) {
	client := sysclient.NewUserServiceClient(m.cli.Conn())
	return client.UserAdd(ctx, in, opts...)
}

func (m *defaultUserService) UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error) {
	client := sysclient.NewUserServiceClient(m.cli.Conn())
	return client.UserList(ctx, in, opts...)
}

func (m *defaultUserService) UserUpdate(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*UserUpdateResp, error) {
	client := sysclient.NewUserServiceClient(m.cli.Conn())
	return client.UserUpdate(ctx, in, opts...)
}

func (m *defaultUserService) UserDelete(ctx context.Context, in *UserDeleteReq, opts ...grpc.CallOption) (*UserDeleteResp, error) {
	client := sysclient.NewUserServiceClient(m.cli.Conn())
	return client.UserDelete(ctx, in, opts...)
}

func (m *defaultUserService) ReSetPassword(ctx context.Context, in *ReSetPasswordReq, opts ...grpc.CallOption) (*ReSetPasswordResp, error) {
	client := sysclient.NewUserServiceClient(m.cli.Conn())
	return client.ReSetPassword(ctx, in, opts...)
}

func (m *defaultUserService) UpdateUserStatus(ctx context.Context, in *UserStatusReq, opts ...grpc.CallOption) (*UserStatusResp, error) {
	client := sysclient.NewUserServiceClient(m.cli.Conn())
	return client.UpdateUserStatus(ctx, in, opts...)
}
