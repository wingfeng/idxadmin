package scope

import (
	"context"

	v1 "github.com/wingfeng/idxadmin/api/orgunit/v1"
)

type sScope struct {
}

func (s *sScope) Get(ctx context.Context, _ v1.GetReq) {}
