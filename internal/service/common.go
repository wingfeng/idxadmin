package service

import (
	"context"

	"github.com/wingfeng/idxadmin/internal/model"
)

type (
	ICRUD[do interface{}] interface {
		Get(ctx context.Context, id interface{}) (do, error)
		Delete(ctx context.Context, id interface{}) error
		Save(ctx context.Context, req do) (err error)
		List(ctx context.Context, req *model.PageReq) (*model.PageRes, error)
	}
)

func CRUD[do interface{}]() ICRUD[do] {
	return nil
}
