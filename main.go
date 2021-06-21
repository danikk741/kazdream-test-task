package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"

	mp "kazdream-test-task/map"
	string "kazdream-test-task/string"
	strings "kazdream-test-task/strings"
)

func main() {
	file, err := os.Open("mobydick.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	raw := string.String{
		Bytes: b,
	}
	lower := strings.ToLower(raw)
	arr := strings.SplitWords(lower)

	newMap := mp.NewMap(13)
	for _, ar := range arr {
		newMap.Add(ar)
	}

	arrFromMap := newMap.ToArray()
	sort.Slice(arrFromMap, func(i, j int) bool {
		return arrFromMap[i].V > arrFromMap[j].V
	})
	arrFromMap = arrFromMap[:20]
	for _, ar := range arrFromMap {
		fmt.Printf("%s %d\n", ar.K.Bytes, ar.V)
	}
}
