package model_test

import (
	vo "dddcarpark/internal/model/valueobject"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewItemType_Valid(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  vo.ItemType
	}{
		{name: "work", value: "work", want: vo.Work},
		{name: "component", value: "component", want: vo.Component},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := vo.NewItemType(tt.value)
			require.NoError(t, err)
			require.Equal(t, got, tt.want)
		})
	}
}

func TestNewItemType_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{name: "work1 Invalid", value: "WORK"},
		{name: "work2 Invalid", value: "Work"},
		{name: "component Invalid", value: "COMPONENT"},
		{name: "not existing Invalid", value: "installing"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := vo.NewItemType(tt.value)
			require.Error(t, err)
		})
	}
}
