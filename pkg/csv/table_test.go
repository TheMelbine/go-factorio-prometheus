package csv_test

import (
	"strings"
	"testing"

	"github.com/daanv2/go-factorio-prometheus/pkg/csv"
	"github.com/stretchr/testify/require"
)

const (
	header    = "amount,planet,player,health,speed"
	raw_table = `1,nauvis,daan,100,1
2,vulcanus,frodo,50,3
3,fulgora,gandalf,77,4.5
4,gleba,legolass,88,1
5,aquilo,aragorn,2,3`
)

func Test_CSV_Table_Parse(t *testing.T) {
	table, err := csv.Parse(strings.Split(header, ","), raw_table)
	require.NoError(t, err)
	ValidateTable(t, table, 5)
	require.Equal(t, table.Headers, []string{"amount", "planet", "player", "health", "speed"})
}

func Test_CSV_Table_Filter(t *testing.T) {
	table, err := csv.Parse(strings.Split(header, ","), raw_table)
	require.NoError(t, err)

	subtable, err := table.FilterColumns([]string{"health", "player", "planet"})
	require.NoError(t, err)
	ValidateTable(t, subtable, 3)
	require.Equal(t, subtable.Headers, []string{"health", "player", "planet"})

	subtable, err = table.FilterColumns([]string{"speed", "planet", "player"})
	require.NoError(t, err)
	ValidateTable(t, subtable, 3)
	require.Equal(t, subtable.Headers, []string{"speed", "planet", "player"})
}

func ValidateTable(t *testing.T, table csv.Table, columns int) {
	require.Len(t, table.Headers, columns)

	for r, row := range table.Records {
		require.Len(t, row.Values, columns)

		for c := range columns {
			v, ok := row.GetValue(c)
			require.True(t, ok, "[%v, %v]", c, r)
			require.NotEmpty(t, v, "[%v, %v]", c, r)
			require.Equal(t, v, row.Values[c])

			h, ok := table.GetHeader(c)
			require.True(t, ok)
			require.Equal(t, h, table.Headers[c])

			vh, vn, ok := table.GetRecordWithHeader(c, r)
			require.True(t, ok)
			require.Equal(t, vn, v)
			require.Equal(t, vh, h)
		}
	}
}