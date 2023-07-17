package main

import (
	"flag"
	"fmt"
	"strings"

	"Course_golang/pkg/crawler"
	"Course_golang/pkg/crawler/spider"
)

func main() {
	searchWord := flag.String("s", "", "Слово для поиска в ссылках")
	flag.Parse()

	s := spider.New()

	arr := make([]crawler.Document, 0)

	goDevData, err := s.Scan("https://go.dev", 2)
	arr = append(arr, goDevData...)
	if err != nil {
		fmt.Printf("Ошибка при сканировании сайта go.dev: %v\n", err)
		return
	}

	golangOrgData, err := s.Scan("https://golang.org", 2)
	arr = append(arr, golangOrgData...)
	if err != nil {
		fmt.Printf("Ошибка при сканировании сайта golang.org: %v\n", err)
		return
	}

	if *searchWord != "" {
		fmt.Printf("Результаты поиска для слова \"%s\":\n", *searchWord)
		for _, document := range arr {
			if strings.Contains(document.URL, *searchWord) || strings.Contains(document.Title, *searchWord) {
				fmt.Println("URL:", document.URL)
				fmt.Println("Заголовок:", document.Title)
				fmt.Println()
			}
		}
	} else {
		fmt.Println("Флаг не указан")
	}
}
