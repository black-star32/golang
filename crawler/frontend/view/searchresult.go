package view

import (
	mm "golang/crawler/frontend/model"
	"html/template"
	"io"
	"log"
)

type SearchReusltView struct {
	template *template.Template
}

func CreateSearchResultView(filename string) SearchReusltView{
	//return SearchReusltView{
	//	template: template.Must(
	//		template.ParseFiles(filename)),
	//}
	template, err := template.ParseFiles(filename)
	if err != nil{
		log.Print("no template file")
		//panic(err)
	}
	return SearchReusltView{
		template: template,
	}
}

func (s SearchReusltView) Render(w io.Writer, data mm.SearchResult) error {
	return s.template.Execute(w, data)
}