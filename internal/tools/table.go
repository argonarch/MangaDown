package tools

import (
	"fmt"
	"strings"
)

func Spaces(totalLength int, uString string) (string, int, int) {
	numLength := 0
	name := ""
	if len(uString) <= totalLength {
		numLength = len(uString)
		name = uString
	} else {
		numLength = totalLength
		name = uString[:numLength]
	}

	padding := totalLength - numLength
	leftPadding := padding / 2
	rightPadding := padding - leftPadding
	return name, leftPadding, rightPadding
}

func PrintTable2Parts(
	firstString string,
	firstSpace int,
	secondString string,
	secondSpace int,
	initialSpace bool,
	finalSpace bool,
) {
	line := "+" + strings.Repeat("-", firstSpace) + "+" + strings.Repeat("-", secondSpace) + "+"
	if initialSpace {
		fmt.Println(line)
	}
	firstName, idLeft, idRight := Spaces(firstSpace, firstString)
	secondName, mangaLeft, mangaRight := Spaces(secondSpace, secondString)

	formattedString := fmt.Sprintf("|%s%s%s|%s%s%s|",
		strings.Repeat(" ", idLeft),
		firstName,
		strings.Repeat(" ", idRight),
		strings.Repeat(" ", mangaLeft),
		secondName,
		strings.Repeat(" ", mangaRight),
	)
	fmt.Println(formattedString)
	if finalSpace {
		fmt.Println(line)
	}
}

func PrintTable1Part(
	firstString string,
	firstSpace int,
	initialSpace bool,
	finalSpace bool,
) {
	line := "+" + strings.Repeat("-", firstSpace) + "+"
	if initialSpace {
		fmt.Println(line)
	}
	firstName, idLeft, idRight := Spaces(firstSpace, firstString)

	formattedString := fmt.Sprintf("|%s%s%s|",
		strings.Repeat(" ", idLeft),
		firstName,
		strings.Repeat(" ", idRight),
	)
	fmt.Println(formattedString)
	if finalSpace {
		fmt.Println(line)
	}
}

func PrintTableEqualParts(
	allStrings []string,
	space int,
	parts int,
	initialSpace bool,
	finalSpace bool,
) {
	unit := strings.Repeat("-", space) + "+"
	line := "+" + strings.Repeat(unit, parts)

	if initialSpace {
		fmt.Println(line)
	}

	text := "|"
	for _, value := range allStrings {

		name, nLeft, nRight := Spaces(space, value)

		formattedString := fmt.Sprintf("%s%s%s|",
			strings.Repeat(" ", nLeft),
			name,
			strings.Repeat(" ", nRight),
		)
		text += formattedString
	}

	if len(allStrings) != parts {
		for i := 0; i < parts-len(allStrings); i++ {
			text += strings.Repeat(" ", space) + "|"
		}
	}

	fmt.Println(text)

	if finalSpace {
		fmt.Println(line)
	}

}
