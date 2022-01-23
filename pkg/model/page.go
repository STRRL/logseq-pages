package model

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"strings"
)

type LogseqPage struct {
	Name   string
	Alias  []string
	Path   string
	Public bool
}

type LogseqPageList []LogseqPage

func (it *LogseqPageList) RenderAsTable() (string, error) {
	tableWriter := table.NewWriter()
	tableWriter.AppendHeader(table.Row{
		"#",
		"Name",
		"Public",
		"Alias",
		"Path",
	})

	for index, item := range *it {
		var columnPublic string
		if item.Public {
			columnPublic = "*"
		} else {
			columnPublic = ""
		}
		tableWriter.AppendRow(table.Row{
			index,
			item.Name,
			columnPublic,
			strings.Join(item.Alias, ", "),
			item.Path,
		})
	}
	tableWriter.AppendFooter(table.Row{"Total", len(*it)})
	return tableWriter.Render(), nil
}

func (it *LogseqPageList) Filter(filter func(item LogseqPage) bool) LogseqPageList {
	var result LogseqPageList
	for _, item := range *it {
		if filter(item) {
			result = append(result, item)
		}
	}
	return result
}
