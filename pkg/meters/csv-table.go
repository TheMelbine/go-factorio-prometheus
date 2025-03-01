package meters

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/daanv2/go-factorio-otel/pkg/csv"
	"github.com/daanv2/go-factorio-otel/pkg/meters/cost"
)

type CSVTable struct {
	name    string
	cmd     string
	headers []string
	table   csv.Table
}

func (t *CSVTable) Name() string {
	return t.name
}

func (t *CSVTable) Table() csv.Table {
	return t.table
}

func (t *CSVTable) Cost() cost.Cost {
	return cost.MINIMAL
}

func NewCSVTable(name, cmd string, headers []string) *CSVTable {
	return &CSVTable{
		name:    name,
		cmd:     cmd,
		headers: headers,
		table: csv.Table{
			Headers: headers,
		},
	}
}

func (t *CSVTable) Scrape(ctx context.Context, executor Executor) error {
	logger := log.WithPrefix("datasource-" + t.name)
	logger.Debug("Scraping table")
	table, err := GrabCSVTable(t.headers, t.cmd, executor)
	if err != nil {
		return err
	}
	t.table = table
	return nil
}

func (t *CSVTable) SubTableInt64(headers ...string) Scrape[int64] {
	return func(ctx context.Context, executor Executor) ([]Point[int64], error) {
		data, err := t.table.FilterColumns(headers)
		if err != nil {
			return nil, err
		}

		return CSVTableToPoints[int64](data), nil
	}
}

func (t *CSVTable) SubTableFloat64(headers ...string) Scrape[int64] {
	return func(ctx context.Context, executor Executor) ([]Point[int64], error) {
		data, err := t.table.FilterColumns(headers)
		if err != nil {
			return nil, err
		}

		return CSVTableToPoints[int64](data), nil
	}
}