package api

import (
	"context"
	"fmt"
	"strings"

	"github.com/alexeyco/simpletable"
)

/****************************************
*  Terminal
*****************************************/

type Terminal struct {
	Output
}

func NewTerminal() *Terminal {
	return &Terminal{Output: NewTerminalTable()}
}

// PrintModelInfo print model info into terminal.
func (terminal *Terminal) PrintModelInfo(ctx context.Context, m *ModelInfo, _ *FormatModelInfoOption) error {
	table := NewConsoleTable()

	cells := make([]*ConsoleCell, 0)
	cells = append(cells, &ConsoleCell{
		Key:  "ID",
		Text: fmt.Sprintf("%d", m.ID),
	})
	cells = append(cells, &ConsoleCell{
		Key:  "Name",
		Text: fmt.Sprintf("%s", m.Name),
	})
	cells = append(cells, &ConsoleCell{
		Key:  "Type",
		Text: fmt.Sprintf("%s", m.Type),
	})
	cells = append(cells, &ConsoleCell{
		Key:  "Creator",
		Text: fmt.Sprintf("%s", m.Creator.Username),
	})
	cells = append(cells, &ConsoleCell{
		Key:  "Stat",
		Text: fmt.Sprintf("download:%d\nthumbsUp:%d", m.Stats.DownloadCount, m.Stats.ThumbsUpCount),
	})
	// Versions
	versionTexts := make([]string, 0)
	for _, ver := range m.ModelVersions {
		// Too many ...
		if len(versionTexts) >= 3 {
			versionTexts = append(versionTexts, "...")
			break
		} else {
			versionTexts = append(versionTexts, fmt.Sprintf("[%d][%s] %s", ver.ID, ver.Name, ver.DownloadURL))
		}
	}

	cells = append(cells, &ConsoleCell{
		Key:  "Versions",
		Text: strings.Join(versionTexts, "\n"),
	})

	table.Cells = append(table.Cells, cells)

	// print into terminal
	output, err := terminal.Output.Format(ctx, table)
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}

type FormatModelInfoOption struct {
	Full bool
}

/****************************************
*  TerminalTable
*****************************************/

type TerminalTable struct {
	Style string
}

func (out *TerminalTable) Format(_ context.Context, input *ConsoleTable) ([]byte, error) {
	if len(input.Cells) == 0 {
		return []byte{}, nil
	}

	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: make([]*simpletable.Cell, 0),
	}
	for _, col := range input.Cells[0] {
		table.Header.Cells = append(table.Header.Cells, &simpletable.Cell{
			Align: simpletable.AlignLeft,
			Text:  col.Key,
		})
	}

	tableCells := make([][]*simpletable.Cell, 0)
	for _, cells := range input.Cells {
		curr := make([]*simpletable.Cell, 0)
		for _, cell := range cells {
			curr = append(curr, &simpletable.Cell{
				Align: simpletable.AlignLeft,
				Span:  0,
				Text:  cell.Text,
			})
		}
		tableCells = append(tableCells, curr)
	}
	table.Body.Cells = tableCells

	switch out.Style {
	default:
		table.SetStyle(simpletable.StyleMarkdown)
	case "markdown":
		table.SetStyle(simpletable.StyleMarkdown)
	}

	output := table.String()
	return []byte(output), nil
}

func NewTerminalTable() *TerminalTable {
	return &TerminalTable{Style: "markdown"}
}

/****************************************
*  ConsoleOutput
*****************************************/

type Output interface {
	Format(ctx context.Context, table *ConsoleTable) ([]byte, error)
}

type ConsoleTable struct {
	Cells [][]*ConsoleCell `json:"cells,omitempty"`
}

func NewConsoleTable() *ConsoleTable {
	return &ConsoleTable{Cells: make([][]*ConsoleCell, 0)}
}

type ConsoleCell struct {
	Key  string `json:"key,omitempty"`
	Text string `json:"text,omitempty"`
}
