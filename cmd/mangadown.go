package main

import (
	"MangaDown/internal/chapter"
	"MangaDown/internal/image"
	"MangaDown/internal/manga"

	"context"
	"log"
	"os"

	cli "github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:    "manga-down",
		Version: "v1.5.0",
		Action: func(context.Context, *cli.Command) error {
			setup(manga.InputNameManga())
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "manga-search",
				Aliases: []string{"ms"},
				Usage:   "Search manga",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					setup(cmd.Args().First())
					return nil
				},
			},
		}}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func setup(name string) {
	var mangas []manga.Manga = manga.ScrapeMangas(name)
	manga.PrintMangas(mangas)
	var mangaSelected manga.Manga = manga.SelectManga(mangas, manga.InputNumberManga())
	var chapters []chapter.Chapter = chapter.ScrapeChapters(mangaSelected.Url)
	chapter.PrintChapters(chapters, mangaSelected.Name)
	minRangeChapters, maxRangeChapters := chapter.InputRangeChapter()
	var rangeChapterSelected []chapter.Chapter = chapter.SelectRangeChapters(chapters, minRangeChapters, maxRangeChapters)
	for i := len(rangeChapterSelected) - 1; i >= 0; i-- {
		image.DownloadImages(
			image.ScrapeImages(rangeChapterSelected[i].Url),
			mangaSelected.Name,
			rangeChapterSelected[i].Name,
		)
	}
}
