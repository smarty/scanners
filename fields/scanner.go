// Package fields scans fields, splitting on whitespace—fields
// themselves cannot contain whitespace.
//
// Advance the scanner with the Scan method and check errors with
// the Err method, both from the underlying bufio.Scanner.
package fields

import (
	"bufio"
	"io"
	"strings"
)

// Scanner provides access to the whitespace-separated fields of
// data. Field values cannot contain any whitespace.
//
// For a file that follows the encoding scheme of a so-called TSV,
// use [github.com/smarty/scanners/csv.Scanner] and configure it
// for tabs with [github.com/smarty/scanners/csv.Comma].
type Scanner struct {
	*bufio.Scanner
}

// NewScanner returns a fields scanner.
func NewScanner(reader io.Reader) *Scanner {
	return &Scanner{Scanner: bufio.NewScanner(reader)}
}

// Fields returns the most recent fields generated by a call to Scan as a
// []string.
func (this *Scanner) Fields() []string {
	return strings.Fields(this.Text())
}
