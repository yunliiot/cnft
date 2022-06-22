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
[接口文档](/docs/openapi.md)

---

