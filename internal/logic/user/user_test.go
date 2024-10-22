package user

import (
	"context"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	v1 "github.com/wingfeng/idxadmin/api/user/v1"
	"github.com/wingfeng/idxadmin/internal/service"
)

func TestResetPwd(t *testing.T) {
	ctx := context.Background()
	id := int64(1838872840128958465)
	s := service.User()

	u, _ := s.Get(ctx, id)
	oldPwd := u.PasswordHash.(*g.Var)
	t.Log(oldPwd)
	s.ResetPwd(ctx, v1.ResetPwdReq{Id: "1838872840128958465"})
	u, _ = s.Get(ctx, id)
	newPwd := u.PasswordHash.(*g.Var)
	t.Logf("%v", newPwd)
	gtest.AssertNE(oldPwd, newPwd)
}
