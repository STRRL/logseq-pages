package command

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/strrl/logseq-pages/pkg/model"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func NewListCommand(rootOption *RootOption) *cobra.Command {
	listOption := ListOption{
		FilterType: AllFilter,
		Output:     TableOutput,
	}

	listCommand := &cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listLogseqPages(rootOption, &listOption)
		},
	}

	listCommand.Flags().StringVarP((*string)(&listOption.FilterType), "filter", "f", string(AllFilter), "filter type, available values: all, public, private")
	listCommand.Flags().StringVarP((*string)(&listOption.Output), "output", "o", string(TableOutput), "output type, available values: table, json")

	return listCommand
}

type ListOption struct {
	FilterType FilterType
	Output     OutputType
}

type OutputType string

const TableOutput OutputType = "table"
const JsonOutput OutputType = "json"

type FilterType string

const AllFilter FilterType = "all"
const PublicFilter FilterType = "public"
const PrivateFilter FilterType = "private"

func listLogseqPages(rootOption *RootOption, listOption *ListOption) error {
	var markdownFiles []string
	appendToMarkdownFiles := func(path string, d fs.DirEntry, err error) error {
		if strings.Contains(path, "md") {
			markdownFiles = append(markdownFiles, path)
		}
		return nil
	}

	err := filepath.WalkDir(rootOption.WorkDirectory, appendToMarkdownFiles)
	if err != nil {
		return err
	}
	var pages model.LogseqPageList
	for _, file := range markdownFiles {
		logseqPage, err := loadLogseqPageFromFile(file)
		if err != nil {
			continue
		}
		pages = append(pages, *logseqPage)
	}

	if listOption.FilterType == PublicFilter {
		pages = pages.Filter(func(page model.LogseqPage) bool {
			return page.Public
		})
	}
	if listOption.FilterType == PrivateFilter {
		pages = pages.Filter(func(page model.LogseqPage) bool {
			return !page.Public
		})
	}

	if listOption.Output == TableOutput {
		out, err := pages.RenderAsTable()
		if err != nil {
			return err
		}
		fmt.Println(out)
		return nil
	} else if listOption.Output == JsonOutput {
		out, err := json.Marshal(pages)
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", out)
		return nil
	} else {
		fmt.Println("unknown output type " + string(listOption.Output))
		return nil
	}
}

func loadLogseqPageFromFile(file string) (*model.LogseqPage, error) {
	result := model.LogseqPage{
		Name:   filepath.Base(file),
		Alias:  nil,
		Path:   file,
		Public: false,
	}

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	explicitlyPublic := false

	reader := bufio.NewReader(f)
	var line string
	for {
		bytes, prefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		line += string(bytes)
		if prefix {
			continue
		}

		if strings.HasPrefix(line, "public:: ") {
			explicitlyPublic = true
			isPublic := strings.TrimPrefix(line, "public:: ")
			result.Public, err = strconv.ParseBool(isPublic)
			if err != nil {
				return nil, err
			}
		}
		if strings.HasPrefix(line, "alias:: ") {
			result.Alias = append(result.Alias, strings.TrimPrefix(line, "alias:: "))
		}

		line = ""
	}
	if !explicitlyPublic {
		result.Public = false
	}

	return &result, nil
}
