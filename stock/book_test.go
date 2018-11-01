package stock

import "testing"

func TestQueryBook(t *testing.T) {
	type args struct {
		ticker     string
		subsection string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "AAPL - quote",
			args: args{
				ticker:     "AAPL",
				subsection: "quote",
			},
		},
		{
			name: "AAPL - bids",
			args: args{
				ticker:     "AAPL",
				subsection: "bids",
			},
		},
		{
			name: "AAPL - asks",
			args: args{
				ticker:     "AAPL",
				subsection: "asks",
			},
		},
		{
			name: "AAPL - trades",
			args: args{
				ticker:     "AAPL",
				subsection: "trades",
			},
		},
		{
			name: "QQQ - quote",
			args: args{
				ticker:     "QQQ",
				subsection: "quote",
			},
		},
		{
			name: "QQQ - bids",
			args: args{
				ticker:     "QQQ",
				subsection: "bids",
			},
		},
		{
			name: "QQQ - asks",
			args: args{
				ticker:     "QQQ",
				subsection: "asks",
			},
		},
		{
			name: "QQQ - trades",
			args: args{
				ticker:     "QQQ",
				subsection: "trades",
			},
		},
		{
			name: "SYX - quote",
			args: args{
				ticker:     "SYX",
				subsection: "quote",
			},
		},
		{
			name: "SYX - bids",
			args: args{
				ticker:     "SYX",
				subsection: "bids",
			},
		},
		{
			name: "SYX - asks",
			args: args{
				ticker:     "SYX",
				subsection: "asks",
			},
		},
		{
			name: "SYX - trades",
			args: args{
				ticker:     "SYX",
				subsection: "trades",
			},
		},
		{
			name: "KEM - quote",
			args: args{
				ticker:     "KEM",
				subsection: "quote",
			},
		},
		{
			name: "KEM - bids",
			args: args{
				ticker:     "KEM",
				subsection: "bids",
			},
		},
		{
			name: "KEM - asks",
			args: args{
				ticker:     "KEM",
				subsection: "asks",
			},
		},
		{
			name: "KEM - trades",
			args: args{
				ticker:     "KEM",
				subsection: "trades",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QueryBook(tt.args.ticker, tt.args.subsection)
		})
	}
}
