package role

import (
	"context"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/test/gtest"
	v1 "github.com/wingfeng/idxadmin/api/role/v1"
	"github.com/wingfeng/idxadmin/internal/service"
)

func TestMembers(t *testing.T) {
	s := service.Role()
	req := &v1.MembersReq{Id: 1838872840128958465}
	req.Page = 1
	req.PageSize = 10
	res, err := s.Members(context.Background(), req)
	gtest.AssertNil(err)
	gtest.Assert(res.Total, 1)
	t.Logf("%v", res)
}
