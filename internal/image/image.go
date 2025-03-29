package image

import (
	"MangaDown/internal/tools"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Image struct {
	Id  int
	Url string
}

func ScrapeImages(urlChapter string) []Image {
	var doc *goquery.Document = tools.RequestHtml(urlChapter)
	var images []Image
	// Find url for each item
	doc.Find("div.read-container div.reading-content div.page-break").
		Each(func(i int, s *goquery.Selection) {
			s.Find("img").
				First().
				Each(func(e int, a *goquery.Selection) {

					href, existsHref := a.Attr("src")
					if existsHref {
						image := Image{
							Id:  i,
							Url: tools.StringCleaner(href),
						}
						images = append(images, image)
					}
				})
		})

	return images
}

func DownloadImages(images []Image, nameManga string, chapterName string) {
	tools.Mip(images, func(u Image) { DownloadImage(u, nameManga, chapterName) })
}

func DownloadImage(image Image, nameManga string, chapterName string) {
	directory := "/hdd/kael/Documentos/Mangas/" + nameManga + "/" + chapterName + "/"
	fileName := strconv.Itoa(image.Id) + ".jpg"

	// Crear la carpeta (solo si no existe)
	os.MkdirAll(directory, 0755)

	// Hacer la solicitud GET
	resp, err := http.Get(image.Url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Crear el archivo local para guardar la imagen
	out, err := os.Create(directory + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// Copiar el contenido de la respuesta al archivo local
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

}
