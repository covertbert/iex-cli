package stock

import (
	"reflect"
	"testing"
)

var testChart = Chart{
	{
		Open:  1,
		Close: 1,
		High:  1,
		Low:   1,
		Date:  "2018-11-05",
	},
	{
		Open:  2,
		Close: 2,
		High:  2,
		Low:   2,
		Date:  "2018-11-05",
	},
	{
		Open:  3,
		Close: 3,
		High:  3,
		Low:   3,
		Date:  "2018-11-05",
	},
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
		{
			name: "Data labels",
			args: args{
				c: testChart,
			},
			want: []string{"2018-11-05", "2018-11-05", "2018-11-05", "2018-11-05", "2018-11-05", "2018-11-05"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dataLabels(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dataLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
