package csv

import (
	"fmt"
	"io"
	"reflect"
)

// StructScanner provides access to the fields of CSV-encoded
// data through a struct's fields.
//
// Like unmarshaling with the standard JSON or XML decoders, the
// fields of the struct must be exported and tagged with a `"csv:"`
// prefix.
//
// All configurations of the underlying *csv.Reader are available
// through an [Option].
type StructScanner struct {
	*ColumnScanner
}

// NewStructScanner returns a StructScanner that reads from reader,
// configured with the provided options.
func NewStructScanner(reader io.Reader, options ...Option) (*StructScanner, error) {
	inner, err := NewColumnScanner(NewScanner(reader, options...))
	if err != nil {
		return nil, err
	}
	return &StructScanner{ColumnScanner: inner}, nil
}

// Populate gets the most recent record generated by a call to Scan
// and stores the values for tagged fields in the value pointed to
// by v.
func (this *StructScanner) Populate(v interface{}) error {
	type_ := reflect.TypeOf(v)
	if type_.Kind() != reflect.Ptr {
		return fmt.Errorf("Provided value must be reflect.Ptr. You provided [%v] ([%v]).", v, type_.Kind())
	}

	value := reflect.ValueOf(v)
	if value.IsNil() {
		return fmt.Errorf("The provided value was nil. Please provide a non-nil pointer.")
	}

	this.populate(type_.Elem(), value.Elem())
	return nil
}

func (this *StructScanner) populate(type_ reflect.Type, value reflect.Value) {
	for x := 0; x < type_.NumField(); x++ {
		column := type_.Field(x).Tag.Get("csv")

		_, found := this.columnIndex[column]
		if !found {
			continue
		}

		field := value.Field(x)
		if field.Kind() != reflect.String {
			continue // Future: return err?
		} else if !field.CanSet() {
			continue // Future: return err?
		}

		field.SetString(this.Column(column))
	}
}
