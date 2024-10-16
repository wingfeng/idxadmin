// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package client

import (
	"context"

	"github.com/wingfeng/idxadmin/api/client/v1"
)

type IClientV1 interface {
	Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error)
	Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error)
	Save(ctx context.Context, req *v1.SaveReq) (res *v1.SaveRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	GenerateSecret(ctx context.Context, req *v1.GenerateSecretReq) (res *v1.GenerateSecretRes, err error)
	DeleteSecret(ctx context.Context, req *v1.DeleteSecretReq) (res *v1.DeleteSecretRes, err error)
	GetClientSecrets(ctx context.Context, req *v1.GetClientSecretsReq) (res *v1.GetClientSecretsRes, err error)
}
