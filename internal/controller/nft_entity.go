package controller

import (
	"cnft/apiv1"
	"cnft/internal/model"
	"cnft/internal/service/nft_entity"
	"cnft/internal/service/token"
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/os/gtime"
)

var (
	NftEntity = cNftEntity{}
)

type cNftEntity struct{}

func (c *cNftEntity) Add(ctx context.Context, req *apiv1.NftEntityAddReq) (res *apiv1.NftEntityAddRes, err error) {
	in := &model.NftEntityAddInput{}
	in.EntityId = req.EntityId
	in.EntityType = "1"
	in.Token = token.Token().Generate(ctx, in.EntityId)
	in.CreatedAt = gtime.TimestampMilli()
	in.UpdatedAt = in.CreatedAt
	err = nft_entity.NftEntity().Add(ctx, in)
	return
}

func (c *cNftEntity) Update(ctx context.Context, req *apiv1.NftEntityUpdateReq) (res *apiv1.NftEntityUpdateRes, err error) {
	in := &model.NftEntityUpdateInput{}
	in.NftEntity = req.NftEntity
	err = nft_entity.NftEntity().Update(ctx, in)
	return
}

func (c *cNftEntity) NftEntityList(ctx context.Context, req *apiv1.NftEntityListReq) (res apiv1.NftEntityListRes, err error) {
	in := &model.NftEntityListInput{}
	in.NftEntityListInputQuery = req.NftEntityListInputQuery
	in.NftEntityListInputPage = req.NftEntityListInputPage
	list, err := nft_entity.NftEntity().List(ctx, in)
	if err != nil {
		return res, err
	}
	res.List = list.List
	res.Total = list.Total
	return res, err
}

func (c *cNftEntity) NftEntityGetById(ctx context.Context, req *apiv1.NftEntityGetByIdReq) (res *apiv1.NftEntityGetByIdRes, err error) {
	info, err := nft_entity.NftEntity().GetById(ctx, req.Id)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	res = &apiv1.NftEntityGetByIdRes{}
	res.NftEntity = info.NftEntity
	return
}

func (c *cNftEntity) NftEntityGetByEntityId(ctx context.Context, req *apiv1.NftEntityGetByEntityIdReq) (res *apiv1.NftEntityGetByIdRes, err error) {
	info, err := nft_entity.NftEntity().GetByEntityId(ctx, req.EntityId)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	res = &apiv1.NftEntityGetByIdRes{}
	res.NftEntity = info.NftEntity
	return
}

func (c *cNftEntity) Delete(ctx context.Context, req *apiv1.NftEntityDeleteReq) (res apiv1.NftEntityDeleteRes, err error) {
	err = nft_entity.NftEntity().Delete(ctx, req.Id)
	return res, err
}
