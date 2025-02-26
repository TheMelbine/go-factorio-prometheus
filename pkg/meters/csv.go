package meters

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/daanv2/go-factorio-otel/pkg/generics"
	"go.opentelemetry.io/otel/attribute"
	"golang.org/x/exp/constraints"
)

func CSVScraper[T constraints.Integer | constraints.Float](header string, cmd string) Scrape[T] {
	headers := strings.Split(header, ",")
	if len(headers) == 0 {
		panic("header must have at least one field")
	}
	if headers[0] != "amount" {
		panic("first field of header must be 'amount'")
	}

	parse := CSVToAttributesParser[T](headers)
	return func(ctx context.Context, executor Executor) ([]Point[T], error) {
		out, err := executor.Execute(cmd)
		if err != nil {
			return nil, err
		}
		return parse(out)
	}
}

// CSVToAttributesParser parses CSV data into a list of OTEL attributes.
// and returns a slice where each element is a list of OTEL attributes for one record,
// excluding the index field.
func CSVToAttributesParser[T constraints.Integer | constraints.Float](headers []string) func(csvData string) ([]Point[T], error) {
	return func(csvData string) ([]Point[T], error) {
		reader := csv.NewReader(strings.NewReader(csvData))
		var points []Point[T]
		for {
			record, err := reader.Read()
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				return nil, err
			}
			p, err := parsePoint[T](headers, record)
			if err != nil {
				log.Warnf("failed to parse record: %v \n%v\n%v", err, headers, record)
				continue
			}

			points = append(points, p)
		}
		return points, nil
	}
}

func parsePoint[T constraints.Integer | constraints.Float](header []string, record []string) (Point[T], error) {
	amount, err := generics.ParseNumber[T](record[0])
	if err != nil {
		return Point[T]{}, err
	}

	labelsValues := record[1:]
	labelsKeys := header[1:]
	if len(labelsKeys) != len(labelsValues) {
		return Point[T]{}, errors.New("header and record have different lengths")
	}

	p := Point[T]{
		Amount: amount,
		Labels: make([]attribute.KeyValue, 0, len(labelsKeys)),
	}
	for i, key := range labelsKeys {
		p.Labels = append(p.Labels, attribute.String(strings.TrimSpace(key), labelsValues[i]))
	}
	return p, nil
}