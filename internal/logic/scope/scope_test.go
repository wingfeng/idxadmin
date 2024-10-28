package scope

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/wingfeng/idxadmin/internal/service"
)

func TestDelet(t *testing.T) {
	ctx := context.Background()
	s := service.Scope()
	err := s.Delete(ctx, 1)
	gtest.AssertNil(err)
}
