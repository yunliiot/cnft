package token

import (
	"cnft/internal/config"
	"context"
	"fmt"
	"github.com/33cn/chain33-sdk-go/client"
	"github.com/33cn/chain33/rpc/jsonclient"
	"github.com/33cn/chain33/types"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/os/glog"
)

var cClient *CClient

type CClient struct {
	*client.JSONClient
	URL string
}

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
	if cClient == nil {
		hash, _ := gmd5.EncryptString(id)
		return hash
	}
	return s.GenerateV2(ctx, id)
}

func (s *sToken) GenerateV2(ctx context.Context, id string) string {
	res, err := cClient.AccountCreate(ctx, fmt.Sprintf("game_%s", id), 0)
	if err != nil {
		return ""
	}
	return res.Acc.Addr
}

func InitJSONClient() {
	url := config.Config.AppConfigServer.Chain33Url
	cClient, _ = NewJSONClient(url)
}

func NewJSONClient(url string) (*CClient, error) {
	cli, err := client.NewJSONClient("", url)
	if err != nil {
		glog.Warning(context.Background(), "NewJSONClient failed", err)
		return nil, err
	}
	c := CClient{}
	c.JSONClient = cli
	c.URL = url
	return &c, nil
}

func (client *CClient) AccountCreate(ctx context.Context, label string, addressType int32) (*types.WalletAccount, error) {
	params := types.ReqNewAccount{
		Label:     label,
		AddressID: addressType,
	}
	var res types.WalletAccount
	rCtx := jsonclient.NewRPCCtx(client.URL, "Chain33.NewAccount", &params, &res)
	_, err := rCtx.RunResult()
	if err != nil {
		glog.Warning(ctx, "NewJSONClient failed", err)
		return nil, err
	}
	return &res, nil
}
