## CNFT

将实体信息（游戏装备、图片等）上链，并提供相关查询、交易接口

## 准备

### 启动 mysql

```
docker run -itd --name mysql-cnft -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql
## 初始化数据库
docker exec -it mysql-cnft bash
mysql -u root -p
## 输入密码
## 将manifest/config/init.sql内容粘贴执行
```

### 启动 chanin33

```
git clone git@github.com:33cn/chain33.git
## 编译前准备
export GOPROXY=https://mirrors.aliyun.com/goproxy
make
cd chain33/cmd/chain3
go build
./chain33 -f chain33.toml
```

### 使用 chanin33 cli工具，解锁钱包
chain33钱包默认状态是lock，需要使用命令解锁之后才能创建账户
```
## 切换到chain33代码根目录
cd chain33/cmd/cli
go build
## 生成seed，注意：生成的seed要记住，如果不小心误删钱包，通过seed可以找回钱包
./cli seed generate -l 0
## 保存seed，-p后跟的是钱包密码。 注意：密码可以自定义，并且牢牢记住，后面解锁钱包时会用到。密码需满足8位字母数字组合的要求。
./cli seed save -s [上一步生成的seed值] -p tech1234
## 解锁钱包
./cli wallet unlock -p tech1234 -t 0 
```

### 启动 cnft

```
# 上链
curl --location --request POST 'http://127.0.0.1:8000/nft-entity' \
--header 'Content-Type: application/json' \
--header 'Accept: application/json' \
--data-raw '{
  "entity_id": "test"
}'

# 使用cnft接口查询上链信息
curl --location --request GET 'http://127.0.0.1:8000/nft-entity/list' \
--header 'Accept: application/json'

# 使用chain33 cli工具查询
./cli account list
```

---
## 接口文档

---


## post__nft-entity

> Code samples

`POST /nft-entity`

上链

> Body parameter

```json
{
  "entity_id": "string"
}
```

<h3 id="post__nft-entity-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[cnft.apiv1.NftEntityAddReq](#schemacnft.apiv1.nftentityaddreq)|true|none|

> Example responses

> 200 Response

```json
{}
```

<h3 id="post__nft-entity-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|[cnft.apiv1.NftEntityAddRes](#schemacnft.apiv1.nftentityaddres)|

<aside class="success">
This operation does not require authentication
</aside>

## get__nft-entity_entity_{id}

> Code samples

`GET /nft-entity/entity/{id}`

根据entityId查询链上信息

<h3 id="get__nft-entity_entity_{id}-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|entity_id|query|number(int)|false|none|

> Example responses

> 200 Response

```json
{
  "id": 0,
  "entity_id": "string",
  "entity_type": "string",
  "token": "string",
  "created_at": 0,
  "updated_at": 0,
  "deleted_at": 0
}
```

<h3 id="get__nft-entity_entity_{id}-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|[cnft.apiv1.NftEntityGetByIdRes](#schemacnft.apiv1.nftentitygetbyidres)|

<aside class="success">
This operation does not require authentication
</aside>

## get__nft-entity_list

> Code samples

`GET /nft-entity/list`

查询已上链的entity信息

<h3 id="get__nft-entity_list-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|PageNum|query|number(int)|false|none|
|PageSize|query|number(int)|false|none|

> Example responses

> 200 Response

```json
{
  "List": [
    {
      "id": 0,
      "entity_id": "string",
      "entity_type": "string",
      "token": "string",
      "created_at": 0,
      "updated_at": 0,
      "deleted_at": 0
    }
  ],
  "Total": 0
}
```

<h3 id="get__nft-entity_list-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|[cnft.apiv1.NftEntityListRes](#schemacnft.apiv1.nftentitylistres)|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_cnft.apiv1.NftEntityAddReq">cnft.apiv1.NftEntityAddReq</h2>
<!-- backwards compatibility -->
<a id="schemacnft.apiv1.nftentityaddreq"></a>
<a id="schema_cnft.apiv1.NftEntityAddReq"></a>
<a id="tocScnft.apiv1.nftentityaddreq"></a>
<a id="tocscnft.apiv1.nftentityaddreq"></a>

```json
{
  "entity_id": "string"
}

```

上链

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|entity_id|string(string)|true|none|none|

<h2 id="tocS_cnft.apiv1.NftEntityAddRes">cnft.apiv1.NftEntityAddRes</h2>
<!-- backwards compatibility -->
<a id="schemacnft.apiv1.nftentityaddres"></a>
<a id="schema_cnft.apiv1.NftEntityAddRes"></a>
<a id="tocScnft.apiv1.nftentityaddres"></a>
<a id="tocscnft.apiv1.nftentityaddres"></a>

