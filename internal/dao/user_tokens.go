// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/wingfeng/idxadmin/internal/dao/internal"
)

// internalUserTokensDao is internal type for wrapping internal DAO implements.
type internalUserTokensDao = *internal.UserTokensDao

// userTokensDao is the data access object for table user_tokens.
// You can define custom methods on it to extend its functionality as you wish.
type userTokensDao struct {
	internalUserTokensDao
}

var (
	// UserTokens is globally public accessible object for table user_tokens operations.
	UserTokens = userTokensDao{
		internal.NewUserTokensDao(),
	}
)

// Fill with you ideas below.
