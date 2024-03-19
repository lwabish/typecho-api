package main

import (
	"fmt"
	"github.com/lwabish/typecho-api/models"
	"gopkg.in/yaml.v3"
	"log"
	"regexp"
)

type FrontMatter struct {
	Cid        int               `yaml:"cid"`
	Categories []models.Category `yaml:"categories"`
	Tags       []models.Tag      `yaml:"tags"`
}

func (m *FrontMatter) String() string {
	out, err := yaml.Marshal(m)
	if err != nil {
		log.Printf("error: %v", err)
		return ""
	}
	return fmt.Sprintf("---\n%s\n---\n", string(out))
}

func ParseFrontMatter(markdown string) *FrontMatter {
	f := &FrontMatter{}
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
