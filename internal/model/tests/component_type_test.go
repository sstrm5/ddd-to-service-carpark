package model_test

import (
	vo "dddcarpark/internal/model/valueobject"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewComponentType_Valid(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  vo.ComponentType
	}{
		{name: "Oil Filter", value: "oil_filter", want: vo.OilFilter},
		{name: "Air Filter", value: "air_filter", want: vo.AirFilter},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := vo.NewComponentType(tt.value)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestNewComponentType_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{name: "Oil Filter Invalid", value: "OIL_FILTER"},
		{name: "Air Filter Invalid", value: "air filter"},
		{name: "Not existing Invalid", value: "wheel"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := vo.NewComponentType(tt.value)
			require.Error(t, err)
		})
	}
}
