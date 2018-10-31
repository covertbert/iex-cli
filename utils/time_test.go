package utils

import "testing"

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
