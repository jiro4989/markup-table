/*
Table is a package for generating tables of markup languages.
*/
package table

import "strings"

// Asciidoc returns a asciidoc table from matrix.
//
// Asciidoc always has a table header.
func Asciidoc(matrix [][]string) []string {
	if len(matrix) < 1 {
		return []string{}
	}

	var tbl []string

	tbl = append(tbl, `[options="header"]`)
	tbl = append(tbl, "|=================")
	for _, row := range matrix {
		tbl = append(tbl, asciidocRow(row...))
	}

	tbl = append(tbl, "|=================")
	return tbl
}

func asciidocRow(cols ...string) string {
	var cols2 []string
	for _, col := range cols {
		col = strings.Replace(col, "\n", "+\n", -1)
		cols2 = append(cols2, col)
	}
	return "| " + strings.Join(cols2, " | ")
}
