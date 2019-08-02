package table_test

import (
	"fmt"

	table "github.com/jiro4989/markup-table"
)

func ExampleMarkdown() {
	matrix := [][]string{

		{"Head 1", "Head 2", "Head 3"},
		{"Body 1", "Body 2", "Body 3"},
		{"Body 4", "Body 5", "Body 6"},
	}
	for _, row := range table.Markdown(matrix) {
		fmt.Println(row)
	}
	// Output:
	// | Head 1 | Head 2 | Head 3 |
	// | :---: | :---: | :---: |
	// | Body 1 | Body 2 | Body 3 |
	// | Body 4 | Body 5 | Body 6 |
}
