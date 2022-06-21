package nft_entity

import (
	"cnft/internal/service/internal/do"
	"context"
	"errors"

	"cnft/internal/model"
	"cnft/internal/service/internal/dao"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	// sNftEntity is service struct of module NftEntity.
	sNftEntity struct{}
)

var (
	// insNftEntity is the instance of service User.
	insNftEntity = sNftEntity{}
)

// NftEntity returns the interface of NftEntity service.
func NftEntity() *sNftEntity {
	return &insNftEntity
}

func (s *sNftEntity) Add(ctx context.Context, input *model.NftEntityAddInput) error {
	input.Id = 0
	ct, _ := dao.NftEntity.Ctx(ctx).Where(do.NftEntity{EntityId: input.EntityId}).Count()
	if ct > 0 {
		return errors.New("data already exists")
	}
	_, err := dao.NftEntity.Ctx(ctx).OmitEmpty().Insert(input)
	if err != nil {
	}
	return err
}

func (s *sNftEntity) Update(ctx context.Context, input *model.NftEntityUpdateInput) error {
	_, err := dao.NftEntity.Ctx(ctx).Save(input)
	if err != nil {
	}
	return err
}

func (s *sNftEntity) List(ctx context.Context, req *model.NftEntityListInput) (ret model.NftEntityListOutput, err error) {
	count, err := dao.NftEntity.Ctx(ctx).OmitEmpty().Where(req.NftEntityListInputQuery).Count()
	if err != nil {
		return ret, err
	}

	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	offset := (req.PageNum - 1) * req.PageSize
	err = dao.NftEntity.Ctx(ctx).OmitEmpty().Where(req.NftEntityListInputQuery).
		Offset(offset).Limit(req.PageSize).
		Scan(&ret.List)
	if err != nil {
		return ret, err
	}
	ret.Total = count
	return ret, nil
}

func (s *sNftEntity) GetById(ctx context.Context, id int) (ret model.NftEntityGetByIdOutput, err error) {
	err = dao.NftEntity.Ctx(ctx).Where(g.Map{"id": id}).Scan(&ret)
	return ret, err
}

func (s *sNftEntity) GetByEntityId(ctx context.Context, id int) (ret model.NftEntityGetByIdOutput, err error) {
	err = dao.NftEntity.Ctx(ctx).Where(g.Map{"entity_id": id}).Scan(&ret)
	return ret, err
}

func (s *sNftEntity) Delete(ctx context.Context, id int) (err error) {
	_, err = dao.NftEntity.Ctx(ctx).Where(g.Map{"id": id}).Delete()
	return err
}
