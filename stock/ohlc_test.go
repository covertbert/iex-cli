package stock

import "testing"

func TestQueryOHLC(t *testing.T) {
	type args struct {
		ticker string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "AAPL",
			args: args{
				ticker: "AAPL",
			},
		},
		{
			name: "QQQ",
			args: args{
				ticker: "QQQ",
			},
		},
		{
			name: "SYX",
			args: args{
				ticker: "SYX",
			},
		},
		{
			name: "KEM",
			args: args{
				ticker: "KEM",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QueryOHLC(tt.args.ticker)
		})
	}
}
