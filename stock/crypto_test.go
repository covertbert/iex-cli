package stock

import "testing"

func TestQueryCrypto(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QueryCrypto()
		})
	}
}
