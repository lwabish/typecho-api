# typecho-api

- 提供http api，直接对接操作typecho数据库
- 目前仅实现了创建/更新文章接口，对接typora的导出功能即可实现typora写完文章后直接发布到typecho

## usage

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

## interface

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