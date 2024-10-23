package common

import (
	"context"

	"github.com/wingfeng/idx/utils"
	"github.com/wingfeng/idxadmin/internal/model"
	"github.com/wingfeng/idxadmin/internal/service"
)

type sUtils struct{}

func init() {
	service.RegisterUtils(NewUtils())
}
func NewUtils() service.IUtils {
	return &sUtils{}
}
func (s *sUtils) NewId(ctx context.Context) model.ID {
	id := utils.GeneratID()
	return model.ID(id.Int64())
}
