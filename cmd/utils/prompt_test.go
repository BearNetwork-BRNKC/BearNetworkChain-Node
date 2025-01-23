

// Package utils contains internal helper functions for go-ethereum commands.
package utils

import (
	"testing"
)

func TestGetPassPhraseWithList(t *testing.T) {
	t.Parallel()
	type args struct {
		text         string
		confirmation bool
		index        int
		passwords    []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{
				"text1",
				false,
				0,
				[]string{"zero", "one", "two"},
			},
			"zero",
		},
		{
			"test2",
			args{
				"text2",
				false,
				5,
				[]string{"zero", "one", "two"},
			},
			"two",
		},
		{
			"test3",
			args{
				"text3",
				true,
				1,
				[]string{"zero", "one", "two"},
			},
			"one",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := GetPassPhraseWithList(tt.args.text, tt.args.confirmation, tt.args.index, tt.args.passwords); got != tt.want {
				t.Errorf("GetPassPhraseWithList() = %v, want %v", got, tt.want)
			}
		})
	}
}
