package scrapper

import (
	"fmt"
	"regexp"
)

type Tag struct {
	tags  []Tag
	typ   string
	class string
	text  string
}

func ParseHtml(html string) []Tag {
	re := regexp.MustCompile(`/<([A-Za-z][A-Za-z0-9]*)\b[^>]*\bclass="([^"]*)"[^>]*>(.*?)/`)

	regix_parsed := re.FindAllStringSubmatch(html, -1)
	fmt.Println("Regix parsed: ", regix_parsed)

	result_html := []Tag{}

	for i := range regix_parsed {
		inner_tags := []Tag{}
		if regix_parsed[i][3] != "" {
			inner_tags = ParseHtml(regix_parsed[i][3])
		}

		tag := Tag{
			typ:   regix_parsed[i][1],
			class: regix_parsed[i][2],
			text:  regix_parsed[i][3],
			tags:  inner_tags,
		}
		result_html = append(result_html, tag)
	}
	return result_html
}

func (t *Tag) PrintTags() {
	fmt.Println(t.typ, t.class, t.text)
	for i := range t.tags {
		t.tags[i].PrintTags()
	}
}
