package orgunit

import (
	"context"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/wingfeng/idxadmin/internal/model"
	"github.com/wingfeng/idxadmin/service"
)

func Test_Tree(t *testing.T) {
	s := service.OrgUnit()
	r, err := s.Tree(context.Background())
	t.Logf("err %v", err)
	gtest.AssertNil(err)
	gtest.AssertGT(len(r), 0)

	t.Logf("tree:%v", r)
}

func TestPage(t *testing.T) {
	s := service.OrgUnit()
	r, err := s.List(context.Background(), &model.PageReq{
		Page:     1,
		PageSize: 10,
	})
	gtest.AssertNil(err)
	t.Logf("page:%v", r)
}
