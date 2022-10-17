package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	testString := "яЯАAAxctcbaaaaaaaaacbbbccccaa"
	collapseString(testString)
}

func collapseString(testString string) {

	fmt.Println(`Я бы воспользовался regexp backtracking("([a-zA-Zа-яА-Я])\1+"), но стандартная библиотека regexp такой функционал не поддерживает`)

	asSymbols := strings.Split(testString, "")
	counter := make(map[string]int)
	sorted := make([]string, 0)
	for _, v := range asSymbols {
		if _, ok := counter[v]; !ok {
			counter[v] = strings.Count(testString, v)
			sorted = append(sorted, v)
		}
	}

	sort.Strings(sorted)
	var b strings.Builder
	for _, v := range sorted {
		b.WriteString(fmt.Sprintf("%s%d", v, counter[v]))
	}
	result := b.String()

	fmt.Println(result)
}
