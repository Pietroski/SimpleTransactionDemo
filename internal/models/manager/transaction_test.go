package manager_models

import "testing"

func TestCryptoCurrencies_String(t *testing.T) {
	tests := []struct {
		name string
		cc   CryptoCurrencies
		want string
	}{
		{
			name: "bitcoin test case",
			cc:   "BITCOIN",
			want: "BITCOIN",
		},
		{
			name: "bitcoin test case",
			cc:   CryptoCurrenciesBITCOIN,
			want: CryptoCurrenciesBITCOIN.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cc.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCryptoCurrencies_IsCryptoCurrency(t *testing.T) {
	tests := []struct {
		name string
		cc   CryptoCurrencies
		want bool
	}{
		{
			name: "Pietroski-coin test case",
			cc:   CryptoCurrenciesPIETROSKICOIN,
			want: true,
		},
		{
			name: "dodge-coin test case",
			cc:   CryptoCurrenciesDODGECOIN,
			want: true,
		},
		{
			name: "Bitcoin test case",
			cc:   CryptoCurrenciesBITCOIN,
			want: true,
		},
		{
			name: "ethereum test case",
			cc:   CryptoCurrenciesETHEREUM,
			want: true,
		},
		{
			name: "failure test case",
			cc:   "invalid-coin",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cc.IsCryptoCurrency(); got != tt.want {
				t.Errorf("IsCryptoCurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}
