package utils

import (
	"testing"
)

func TestUNIXToHumanReadable(t *testing.T) {
	type args struct {
		u int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "UNIX to 24hr",
			args: args{
				u: 1540906290094,
			},
			want: "13:01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UNIXToHumanReadable(tt.args.u); got != tt.want {
				t.Errorf("UNIXToHumanReadable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateStringToHumanReadable(t *testing.T) {
	type args struct {
		d string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Date string",
			args: args{
				d: "2018-11-02T11:01:00-04:00",
			},
			want: "Fri, 02 Nov 2018 11:01:00 -0400",
		},
		{
			name: "Date string",
			args: args{
				d: "2018-11-02T12:00:00-04:00",
			},
			want: "Fri, 02 Nov 2018 12:00:00 -0400",
		},
		{
			name: "Date string",
			args: args{
				d: "2018-11-02T11:30:26-04:00",
			},
			want: "Fri, 02 Nov 2018 11:30:26 -0400",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DateStringToHumanReadable(tt.args.d); got != tt.want {
				t.Errorf("DateStringToHumanReadable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimYearFromDate(t *testing.T) {
	type args struct {
		d string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Date string",
			args: args{
				d: "2018-11-12",
			},
			want: "12/11",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShortDate(tt.args.d); got != tt.want {
				t.Errorf("ShortDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
