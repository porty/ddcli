package markdown

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTableWriter(t *testing.T) {
	b := bytes.Buffer{}
	w := NewTableWriter(&b)

	require.NoError(t, w.Write([]string{"ID", "Name"}))
	require.NoError(t, w.Write([]string{"1", "Freddy"}))
	require.NoError(t, w.Write([]string{"2", "Sarah"}))
	w.Flush()

	actual := b.String()
	expected := "| ID | Name |\n" +
		"|---|---|\n" +
		"| 1 | Freddy |\n" +
		"| 2 | Sarah |\n"

	require.Equal(t, expected, actual)
}
