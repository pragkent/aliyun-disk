package main

import (
	"reflect"
	"testing"
)

func TestFixArgs(t *testing.T) {
	tests := []struct {
		args []string
		want []string
	}{
		{
			[]string{"program", "a", "b", "-c"},
			[]string{"a", "b", "-c"},
		},
		{
			[]string{"program", "a", "-v"},
			[]string{"version", "a", "-v"},
		},
		{
			[]string{"program", "a", "-version"},
			[]string{"version", "a", "-version"},
		},
		{
			[]string{"program", "a", "--version"},
			[]string{"version", "a", "--version"},
		},
	}

	for _, c := range tests {
		got := fixArgs(c.args)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("fixArgs(%v) = %v; want %v", c.args, got, c.want)
		}
	}
}
