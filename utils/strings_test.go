package utils

import "testing"

func TestReplaceEmptyValue(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty no whitespace",
			args: args{
				value: "",
			},
			want: "N/A",
		},
		{
			name: "Not empty",
			args: args{
				value: "Chicken",
			},
			want: "Chicken",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceEmptyValue(tt.args.value); got != tt.want {
				t.Errorf("ReplaceEmptyValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
