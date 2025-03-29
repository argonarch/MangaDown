package tools

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func RequestHtml(url string) *goquery.Document {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func Map[T any, U any](arr []T, fn func(T) U) []U {
	result := make([]U, len(arr))
	for i, v := range arr {
		result[i] = fn(v)
	}
	return result
}

func Mip[T any](arr []T, fn func(T)) {
	if len(arr) == 0 {
		return // Caso base: la lista está vacía
	}

	fn(arr[0])       // Aplicar la función al elemento actual
	Mip(arr[1:], fn) // Llamada recursiva con el resto de la lista
}

func Filter[T any](arr []T, fn func(T) bool) []T {
	result := []T{}
	for _, v := range arr {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func FilterU[T any](arr []T, fn func(T) bool) T {
	var result T
	for _, v := range arr {
		if fn(v) {
			return v
		}
	}
	return result
}

func FormatFloat(f float64) string {
	// Formatea el float con una precisión razonable (puedes ajustarla)
	s := strconv.FormatFloat(f, 'f', 10, 64)
	// Elimina los ceros finales
	s = strings.TrimRight(s, "0")
	// Elimina el punto decimal si no quedan decimales
	s = strings.TrimRight(s, ".")
	return s
}

func StringCleaner(s string) string {
	text := strings.ReplaceAll(s, "\t", "")
	text = strings.ReplaceAll(text, "\n", "")
  return text
}

func CaesarCipher(text string, shift int) string {
	result := strings.Builder{}
	for _, char := range text {
		if 'a' <= char && char <= 'z' {
			shifted := (char - 'a' + rune(shift%26) + 26) % 26 + 'a'
			result.WriteRune(shifted)
		} else if 'A' <= char && char <= 'Z' {
			shifted := (char - 'A' + rune(shift%26) + 26) % 26 + 'A'
			result.WriteRune(shifted)
		} else {
			shifted := char + rune(shift) // Aplicar desplazamiento a otros caracteres también
			result.WriteRune(shifted)
		}
	}
	return result.String()
}
