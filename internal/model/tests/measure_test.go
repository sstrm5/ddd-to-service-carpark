package model_test

import (
	vo "dddcarpark/internal/model/valueobject"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewMeasure_Valid(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  vo.Measure
	}{
		{name: "km", value: "km", want: vo.Km},
		{name: "miles", value: "miles", want: vo.Miles},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := vo.NewMeasure(tt.value)
			require.NoError(t, err)
			require.Equal(t, got, tt.want)
		})
	}
}

func TestNewMeasure_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{
			name:  "test1 Invalid",
			value: "KM",
		},
		{
			name:  "test2 Invalid",
			value: "MILES",
		},
		{
			name:  "test3 empty Invalid",
			value: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := vo.NewMeasure(tt.value)
			require.Error(t, err)
		})
	}
}
