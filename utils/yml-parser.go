package utils

import (
	"fmt"

	"github.com/goccy/go-yaml"
)

// return's a yaml string
// func ParseYML(path string) string {

// }

type Book struct {
	Author string `yaml:"author"`
	Price  int    `yaml:"price"`
}

type Store struct {
	BookStore []Book `yaml:"book"`
	Bicycle   any    `yaml:"bicycle"`
}

type Ymlmodel struct {
	Store Store `yaml:"store"`
}

func ParserTest() {
	yml := `
store:
  book:
    - author: john
      price: 10
    - author: ken
      price: 12
  bicycle:
    color: red
    price: 19.95
`
	var ym Ymlmodel
	err := yaml.Unmarshal([]byte(yml), &ym)
	if err != nil {
		panic(err)
	}

	for _, book := range ym.Store.BookStore {
		fmt.Println(book.Author, "quotes the price to be : ", book.Price)
	}

}
