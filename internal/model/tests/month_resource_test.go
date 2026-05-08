package model_test

import (
	vo "dddcarpark/internal/model/valueobject"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewResource_Valid(t *testing.T) {
	tests := []struct {
		name  string
		value uint32
		want  vo.MonthResource
	}{
		{
			name:  "test1",
			value: 1,
			want:  vo.MonthResource(1),
		},
		{
			name:  "test2",
			value: 10,
			want:  vo.MonthResource(10),
		},
		{
			name:  "test3",
			value: 1000,
			want:  vo.MonthResource(1000),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := vo.NewMonthResource(tt.value)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestNewMonthResource_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		value uint32
	}{
		{
			name:  "value less than 1 Invalid",
			value: 0,
		},
		{
			name:  "value more than 1200 Invalid",
			value: 1201,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := vo.NewMonthResource(tt.value)
			require.Error(t, err)
		})
	}
}
