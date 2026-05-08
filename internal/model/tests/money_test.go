package model_test

import (
	vo "dddcarpark/internal/model/valueobject"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewMoney_Valid(t *testing.T) {
	tests := []struct {
		name     string
		value    uint32
		currency vo.Currency
		want     vo.Money
	}{
		{
			name:     "test1 zero money rub",
			value:    0,
			currency: vo.Rub,
			want: vo.Money{
				Value:    0,
				Currency: vo.Rub,
			},
		},
		{
			name:     "test2 zero money usd",
			value:    0,
			currency: vo.Usd,
			want: vo.Money{
				Value:    0,
				Currency: vo.Usd,
			},
		},
		{
			name:     "test3",
			value:    150000,
			currency: vo.Usd,
			want: vo.Money{
				Value:    150000,
				Currency: vo.Usd,
			},
		},
		{
			name:     "test4",
			value:    150000,
			currency: vo.Rub,
			want: vo.Money{
				Value:    150000,
				Currency: vo.Rub,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := vo.NewMoney(tt.value, tt.currency)
			require.Equal(t, tt.want, got)
			require.NoError(t, err)
		})
	}
}

func TestNewMoney_Invalid(t *testing.T) {
	tests := []struct {
		name     string
		value    uint32
		currency vo.Currency
		want     vo.Money
	}{
		{
			name:     "empty currency",
			value:    999_999,
			currency: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := vo.NewMoney(tt.value, tt.currency)
			require.Error(t, err)
		})
	}
}
