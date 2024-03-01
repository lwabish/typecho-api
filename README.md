# Typecho-api

- 提供http api，直接对接操作typecho数据库
- 目前仅实现了创建/更新文章接口，对接typora的导出功能即可实现typora写完文章后直接发布到typecho

## Roadmap

- [ ] typora对接api方案和实现
- [x] 创建文章
- [x] 更新文章
- [x] 处理分类和标签
- [ ] 入库时正文添加markdown wrapper
- [ ] 文章字数更新
- [ ] 分类标签计数更新
- [ ] 高级功能
  - [ ] 自动添加TOC
  - [ ] 自动添加文章摘要（调用AI服务生成）
  - [ ] 自动添加文章封面（调用AI服务生成）
  - [ ] 自动拼写检查（调用AI服务实现）
  - [ ] 自动添加文末原文链接，防采集
  - [ ] 鉴权
- [ ] 验证：在本地对接typecho，验证对数据库的操作无致命问题
- [ ] 部署
  - [ ] 容器化（试用https://github.com/ko-build/ko）
  - [ ] k8s

## Usage

```bash
  -h string
    	database host (default "localhost")
  -n string
    	database name (default "typecho")
  -p int
    	database port (default 3307)
  -pd string
    	database password (default "root")
  -t string
    	database type <mysql> (default "mysql")
  -u string
    	database user (default "root")
```

## 接口文档

新建文章
```bash
curl --location --request PUT 'http://localhost:8080/contents' \
--header 'Content-Type: application/json' \
--data '{
    "content": {
        "title": "title",
        "slug": "slug-1709260916102",
        "text": "content"
    },
    "meta": {
        "categories": ["计算机", "其他", "7CWZo"],
        "tags": ["a","b","7CWZo"]
    }
}'
```

更新文章
```bash
curl --location --request PUT 'http://localhost:8080/contents' \
--header 'Content-Type: application/json' \
--data '{
    "content": {
        "cid": 302,
        "title": "title",
        "slug": "slug",
        "text": "content1111111"
    }
}'
```
