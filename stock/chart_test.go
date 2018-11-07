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

func Test_queryPath(t *testing.T) {
	type args struct {
		symbol string
		rng    string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "20180129",
			args: args{
				rng: "20180129",
			},
			want:    "/stock//chart/date/20180129",
			wantErr: false,
		},
		{
			name: "1405gg4146",
			args: args{
				rng: "1405gg4146",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "14055441463",
			args: args{
				rng: "14055441463",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "chicken",
			args: args{
				rng: "chicken",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "5yy",
			args: args{
				rng: "5yy",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "5y",
			args: args{
				rng: "5y",
			},
			want:    "/stock//chart/5y",
			wantErr: false,
		},
		{
			name: "2y",
			args: args{
				rng: "2y",
			},
			want:    "/stock//chart/2y",
			wantErr: false,
		},
		{
			name: "1y",
			args: args{
				rng: "1y",
			},
			want:    "/stock//chart/1y",
			wantErr: false,
		},
		{
			name: "ytd",
			args: args{
				rng: "ytd",
			},
			want:    "/stock//chart/ytd",
			wantErr: false,
		},
		{
			name: "6m",
			args: args{
				rng: "6m",
			},
			want:    "/stock//chart/6m",
			wantErr: false,
		},
		{
			name: "3m",
			args: args{
				rng: "3m",
			},
			want:    "/stock//chart/3m",
			wantErr: false,
		},
		{
			name: "1m",
			args: args{
				rng: "1m",
			},
			want:    "/stock//chart/1m",
			wantErr: false,
		},
		{
			name: "1d",
			args: args{
				rng: "1d",
			},
			want:    "/stock//chart/1d",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := queryPath(tt.args.symbol, tt.args.rng)
			if (err != nil) != tt.wantErr {
				t.Errorf("queryPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("queryPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
