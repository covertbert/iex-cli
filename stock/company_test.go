package stock

import "testing"

func Test_replaceEmpty(t *testing.T) {
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
			if got := replaceEmpty(tt.args.value); got != tt.want {
				t.Errorf("replaceEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
