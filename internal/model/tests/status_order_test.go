package model_test

import (
	vo "dddcarpark/internal/model/valueobject"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewStatusOrder_Valid(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  vo.StatusOrder
	}{
		{
			name:  "draft",
			value: "draft",
			want:  vo.Draft,
		},
		{
			name:  "completed",
			value: "completed",
			want:  vo.Completed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := vo.NewStatusOrder(tt.value)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestNewStatusOrder_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{
			name:  "test1 empty",
			value: "",
		},
		{
			name:  "test2",
			value: "active",
		},
		{
			name:  "test3",
			value: "DRAFT",
		},
		{
			name:  "test4",
			value: "Completed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := vo.NewStatusOrder(tt.value)
			require.Error(t, err)
		})
	}
}
