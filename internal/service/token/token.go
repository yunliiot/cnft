package token

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
)

type (
	// sToken is service struct of module NftEntity.
	sToken struct{}
)

var (
	// insToken is the instance of service User.
	insToken = sToken{}
)

// Token returns the interface of Token service.
func Token() *sToken {
	return &insToken
}

func (s *sToken) Generate(ctx context.Context, id string) string {
	hash, _ := gmd5.EncryptString(id)
	return hash
}
