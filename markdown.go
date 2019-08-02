package table

import "strings"

func Markdown(matrix [][]string) []string {
	if len(matrix) < 1 {
		return []string{}
	}

	var tbl []string

	// Header
	head := matrix[0]
	tbl = append(tbl, markdownRow(head...))
	tbl = append(tbl, markdownHeaderLine(len(head)))

	// Body
	for _, row := range matrix[1:] {
		tbl = append(tbl, markdownRow(row...))
	}
	return tbl
}

func markdownRow(cols ...string) string {
	return "| " + strings.Join(cols, " | ") + " |"
}

func markdownHeaderLine(colCount int) string {
	const col = ":---:"
	var cols []string
	for i := 0; i < colCount; i++ {
		cols = append(cols, col)
	}
	return markdownRow(cols...)
}
