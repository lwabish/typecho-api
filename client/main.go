package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/lwabish/typecho-api/handlers/content"
	"github.com/lwabish/typecho-api/utils"
	"io"
	"log"
	"net/http"
	"os"
	p "path"
	"strconv"
	"strings"
)

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

	_, filename := p.Split(path)
	title := strings.SplitN(filename, "-", 2)[1]

	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	c := string(bytes)
	pureContent := utils.RemoveFrontMatter(c)

	fm := ParseFrontMatter(c)
	log.Println("fm:", fm)

	payload := buildPayload(fm, title, pureContent)
	log.Println("payload:", payload)

	res, err := putContent(payload)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("res:", res)
	cid, err := strconv.Atoi(res)
	if err != nil {
		log.Fatalf("%s is not a number: %v", res, err)
	}
	fm.Cid = cid

	finalContent := strings.Join(append([]string{
		fm.String(),
	}, pureContent), "")

	err = os.WriteFile(path, []byte(finalContent), 0644)
	if err != nil {
		log.Printf("write file failed: %v", err)
		return
	}
}

func buildPayload(matter *FrontMatter, title string, text string) *content.VO {
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
	vo := content.NewVo().
		SetText(text).
		SetTitle(title).
		SetSlug(title).
		SetCategories(matter.Categories).
		SetTags(matter.Tags)

	if matter.Cid != 0 {
		vo.SetCid(matter.Cid)
	}
	return vo
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
