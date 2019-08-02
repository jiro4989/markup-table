package table

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsciidoc(t *testing.T) {
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
				`[options="header"]`,
				`|=================`,
				"| Head 1 | Head 2 | Head 3",
				"| Body 1 | Body 2 | Body 3",
				"| Body 4 | Body 5 | Body 6",
				`|=================`,
			},
		},
		{
			desc: "1 row",
			matrix: [][]string{
				{"Head 1", "Head 2", "Head 3"},
			},
			expect: []string{
				`[options="header"]`,
				`|=================`,
				"| Head 1 | Head 2 | Head 3",
				`|=================`,
			},
		},
		{
			desc:   "0 row",
			matrix: [][]string{},
			expect: []string{},
		},
	}
	for _, v := range tds {
		got := Asciidoc(v.matrix)
		assert.Equal(t, v.expect, got, v.desc)
	}
}
