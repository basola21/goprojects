package scrapper

type Tag struct {
	tags  []Tag
	typ   string
	class string
}

type Html struct{ tags []Tag }

func ParseHtml(html string) Html {
	return Html{}
}