```json
{}

```

### Properties

*None*

<h2 id="tocS_cnft.apiv1.NftEntityGetByEntityIdReq">cnft.apiv1.NftEntityGetByEntityIdReq</h2>
<!-- backwards compatibility -->
<a id="schemacnft.apiv1.nftentitygetbyentityidreq"></a>
<a id="schema_cnft.apiv1.NftEntityGetByEntityIdReq"></a>
<a id="tocScnft.apiv1.nftentitygetbyentityidreq"></a>
<a id="tocscnft.apiv1.nftentitygetbyentityidreq"></a>

```json
{
  "entity_id": 0
}

```

根据entityId查询链上信息

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|entity_id|number(int)|false|none|none|

<h2 id="tocS_cnft.apiv1.NftEntityGetByIdRes">cnft.apiv1.NftEntityGetByIdRes</h2>
<!-- backwards compatibility -->
<a id="schemacnft.apiv1.nftentitygetbyidres"></a>
<a id="schema_cnft.apiv1.NftEntityGetByIdRes"></a>
<a id="tocScnft.apiv1.nftentitygetbyidres"></a>
<a id="tocscnft.apiv1.nftentitygetbyidres"></a>

```json
{
  "id": 0,
  "entity_id": "string",
  "entity_type": "string",
  "token": "string",
  "created_at": 0,
  "updated_at": 0,
  "deleted_at": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|number(int64)|false|none|主键|
|entity_id|string(string)|false|none|实体id##|
|entity_type|string(string)|false|none|实体类型，1:装备##|
|token|string(string)|false|none|区块链token##|
|created_at|number(int64)|false|none|创建时间##|
|updated_at|number(int64)|false|none|创建时间##|
|deleted_at|number(int64)|false|none|创建时间##|

<h2 id="tocS_cnft.apiv1.NftEntityListReq">cnft.apiv1.NftEntityListReq</h2>
<!-- backwards compatibility -->
<a id="schemacnft.apiv1.nftentitylistreq"></a>
<a id="schema_cnft.apiv1.NftEntityListReq"></a>
<a id="tocScnft.apiv1.nftentitylistreq"></a>
<a id="tocscnft.apiv1.nftentitylistreq"></a>

```json
{
  "PageNum": 0,
  "PageSize": 0
}

```

查询已上链的entity信息

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|PageNum|number(int)|false|none|none|
|PageSize|number(int)|false|none|none|

<h2 id="tocS_cnft.apiv1.NftEntityListRes">cnft.apiv1.NftEntityListRes</h2>
<!-- backwards compatibility -->
<a id="schemacnft.apiv1.nftentitylistres"></a>
<a id="schema_cnft.apiv1.NftEntityListRes"></a>
<a id="tocScnft.apiv1.nftentitylistres"></a>
<a id="tocscnft.apiv1.nftentitylistres"></a>

```json
{
  "List": [
    {
      "id": 0,
      "entity_id": "string",
      "entity_type": "string",
      "token": "string",
      "created_at": 0,
      "updated_at": 0,
      "deleted_at": 0
    }
  ],
  "Total": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|List|[[cnft.internal.model.entity.NftEntity](#schemacnft.internal.model.entity.nftentity)]|false|none|none|
|Total|number(int)|false|none|none|

<h2 id="tocS_cnft.internal.model.entity.NftEntity">cnft.internal.model.entity.NftEntity</h2>
<!-- backwards compatibility -->
<a id="schemacnft.internal.model.entity.nftentity"></a>
<a id="schema_cnft.internal.model.entity.NftEntity"></a>
<a id="tocScnft.internal.model.entity.nftentity"></a>
<a id="tocscnft.internal.model.entity.nftentity"></a>

```json
{
  "id": 0,
  "entity_id": "string",
  "entity_type": "string",
  "token": "string",
  "created_at": 0,
  "updated_at": 0,
  "deleted_at": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|number(int64)|false|none|主键|
|entity_id|string(string)|false|none|实体id##|
|entity_type|string(string)|false|none|实体类型，1:装备##|
|token|string(string)|false|none|区块链token##|
|created_at|number(int64)|false|none|创建时间##|
|updated_at|number(int64)|false|none|创建时间##|
|deleted_at|number(int64)|false|none|创建时间##|
