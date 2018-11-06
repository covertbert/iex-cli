package stock

import (
	"reflect"
	"testing"
)

var testChart = Chart{
	{
		Open:  1,
		Close: 1,
	},
	{
		Open:  2,
		Close: 2,
	},
	{
		Open:  3,
		Close: 3,
	},
}

func Test_dataPoints(t *testing.T) {
	type args struct {
		c Chart
	}
	tests := []struct {
		name string
		args args
		want map[string][]float64
	}{
		{
			name: "Demo",
			args: args{
				c: testChart,
			},
			want: map[string][]float64{
				"first":  {1, 2, 3},
				"second": {1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dataPoints(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dataPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataLabels(t *testing.T) {
	type args struct {
		c Chart
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dataLabels(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dataLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
