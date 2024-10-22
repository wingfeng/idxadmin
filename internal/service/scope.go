// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	v1 "github.com/wingfeng/idxadmin/api/orgunit/v1"
)

type (
	IScope interface {
		Get(ctx context.Context, _ v1.GetReq)
	}
)

var (
	localScope IScope
)

func Scope() IScope {
	if localScope == nil {
		panic("implement not found for interface IScope, forgot register?")
	}
	return localScope
}

func RegisterScope(i IScope) {
	localScope = i
}
