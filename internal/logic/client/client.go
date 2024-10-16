package client

import (
	"context"
	"fmt"

	"log/slog"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/wingfeng/idx-oauth2/utils"
	v1 "github.com/wingfeng/idxadmin/api/client/v1"
	"github.com/wingfeng/idxadmin/internal/dao"
	"github.com/wingfeng/idxadmin/internal/model/do"
	"github.com/wingfeng/idxadmin/internal/model/entity"
	"github.com/wingfeng/idxadmin/internal/service"
)

type (
	sClient struct{}
)

func init() {
	service.RegisterClient(New())
}
func New() service.IClient {
	slog.Info("Client service register")
	return &sClient{}
}

// Get one client by id.
func (s *sClient) Get(ctx context.Context, id int) (*do.Clients, error) {

	if row, err := dao.Clients.Ctx(ctx).One("id=?", id); err != nil {
		return nil, err
	} else {
		result := &do.Clients{}
		err = row.Struct(result)
		return result, err
	}

}
func (s *sClient) List(ctx context.Context, req v1.PageReq) (*v1.PageRes, error) {
	g.Log().Info(ctx, "req", req)
	items := make([]entity.Clients, 0)
	count := 0
	err := dao.Clients.Ctx(ctx).Order(req.SortField, req.GetSortOrder()).ScanAndCount(&items, &count, true)
	res := &v1.PageRes{}
	res.PageSize = req.PageSize
	res.Page = req.Page
	res.Total = count
	res.List = items
	return res, err
}
func (s *sClient) Save(ctx context.Context, req v1.SaveReq) (err error) {

	result, err := dao.Clients.Ctx(ctx).OnConflict("id").Save(req.Clients)
	if ra, er := result.RowsAffected(); ra == 0 || er != nil {
		return fmt.Errorf("save error ,no rows affected %s", er)
	}
	return err
}
func (s *sClient) Delete(ctx context.Context, req *v1.DeleteReq) (err error) {
	_, err = dao.Clients.Ctx(ctx).Delete("id=?", req.Id)
	return err
}
func (s *sClient) GenerateSecret(ctx context.Context, req *v1.GenerateSecretReq) (string, error) {
	count, err := dao.ClientSecrets.Ctx(ctx).Where("client_id=?", req.ClientId).Count()
	if err != nil {
		return "", err
	}
	if count > 2 {
		return "", fmt.Errorf("client %d already has 3 secrets, please delete some of them first", req.ClientId)
	}

	secret := utils.GenerateRandomString(32)
	newHash, _ := utils.HashPassword(secret)
	// newSecrets := &struct {
	// 	ClientId   int64       `json:"clientId" form:"clientId"`
	// 	Value      string      `json:"value" form:"value"`
	// 	Name       string      `json:"name" form:"name"`
	// 	Expiration *gtime.Time `json:"expiration" form:"expiration"`
	// }{
	// 	ClientId:   req.ClientId,
	// 	Value:      newHash,
	// 	Name:       req.Name,
	// 	Expiration: gtime.New(req.Expiration),
	// }
	newSecrets := &entity.ClientSecrets{
		ClientId:   req.ClientId,
		Value:      newHash,
		Name:       req.Name,
		Expiration: gtime.New(req.Expiration),
	}
	_, err = dao.ClientSecrets.Ctx(ctx).OmitEmptyData().InsertAndGetId(newSecrets)
	if err != nil {
		return "", err
	}
	return secret, err
}

func (s *sClient) GetSecrets(ctx context.Context, req *v1.GetClientSecretsReq) ([]*entity.ClientSecrets, error) {
	items := make([]*entity.ClientSecrets, 0)
	err := dao.ClientSecrets.Ctx(ctx).Where("client_id=?", req.ClientId).Scan(&items)
	return items, err
}
func (s *sClient) DeleteSecret(ctx context.Context, req *v1.DeleteSecretReq) error {
	_, err := dao.ClientSecrets.Ctx(ctx).Delete("id=?", req.Id)
	return err
}