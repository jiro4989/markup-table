package table

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTML(t *testing.T) {
	type TestData struct {
		desc   string
		matrix [][]string
		expect []string
	}
	tds := []TestData{
		{
			desc: "3 x 3",
			matrix: [][]string{
				{"Head 1", "Head 2", "Head 3"},
				{"Body 1", "Body 2", "Body 3"},
				{"Body 4", "Body 5", "Body 6"},
			},
			expect: []string{
				"<table>",
				"<thead>",
				"<tr><td>Head 1</td><td>Head 2</td><td>Head 3</td></tr>",
				"</thead>",
				"<tbody>",
				"<tr><td>Body 1</td><td>Body 2</td><td>Body 3</td></tr>",
				"<tr><td>Body 4</td><td>Body 5</td><td>Body 6</td></tr>",
				"</tbody>",
				"</table>",
			},
		},
		{
			desc: "1 row",
			matrix: [][]string{
				{"Head 1", "Head 2", "Head 3"},
			},
			expect: []string{
				"<table>",
				"<thead>",
				"<tr><td>Head 1</td><td>Head 2</td><td>Head 3</td></tr>",
				"</thead>",
				"<tbody>",
				"</tbody>",
				"</table>",
			},
		},
		{
			desc:   "0 row",
			matrix: [][]string{},
			expect: []string{},
		},
	}
	for _, v := range tds {
		got := HTML(v.matrix)
		assert.Equal(t, v.expect, got, v.desc)
	}
}
