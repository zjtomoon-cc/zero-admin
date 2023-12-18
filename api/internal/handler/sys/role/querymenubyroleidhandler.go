package role

import (
	"net/http"
	"zero-admin/api/internal/logic/sys/role"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryMenuByRoleIdHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleMenuReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewQueryMenuByRoleIdLogic(r.Context(), ctx)
		resp, err := l.QueryMenuByRoleId(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
