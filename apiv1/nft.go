package apiv1

import (
	"cnft/internal/model"
	"cnft/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type NftEntityAddReq struct {
	g.Meta   `path:"/nft-entity" tags:"nft-entity" method:"post" dc:"上链"`
	EntityId string `json:"entity_id" v:"required"` // 实体id##
}

type NftEntityAddRes struct {
}

type NftEntityUpdateReq struct {
	g.Meta `path:"/nft-entity" tags:"nft-entity" method:"put" summary:""`
	entity.NftEntity
}

type NftEntityUpdateRes struct {
}

type NftEntityListReq struct {
	g.Meta `path:"/nft-entity/list" tags:"nft-entity" dc:"查询已上链的entity信息" method:"get" summary:""`
	model.NftEntityListInputQuery
	model.NftEntityListInputPage
}

type NftEntityListRes struct {
	List  []entity.NftEntity
	Total int
}

type NftEntityGetByIdReq struct {
	g.Meta `path:"/nft-entity/{id}" tags:"nft-entity" method:"get" summary:""`
	Id     int
}

type NftEntityGetByIdRes struct {
	entity.NftEntity
}

type NftEntityGetByEntityIdReq struct {
	g.Meta   `path:"/nft-entity/entity/{id}" tags:"nft-entity" dc:"根据entityId查询链上信息" method:"get" summary:""`
	EntityId int `json:"entity_id"`
}

type NftEntityGetByEntityIdRes struct {
	entity.NftEntity
}

type NftEntityDeleteReq struct {
	g.Meta `path:"/nft-entity/{id}" tags:"nft-entity" method:"delete" summary:""`
	Id     int
}

type NftEntityDeleteRes struct {
}
