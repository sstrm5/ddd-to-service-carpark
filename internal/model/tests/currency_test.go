package model_test

import (
	vo "dddcarpark/internal/model/valueobject"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCurrency_Valid(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  vo.Currency
	}{
		{
			name:  "rub",
			value: "rub",
			want:  vo.Rub,
		},
		{
			name:  "usd",
			value: "usd",
			want:  vo.Usd,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := vo.NewCurrency(tt.value)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestNewCurrency_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{
			name:  "Rub Invalid",
			value: "Rub",
		},
		{
			name:  "USD Invalid",
			value: "USD",
		},
		{
			name:  "usdt Invalid",
			value: "usdt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := vo.NewCurrency(tt.value)
			require.Error(t, err)
		})
	}
}
