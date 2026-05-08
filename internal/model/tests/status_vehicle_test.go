package model_test

import (
	vo "dddcarpark/internal/model/valueobject"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewStatusVehicle_Valid(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  vo.StatusVehicle
	}{
		{
			name:  "active",
			value: "active",
			want:  vo.Active,
		},
		{
			name:  "under repair",
			value: "under repair",
			want:  vo.UnderRepair,
		},
		{
			name:  "written off",
			value: "written off",
			want:  vo.WrittenOff,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := vo.NewStatusVehicle(tt.value)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestNewStatusVehicle_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{
			name:  "ACTIVE Invalid",
			value: "ACTIVE",
		},
		{
			name:  "under_repair Invalid",
			value: "under_repair",
		},
		{
			name:  "written off Invalid",
			value: "some_status",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := vo.NewStatusVehicle(tt.value)
			require.Error(t, err)
		})
	}
}
