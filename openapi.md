---
接口文档

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

## put__nft-entity

> Code samples

`PUT /nft-entity`

> Body parameter

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

<h3 id="put__nft-entity-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[cnft.apiv1.NftEntityUpdateReq](#schemacnft.apiv1.nftentityupdatereq)|true|none|

> Example responses

> 200 Response

```json
{}
```

<h3 id="put__nft-entity-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|[cnft.apiv1.NftEntityUpdateRes](#schemacnft.apiv1.nftentityupdateres)|

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

## delete__nft-entity_{id}

> Code samples

`DELETE /nft-entity/{id}`

<h3 id="delete__nft-entity_{id}-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|Id|path|number(int)|true|none|

> Example responses

> 200 Response

```json
{}
```

<h3 id="delete__nft-entity_{id}-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|[cnft.apiv1.NftEntityDeleteRes](#schemacnft.apiv1.nftentitydeleteres)|

<aside class="success">
This operation does not require authentication
</aside>

## get__nft-entity_{id}

> Code samples

`GET /nft-entity/{id}`

<h3 id="get__nft-entity_{id}-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|Id|path|number(int)|true|none|

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

<h3 id="get__nft-entity_{id}-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|[cnft.apiv1.NftEntityGetByIdRes](#schemacnft.apiv1.nftentitygetbyidres)|

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

<h2 id="tocS_cnft.apiv1.NftEntityUpdateReq">cnft.apiv1.NftEntityUpdateReq</h2>
<!-- backwards compatibility -->
<a id="schemacnft.apiv1.nftentityupdatereq"></a>
<a id="schema_cnft.apiv1.NftEntityUpdateReq"></a>
<a id="tocScnft.apiv1.nftentityupdatereq"></a>
<a id="tocscnft.apiv1.nftentityupdatereq"></a>

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

<h2 id="tocS_cnft.apiv1.NftEntityUpdateRes">cnft.apiv1.NftEntityUpdateRes</h2>
<!-- backwards compatibility -->
<a id="schemacnft.apiv1.nftentityupdateres"></a>
<a id="schema_cnft.apiv1.NftEntityUpdateRes"></a>
<a id="tocScnft.apiv1.nftentityupdateres"></a>
<a id="tocscnft.apiv1.nftentityupdateres"></a>

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

<h2 id="tocS_cnft.apiv1.NftEntityDeleteReq">cnft.apiv1.NftEntityDeleteReq</h2>
<!-- backwards compatibility -->
<a id="schemacnft.apiv1.nftentitydeletereq"></a>
<a id="schema_cnft.apiv1.NftEntityDeleteReq"></a>
<a id="tocScnft.apiv1.nftentitydeletereq"></a>
<a id="tocscnft.apiv1.nftentitydeletereq"></a>

```json
{
  "Id": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|Id|number(int)|false|none|none|

<h2 id="tocS_cnft.apiv1.NftEntityDeleteRes">cnft.apiv1.NftEntityDeleteRes</h2>
<!-- backwards compatibility -->
<a id="schemacnft.apiv1.nftentitydeleteres"></a>
<a id="schema_cnft.apiv1.NftEntityDeleteRes"></a>
<a id="tocScnft.apiv1.nftentitydeleteres"></a>
<a id="tocscnft.apiv1.nftentitydeleteres"></a>

```json
{}

```

### Properties

*None*

<h2 id="tocS_cnft.apiv1.NftEntityGetByIdReq">cnft.apiv1.NftEntityGetByIdReq</h2>
<!-- backwards compatibility -->
<a id="schemacnft.apiv1.nftentitygetbyidreq"></a>
<a id="schema_cnft.apiv1.NftEntityGetByIdReq"></a>
<a id="tocScnft.apiv1.nftentitygetbyidreq"></a>
<a id="tocscnft.apiv1.nftentitygetbyidreq"></a>

```json
{
  "Id": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|Id|number(int)|false|none|none|

