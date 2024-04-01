# Typecho-api

[![Go Report Card](https://goreportcard.com/badge/github.com/lwabish/typecho-api)](https://goreportcard.com/report/github.com/lwabish/typecho-api)

- 提供http api，直接对接操作typecho数据库
- 基于typora的自定义导出功能，实现从typora一键发布到typecho

## Roadmap

### basic

- [x] 创建文章
- [x] 处理分类和标签
- [x] typora对接api方案和实现
- [x] 分类标签计数更新
- [ ] 更新文章
- [ ] 单元测试
- [ ] e2e测试
- [ ] 部署
  - [ ] 容器化（试用https://github.com/ko-build/ko）
  - [ ] k8s

### further

- [ ] 插件化【高优先级】，为其他各类修改文章的功能提供框架和接口设计
- [ ] 文章字数更新
- [ ] 其他高级功能
  - [ ] 自动添加TOC
  - [ ] 自动添加文章摘要（调用AI服务生成）
  - [ ] 自动添加文章封面（调用AI服务生成）
  - [ ] 自动拼写检查（调用AI服务实现）
  - [ ] 自动添加文末原文链接，防采集
  - [ ] 统一文章中所有英文的风格，比如自动将首字母大写
  - [ ] 基于LLM，实现文字语气的统一
  - [ ] 统一句号
  - [ ] 鉴权

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

## Use-with-typora

[文档](./client/Readme.md)

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
