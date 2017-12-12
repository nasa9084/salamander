package model

import (
	"reflect"
	"unicode"

	strlib "github.com/nasa9084/go-strlib"
)

const colTag = "col"

// Columns extract RDB's database column names from go's struct.
// This function returns nil if given is not struct.
func Columns(m interface{}) []string {
	t := reflect.TypeOf(m)
	if t.Kind() != reflect.Struct {
		return nil
	}
	return columns(t)
}

func columns(t reflect.Type) []string {
	cols := []string{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !isExported(f) {
			continue
		}

		var adds []string
		if f.Type.Kind() == reflect.Struct {
			adds = columns(f.Type)
		} else {
			adds = []string{fieldName(f)}
		}

		cols = append(cols, adds...)
	}
	return cols
}

func isExported(f reflect.StructField) bool {
	return unicode.IsUpper([]rune(f.Name)[0])
}

func fieldName(f reflect.StructField) string {
	tagv, ok := f.Tag.Lookup(colTag)
	if ok {
		return tagv
	}
	return strlib.SnakeCase(f.Name)
}
