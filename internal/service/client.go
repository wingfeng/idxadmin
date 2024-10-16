// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	v1 "github.com/wingfeng/idxadmin/api/client/v1"
	"github.com/wingfeng/idxadmin/internal/model/do"
	"github.com/wingfeng/idxadmin/internal/model/entity"
)

type (
	IClient interface {
		// Get one client by id.
		Get(ctx context.Context, id int) (*do.Clients, error)
		List(ctx context.Context, req v1.PageReq) (*v1.PageRes, error)
		Save(ctx context.Context, req v1.SaveReq) (err error)
		Delete(ctx context.Context, req *v1.DeleteReq) (err error)
		GenerateSecret(ctx context.Context, req *v1.GenerateSecretReq) (string, error)
		GetSecrets(ctx context.Context, req *v1.GetClientSecretsReq) ([]*entity.ClientSecrets, error)
		DeleteSecret(ctx context.Context, req *v1.DeleteSecretReq) error
	}
)

var (
	localClient IClient
)

func Client() IClient {
	if localClient == nil {
		panic("implement not found for interface IClient, forgot register?")
	}
	return localClient
}

func RegisterClient(i IClient) {
	localClient = i
}
