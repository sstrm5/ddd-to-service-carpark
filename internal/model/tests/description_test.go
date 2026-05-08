package model_test

import (
	vo "dddcarpark/internal/model/valueobject"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDescription_Valid(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  vo.Description
	}{
		{
			name:  "test1",
			value: "text",
			want:  vo.Description("Text"),
		},
		{
			name:  "test2",
			value: "text Text",
			want:  vo.Description("Text text"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := vo.NewDescription(tt.value)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestNewDescription_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{
			name:  "empty Invalid",
			value: "",
		},
		{
			name:  "1 symbol",
			value: "a",
		},
		{
			name:  "too long",
			value: strings.Repeat("a", 301),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := vo.NewDescription(tt.value)
			require.Error(t, err)
		})
	}
}
