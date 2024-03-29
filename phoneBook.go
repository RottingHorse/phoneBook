package main

import (
	"fmt"
	"os"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data []Entry

func search(key string) *Entry {
	for i, v := range data {
		if key == v.Surname {
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
		exe := arguments[0]
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	data = append(data, Entry{
		Name:    "FirstName",
		Surname: "FirstSurname",
		Tel:     "12345",
	})
	data = append(data, Entry{
		Name:    "SecondName",
		Surname: "SecondSurname",
		Tel:     "34567",
	})
	data = append(data, Entry{
		Name:    "ThirdName",
		Surname: "ThirdSurname",
		Tel:     "56789",
	})

	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		result := search(arguments[2])
		if result == nil {
			fmt.Println("Entry not found", arguments[2])
			return
		}
		fmt.Println(*result)
	case "list":
		list()
	default:
		fmt.Println("Not a valid option")
	}

}
