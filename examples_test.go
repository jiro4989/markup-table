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

func ExampleHTML() {
	matrix := [][]string{

		{"Head 1", "Head 2", "Head 3"},
		{"Body 1", "Body 2", "Body 3"},
		{"Body 4", "Body 5", "Body 6"},
	}
	for _, row := range table.HTML(matrix) {
		fmt.Println(row)
	}
	// Output:
	// <table>
	// <thead>
	// <tr><td>Head 1</td><td>Head 2</td><td>Head 3</td></tr>
	// </thead>
	// <tbody>
	// <tr><td>Body 1</td><td>Body 2</td><td>Body 3</td></tr>
	// <tr><td>Body 4</td><td>Body 5</td><td>Body 6</td></tr>
	// </tbody>
	// </table>
}

func ExampleAsciidoc() {
	matrix := [][]string{

		{"Head 1", "Head 2", "Head 3"},
		{"Body 1", "Body 2", "Body 3"},
		{"Body 4", "Body 5", "Body 6"},
	}
	for _, row := range table.Asciidoc(matrix) {
		fmt.Println(row)
	}
	// Output:
	// [options="header"]
	// |=================
	// | Head 1 | Head 2 | Head 3
	// | Body 1 | Body 2 | Body 3
	// | Body 4 | Body 5 | Body 6
	// |=================
}
