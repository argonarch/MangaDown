package manga

import (
	"MangaDown/internal/tools"
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Manga struct {
	Id   int
	Name string
	Url  string
}

func ScrapeMangas(nameManga string) []Manga {
	value := "lxxtw>33qerkerexs2ms3"
	urlMangaPoint := tools.CaesarCipher(value, -4)

	var doc *goquery.Document = tools.RequestHtml(urlMangaPoint + "?s=" + nameManga + "&post_type=wp-manga")
	var mangas []Manga
	// Find url for each item
	doc.Find("div.c-tabs-item div.tab-summary div.post-title a").
		Each(func(i int, a *goquery.Selection) {
			href, existsHref := a.Attr("href")
			preTitle := a.Text()
			title := strings.ReplaceAll(preTitle, "’", "'")
			title = tools.StringCleaner(title)
			if existsHref {
				manga := Manga{Id: i, Name: title, Url: tools.StringCleaner(href)}
				mangas = append(mangas, manga)
			}
		})

	return mangas
}

func SelectManga(mangas []Manga, numManga int) Manga {
	return tools.FilterU(mangas, func(m Manga) bool { return m.Id == numManga })
}

// Prints

func PrintManga(manga Manga) {
	idString := strconv.Itoa(manga.Id)
	tools.PrintTable2Parts(idString, 10, manga.Name, 65, false, true)
}

func PrintMangas(mangas []Manga) {
	tools.PrintTable2Parts("index", 10, "Mangas", 65, true, true)
	tools.Mip(mangas, PrintManga)
}

func InputNumberManga() int {
	var numManga int
	fmt.Print("Numero de manga: ")
	fmt.Scanln(&numManga)
	return numManga
}

func InputNameManga() string {
	var nameManga string
	fmt.Print("Nombre del manga: ")
	fmt.Scanln(&nameManga)
	return nameManga
}
