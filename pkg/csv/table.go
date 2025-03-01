package csv

import (
	"encoding/csv"
	"errors"
	"io"
	"slices"
	"strings"
)

type Table struct {
	Headers  []string
	Records []Record
}

type Record struct {
	Values []string
}

func (r Record) GetValue(index int) (string, bool) {
	if index >= len(r.Values) {
		return "", false
	}

	return r.Values[index], true
}

func Parser(headers []string) func(values string) (Table, error) {
	return func(values string) (Table, error) {
		return Parse(headers, values)
	}
}

func Parse(headers []string, data string) (Table, error) {
	if len(headers) == 0 {
		return Table{}, errors.New("no headers provided")
	}

	table := Table{
		Headers: headers,
	}

	reader := csv.NewReader(strings.NewReader(data))
	for {
		record, err := reader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return table, err
		}
		if len(record) != len(table.Headers) {
			if len(record) > len(table.Headers) {
				return table, errors.New("record is longer then headers: " + strings.Join(record, ","))
			}
			for len(record) < len(table.Headers) {
				record = append(record, "")
			}
		}

		table.Records = append(table.Records, Record{
			Values: record,
		})
	}

	return table, nil
}

func (t Table) GetHeader(index int) (string, bool) {
	if index >= len(t.Headers) {
		return "", false
	}

	return t.Headers[index], true
}

func (t Table) HeaderIndex(header string) int {
	return slices.Index(t.Headers, header)
}

func (t Table) GetRecordWithHeader(column, row int) (header string, record string, ok bool) {
	header, ok = t.GetHeader(column)
	if !ok {
		return "", "", false
	}

	r, ok := t.GetRecord(column, row)
	if !ok {
		return header, "", false
	}

	v, ok := r.GetValue(column)
	
	return header, v, ok
}

func (t Table) GetRecord(column, row int) (Record, bool) {
	if row >= len(t.Records) {
		return Record{}, false
	}

	return t.Records[row], true
}

func (t Table) FilterColumns(headers []string) (Table, error) {
	result := Table{
		Headers: headers,
	}

	// Map of old indexes to new ones, maps old header index to new index
	translate := make(map[int]int)
	for j, h := range t.Headers {
		translate[j] = result.HeaderIndex(h)
	}
	

	for _, record := range t.Records {
		nrecord := Record{Values:make([]string, len(record.Values))}
		result.Records = append(result.Records, nrecord)

		// Move values
		for i, v := range record.Values {
			nindex, ok := translate[i]
			if !ok || nindex < 0 { // Skip because new table has no header
				continue
			}
			nrecord.Values[nindex] = v
		}
		
	}

	return result, nil
}