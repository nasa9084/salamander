package model_test

import (
	"reflect"
	"testing"

	"github.com/nasa9084/salamander/salamander/model"
)

type withoutTag struct {
	ColumnA string
	ColumnB string
	columnc string
}

type withTag struct {
	ColumnA string `col:"col_a"`
	ColumnB string `col:"col_b"`
	columnc string `col:"col_c"`
}

type nested struct {
	ColumnA string
	Nested  struct {
		ColumnB string
		columnc string
	}
	columnd string
}

func TestColumns(t *testing.T) {
	candidates := []struct {
		in       interface{}
		expected []string
	}{
		{withoutTag{}, []string{"column_a", "column_b"}},
		{withTag{}, []string{"col_a", "col_b"}},
		{nested{}, []string{"column_a", "column_b"}},
		{0, nil},
	}
	for _, c := range candidates {
		cols := model.Columns(c.in)
		if !reflect.DeepEqual(cols, c.expected) {
			t.Errorf("%s != %s", cols, c.expected)
			return
		}
	}
}
