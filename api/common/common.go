// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package common

import (
	"context"

	"github.com/wingfeng/idxadmin/api/common/v1"
)

type ICommonV1 interface {
	NewId(ctx context.Context, req *v1.NewIdReq) (res *v1.NewIdRes, err error)
}
