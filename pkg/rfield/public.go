package rfield

import "strings"

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
