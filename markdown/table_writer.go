package markdown

import (
	"io"
	"strings"
)

type TableWriter struct {
	w             io.Writer
	headerWritten bool
}

func NewTableWriter(w io.Writer) *TableWriter {
	return &TableWriter{
		w: w,
	}
}

func (t *TableWriter) Write(record []string) error {
	line := "| " + strings.Join(record, " | ") + " |\n"
	if _, err := t.w.Write([]byte(line)); err != nil {
		return err
	}
	if t.headerWritten {
		return nil
	}
	t.headerWritten = true

	dashes := make([]string, len(record))
	for i := range dashes {
		dashes[i] = "---"
	}
	line = "|" + strings.Join(dashes, "|") + "|\n"
	if _, err := t.w.Write([]byte(line)); err != nil {
		return err
	}
	return nil
}

func (t *TableWriter) Flush() {
}
