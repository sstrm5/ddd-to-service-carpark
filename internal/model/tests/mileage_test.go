package model_test

import (
	vo "dddcarpark/internal/model/valueobject"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewMileage_Valid(t *testing.T) {
	tests := []struct {
		name     string
		distance uint32
		measure  vo.Measure
		want     vo.Mileage
	}{
		{
			name:     "15 km",
			distance: 15,
			measure:  vo.Km,
			want: vo.Mileage{
				Distance: 15,
				Measure:  vo.Km,
			},
		},
		{
			name:     "150 miles",
			distance: 150,
			measure:  vo.Miles,
			want: vo.Mileage{
				Distance: 150,
				Measure:  vo.Miles,
			},
		},
		{
			name:     "0 miles",
			distance: 0,
			measure:  vo.Miles,
			want: vo.Mileage{
				Distance: 0,
				Measure:  vo.Miles,
			},
		},
		{
			name:     "0 km",
			distance: 0,
			measure:  vo.Km,
			want: vo.Mileage{
				Distance: 0,
				Measure:  vo.Km,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := vo.NewMileage(tt.distance, tt.measure)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestNewMileage_Invalid(t *testing.T) {
	tests := []struct {
		name     string
		distance uint32
		measure  vo.Measure
	}{
		{
			name:     "distance more 10_000_000, km",
			distance: 10_000_001,
			measure:  vo.Km,
		},
		{
			name:     "distance more 10_000_000, miles",
			distance: 10_000_001,
			measure:  vo.Miles,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := vo.NewMileage(tt.distance, tt.measure)
			require.Error(t, err)
		})
	}
}

func TestLessThan(t *testing.T) {

	mileage1, _ := vo.NewMileage(1, vo.Km)
	mileage2, _ := vo.NewMileage(1, vo.Miles)
	mileage3, _ := vo.NewMileage(2, vo.Km)
	mileage4, _ := vo.NewMileage(200_000, vo.Km)

	tests := []struct {
		name     string
		mileage1 vo.Mileage
		mileage2 vo.Mileage
		want     bool
	}{
		{
			name:     "test1",
			mileage1: mileage1,
			mileage2: mileage2,
			want:     true,
		},
		{
			name:     "test2",
			mileage1: mileage2,
			mileage2: mileage1,
			want:     false,
		},
		{
			name:     "test3",
			mileage1: mileage1,
			mileage2: mileage4,
			want:     true,
		},
		{
			name:     "test4",
			mileage1: mileage2,
			mileage2: mileage3,
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.mileage1.LessThan(tt.mileage2)
			require.Equal(t, tt.want, got)
		})
	}
}
