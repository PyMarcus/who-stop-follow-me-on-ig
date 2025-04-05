package internal

import (
	"fmt"
	"sort"
	"strconv"
)

func checkFollowing(name string, followers []string) bool {
	return binarySearch(name, followers)
}

func binarySearch(name string, followers []string) bool {
	if len(followers) <= 0 {
		fmt.Println("Não encontrou " + name)
		return false
	}
	size := len(followers) / 2
	middle := followers[size]
	if name == middle {
		return true
	} else if name < middle {
		return binarySearch(name, followers[:size])
	}
	return binarySearch(name, followers[size+1:])
}

func Compare(followers, following []string) []string {
	sort.Strings(followers)
	names := []string{}
	i := 1
	for _, name := range following {
		if !checkFollowing(name, followers) {
			fmt.Println("Não te segue: " + strconv.Itoa(i) + ") " + name)
			names = append(names, strconv.Itoa(i)+") "+name)
			i++
		}
	}
	return names
}
