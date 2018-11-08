package stock

import (
	"reflect"
	"testing"

	"github.com/rendon/testcli"
)

var testDividends = Dividends{
	{
		Amount: 1,
		ExDate: "2018-11-05",
	},
	{
		Amount: 2,
		ExDate: "2018-11-05",
	},
	{
		Amount: 3,
		ExDate: "2018-11-05",
	},
}

func Test_divDataLabels(t *testing.T) {
	type args struct {
		c Dividends
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Data labels",
			args: args{
				c: testDividends,
			},
			want: []string{"2018-11-05", "2018-11-05", "2018-11-05", "2018-11-05", "2018-11-05", "2018-11-05", "2018-11-05", "2018-11-05", "2018-11-05"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divDataLabels(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divDataLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_divQueryPath(t *testing.T) {
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
			want:    "/stock//dividends/5y",
			wantErr: false,
		},
		{
			name: "2y",
			args: args{
				rng: "2y",
			},
			want:    "/stock//dividends/2y",
			wantErr: false,
		},
		{
			name: "1y",
			args: args{
				rng: "1y",
			},
			want:    "/stock//dividends/1y",
			wantErr: false,
		},
		{
			name: "ytd",
			args: args{
				rng: "ytd",
			},
			want:    "/stock//dividends/ytd",
			wantErr: false,
		},
		{
			name: "6m",
			args: args{
				rng: "6m",
			},
			want:    "/stock//dividends/6m",
			wantErr: false,
		},
		{
			name: "3m",
			args: args{
				rng: "3m",
			},
			want:    "/stock//dividends/3m",
			wantErr: false,
		},
		{
			name: "1m",
			args: args{
				rng: "1m",
			},
			want:    "/stock//dividends/1m",
			wantErr: false,
		},
		{
			name: "1d",
			args: args{
				rng: "1d",
			},
			want:    "/stock//dividends/1d",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := divQueryPath(tt.args.symbol, tt.args.rng)
			if (err != nil) != tt.wantErr {
				t.Errorf("divQueryPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("divQueryPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDividendsNoArgs(t *testing.T) {
	testcli.Run("../iex-cli", "dividends")

	if !testcli.StdoutContains("Error: No argument supplied") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Error: No argument supplied")
	}
}
