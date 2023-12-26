package main

import (
	"fmt"
	"os"
	"path"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}

func search(key string) *Entry {
	for i, v := range data {
		if v.Surname == key {
			return &data[i]
		}
	}
	return nil

}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s list|search <arguments>\n", exe)
		return

	}

	data = append(data, Entry{"Mihalis", "Steven", "2345435678"})
	data = append(data, Entry{"Cornelia", "Ekpenyong", "2345678239"})
	data = append(data, Entry{"Richard", "Edward", "2345678903"})

	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: Search Surname")
			return
		}
		result := search(arguments[2])
		if result == nil {
			fmt.Println("Entry not found:", arguments[2])
			return
		}
		fmt.Println(*result)

	case "list":
		list()

	default:
		fmt.Println("Not a valid option")

	}
}
