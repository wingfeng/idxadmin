package common

import (
	"context"

	v1 "github.com/wingfeng/idxadmin/api/common/v1"
	"github.com/wingfeng/idxadmin/internal/service"
)

func (c *ControllerV1) NewId(ctx context.Context, req *v1.NewIdReq) (res *v1.NewIdRes, err error) {
	id := service.Utils().NewId(ctx)

	return &v1.NewIdRes{Id: id.String()}, nil
}
