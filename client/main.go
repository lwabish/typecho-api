package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/lwabish/typecho-api/handlers/content"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type frontMatter struct {
	Cid int `yaml:"cid"`
}

var (
	server string
	path   string
)

func init() {
	flag.StringVar(&server, "s", "http://127.0.0.1:8080", "typecho-api server address")
	flag.StringVar(&path, "p", "", "post path")
}

func main() {
	flag.Parse()
	log.Println("typecho-api server address:", server)
	log.Println("post path:", path)

	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	fm := newFrontMatter(string(bytes))
	log.Println("fm:", fm)

	payload := buildPayload(fm)
	log.Println("payload:", payload)

	res, err := putContent(payload)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("res:", res)

}

func buildPayload(matter *frontMatter) *content.VO {
	//`{
	//    "content": {
	//        "title": "title",
	//        "slug": "slug-1709260916102",
	//        "text": "content"
	//    },
	//    "meta": {
	//        "categories": ["计算机", "其他", "7CWZo"],
	//        "tags": ["a","b","7CWZo"]
	//    }
	//}`
	vo := content.NewVo()

	if matter.Cid != 0 {
		vo.SetCid(matter.Cid)
	}
	return vo
}

func newFrontMatter(markdown string) *frontMatter {
	f := &frontMatter{}
	re := regexp.MustCompile(`(?s)^---\n(.+?)\n---`)
	matches := re.FindStringSubmatch(markdown)
	if len(matches) >= 2 {
		yamlContent := matches[1]
		if err := yaml.Unmarshal([]byte(yamlContent), f); err != nil {
			log.Println(err)
		}
	}
	return f
}

func putContent(vo *content.VO) (string, error) {
	payload, err := json.Marshal(vo)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/contents", server), strings.NewReader(string(payload)))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), err
}
