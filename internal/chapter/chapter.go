package chapter

import (
	"MangaDown/internal/tools"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Chapter struct {
	Id       float32
	IdString string
	Name     string
	Url      string
}

func OnlyUrl(chapter Chapter) string {
	return chapter.Url
}

func ScrapeChapters(urlManga string) []Chapter {
	var doc *goquery.Document = tools.RequestHtml(urlManga)
	var chapters []Chapter
	// Find url for each item
	doc.Find("div.listing-chapters_wrap ul.main li.wp-manga-chapter").
		Each(func(i int, s *goquery.Selection) {
			s.Find("a").
				First().
				Each(func(e int, a *goquery.Selection) {

					href, existsHref := a.Attr("href")
					preTitle := a.Text()
					preTitle = strings.ReplaceAll(preTitle, "Chapter ", "")
					preTitle = tools.StringCleaner(preTitle)
					floatValue, err := strconv.ParseFloat(preTitle, 32)
					title := "Chapter " + preTitle
					if err == nil {
						if existsHref {
							chapter := Chapter{
								Id:       float32(floatValue),
								IdString: preTitle,
								Name:     title,
								Url:      tools.StringCleaner(href),
							}
							chapters = append(chapters, chapter)
						}
					}
				})
		})

	return chapters
}

func SelectRangeChapters(chapters []Chapter, minChapter int, maxChapter int) []Chapter {
	return tools.Filter(
		chapters,
		func(c Chapter) bool { return c.Id >= float32(minChapter) && c.Id <= float32(maxChapter) },
	)
}

// Prints
func PrintChapters(chapters []Chapter, nameManga string) {

	minC := tools.FormatFloat(float64(chapters[len(chapters)-1].Id))
	maxC := tools.FormatFloat(float64(chapters[0].Id))
	title := "Chapters of " + nameManga + " [" + minC + ":" + maxC + "]"
	tools.PrintTable1Part(title, 76, true, false)
	fmt.Println("+----------+----------+----------+----------+----------+----------+----------+")

	// Crear una nueva lista para almacenar solo los nombres
	idsString := make([]string, len(chapters))

	// Iterar sobre la lista de objetos y extraer el atributo Nombre
	for i, objeto := range chapters {
		idsString[i] = objeto.IdString
	}

	groupSize := 7
	for i := 0; i < len(idsString); i += groupSize {
		end := i + groupSize
		end = slices.Min([]int{end, len(idsString)})
		group := idsString[i:end]
		tools.PrintTableEqualParts(group, 10, groupSize, false, true)
	}
}

func InputRangeChapter() (int, int) {
	var start int
	var final int
	fmt.Print("Min Chapter: ")
	fmt.Scanln(&start)
	fmt.Print("Max Chapter: ")
	fmt.Scanln(&final)
	return start, final
}
