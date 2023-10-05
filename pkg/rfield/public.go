package rfield

import (
	"reflect"
	"strings"
	"unicode"
)

func (f *RField) IsPublic() bool {
	return unicode.IsUpper(rune(f.field.Name[0]))
}

func (f *RField) Type() reflect.Type {
	return f.field.Type
}

func (f *RField) Tag(tagName string) []string {
	tagValue := f.field.Tag.Get(tagName)
	if tagValue == "" {
		return nil
	}

	values := strings.Split(tagValue, ",")
	for i := 0; i < len(values); i++ {
		values[i] = strings.TrimSpace(values[i])
	}
	return values
}
