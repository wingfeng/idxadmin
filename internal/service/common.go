package service

import (
	"context"

	"github.com/wingfeng/idxadmin/internal/model"
)

type (
	ICRUD[entity interface{}] interface {
		Get(ctx context.Context, id interface{}) (entity, error)
		Delete(ctx context.Context, id interface{}) error
		Save(ctx context.Context, req entity) (err error)
		List(ctx context.Context, req *model.PageReq) (*model.PageRes, error)
	}
)

func CRUD[do interface{}]() ICRUD[do] {
	return nil
}

type (
	IUtils interface {
		NewId(ctx context.Context) model.ID
	}
)

var (
	localUtils IUtils
)

func Utils() IUtils {
	if localUtils == nil {
		panic("implement not found for interface IUtils, forgot register?")
	}
	return localUtils
}

func RegisterUtils(i IUtils) {
	localUtils = i
}
