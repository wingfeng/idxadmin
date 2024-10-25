package common

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/wingfeng/idx/utils"
	"github.com/wingfeng/idxadmin/internal/model"
	"github.com/wingfeng/idxadmin/internal/model/do"
	"github.com/wingfeng/idxadmin/internal/model/entity"
)

func TestCrud(t *testing.T) {
	ctx := context.Background()

	id := utils.GeneratID().Int64()
	t.Logf("New ID: %v", id)
	uName := gofakeit.Username()
	email := gofakeit.Email()
	u := &entity.Users{
		UserName: uName,

		Email: email,

		DisplayName: gofakeit.Name(),

		OuId:           1328680589330485249,
		Ou:             "XX软件有限公司",
		EmailConfirmed: true,
	}
	u.LockoutEnd = gtime.Now()
	u.Id = entity.ID(id)

	s := NewCRUD[*entity.Users]()
	err := s.Save(ctx, u)
	gtest.AssertNil(err)
	r, err := s.Get(ctx, id)
	gtest.AssertNil(err)
	gtest.Assert(r.UserName, uName)
	t.Logf("User %v", r)
	err = s.Delete(ctx, id)
	gtest.AssertNil(err)
}

func TestPage(t *testing.T) {
	ctx := context.Background()
	s := NewCRUD[do.Users]()
	req := &model.PageReq{Page: 1, PageSize: 10,
		Filters:   []string{"account like ?"},
		Args:      []string{"%admin%"},
		SortOrder: "asc",
		SortField: "id",
	}

	res, err := s.List(ctx, req)
	gtest.Assert(1, res.Total)
	gtest.AssertNil(err)
	t.Logf("Page %v", res)
	//clear filter and test pagesize
	req.Filters = []string{}
	req.PageSize = 2
	req.Page = 3
	req.Fields = "id,account,email,display_name"
	res, err = s.List(ctx, req)
	gtest.AssertNil(err)
	gtest.Assert(res.Total, 14)
	gtest.Assert(res.Page, 3)
	t.Logf("Page %v", res.List)
}
