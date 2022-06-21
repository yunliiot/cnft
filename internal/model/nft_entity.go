package model

import (
	"cnft/internal/model/entity"
)

type NftEntityAddInput struct {
	entity.NftEntity
}

type NftEntityUpdateInput struct {
	entity.NftEntity
}

// NftEntityListInputQuery 添加自定义的检索字段
type NftEntityListInputQuery struct{}
type NftEntityListInputPage struct {
	PageNum  int
	PageSize int
}

type NftEntityListInput struct {
	NftEntityListInputQuery
	NftEntityListInputPage
}

type NftEntityListOutput struct {
	List  []entity.NftEntity `json:"list"`
	Total int                `json:"total"`
}

type NftEntityGetByIdInput struct {
}

type NftEntityGetByIdOutput struct {
	entity.NftEntity
}
