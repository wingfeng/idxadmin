package user

import (
	"context"
	"strconv"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/test/gtest"
	v1 "github.com/wingfeng/idxadmin/api/user/v1"
	"github.com/wingfeng/idxadmin/internal/service"
)

func TestResetPwd(t *testing.T) {
	ctx := context.Background()
	id := int64(1838872840128958464)
	s := service.User()
	idStr := strconv.Itoa(int(id))
	u, _ := s.Get(ctx, id)
	oldPwd := u.PasswordHash
	t.Log(oldPwd)
	pwd, err := s.ResetPwd(ctx, &v1.ResetPwdReq{Id: idStr})
	gtest.AssertNil(err)
	t.Log(pwd)
	u, _ = s.Get(ctx, id)
	newPwd := u.PasswordHash
	t.Logf("%v", newPwd)
	gtest.AssertNE(oldPwd, newPwd)
}
