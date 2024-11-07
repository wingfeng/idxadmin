package common

import (
	"context"

	"github.com/wingfeng/idx/utils"
	"github.com/wingfeng/idxadmin/internal/model/entity"
	"github.com/wingfeng/idxadmin/service"
)

type sUtils struct{}

func init() {
	service.RegisterUtils(NewUtils())
}
func NewUtils() service.IUtils {
	return &sUtils{}
}
func (s *sUtils) NewId(ctx context.Context) entity.ID {
	id := utils.GeneratID()
	return entity.ID(id.Int64())
}
