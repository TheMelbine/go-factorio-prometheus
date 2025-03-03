package meters

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/charmbracelet/log"
	csvx "github.com/daanv2/go-factorio-otel/pkg/csv"
	"github.com/daanv2/go-factorio-otel/pkg/generics"
	"github.com/daanv2/go-factorio-otel/pkg/lua"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/exp/constraints"
)

func GrabCSVTable(headers []string, cmd string, executor Executor) (csvx.Table, error) {
	out, err := executor.Execute(cmd)
	if err != nil {
		return csvx.Table{}, fmt.Errorf("error executing: %v => %w", cmd, err)
	}
	p, err := csvx.Parse(headers, out)
	if err != nil {
		return csvx.Table{}, fmt.Errorf("error parsing: %v => %w", out, err)
	}
	return p, nil
}

func CSVScraper[T constraints.Integer | constraints.Float](header string, cmd string) Scrape[T] {
	headers := strings.Split(header, ",")
	cmd = lua.SingleLine(cmd)
	return func(ctx context.Context, executor Executor) ([]Point[T], error) {
		table, err := GrabCSVTable(headers, cmd, executor)
		if err != nil {
			return nil, fmt.Errorf("couldn't grab table: %w", err)
		}

		return CSVTableToPoints[T](table), nil
	}
}

func CSVTableToPoints[T constraints.Integer | constraints.Float](table csvx.Table) []Point[T] {
	var points []Point[T]
	for _, record := range table.Records {
		p, err := parsePoint[T](table.Headers, record.Values)
		if err != nil {
			log.Warnf("failed to parse record: %v \n%v\n%v", err, table.Headers, record)
			continue
		}

		points = append(points, p)
	}
	return points
}

func parsePoint[T constraints.Integer | constraints.Float](header []string, record []string) (Point[T], error) {
	v := record[0]
	var (
		amount T
		err    error
	)

	if v != "" {
		amount, err = generics.ParseNumber[T](v)
		if err != nil {
			return Point[T]{}, err
		}
	}

	labelsValues := record[1:]
	labelsKeys := header[1:]
	if len(labelsKeys) != len(labelsValues) {
		log.Debug("header and record have different lengths", "values", labelsValues, "keys", labelsKeys)
		return Point[T]{}, errors.New("header and record have different lengths")
	}

	p := Point[T]{
		Amount: amount,
		Labels: prometheus.Labels{},
	}
	for i, key := range labelsKeys {
		p.Labels[key] = labelsValues[i]
	}
	return p, nil
}
