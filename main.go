package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	mp "kazdream-test-task/map"
	string "kazdream-test-task/string"
	strings "kazdream-test-task/strings"
)

func mergeSort(items []mp.Elem) []mp.Elem {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]mp.Elem, middle)
		right = make([]mp.Elem, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []mp.Elem) (result []mp.Elem) {
	result = make([]mp.Elem, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0].V > right[0].V {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

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
	sorted := mergeSort(arrFromMap)
	sorted = sorted[:20]
	for _, ar := range sorted {
		fmt.Printf("%d %s\n", ar.V, ar.K.Bytes)
	}
}
