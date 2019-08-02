package table

import (
	"html"
	"strings"
)

// HTML returns a HTML table from matrix.
//
// HTML wraps rows with table, thead and tbody tags.
func HTML(matrix [][]string) []string {
	if len(matrix) < 1 {
		return []string{}
	}

	var tbl []string

	tbl = append(tbl, "<table>")

	// Header
	head := matrix[0]
	tbl = append(tbl, "<thead>")
	tbl = append(tbl, htmlRow(head...))
	tbl = append(tbl, "</thead>")

	// Body
	tbl = append(tbl, "<tbody>")
	for _, row := range matrix[1:] {
		tbl = append(tbl, htmlRow(row...))
	}
	tbl = append(tbl, "</tbody>")

	tbl = append(tbl, "</table>")
	return tbl
}

func htmlRow(cols ...string) string {
	var escCols []string
	for _, col := range cols {
		c := html.EscapeString(col)
		escCols = append(escCols, c)
	}
	return "<tr><td>" + strings.Join(escCols, "</td><td>") + "</td></tr>"
}
