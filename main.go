package main

import "fmt"

func main() {
	var list1 []string = []string{"ice", "apple", "book", "test"}
	var list2 []string = []string{"ice", "apple", "bird", "test"}
	fmt.Println(levenshtein(list1, list2))
}

func levenshtein(taglist1, taglist2 []string) float32 {
	maxlen := maxnum(len(taglist1), len(taglist2))
	var tags1 = make([]string, maxlen)
	var tags2 = make([]string, maxlen)

	tags1 = taglist1
	tags2 = taglist2

	num := 0

	for _, j := range tags1 {
		for _, jj := range tags2 {
			if jj == j {
				num++
				break
			}
		}
	}

	return float32(num) / float32(maxlen)
}

func maxnum(vals ...int) int {
	var max int
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}
