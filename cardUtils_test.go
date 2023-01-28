package CardUtils

import (
	"fmt"
	"github.com/mickstar/payment-card-utils-go/Scheme"
	"testing"
)

func Test_maskPan(t *testing.T) {
	type args struct {
		pan string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "pan is less than 10 digits",
			args: args{
				pan: "123456789",
			},
			want: "123456789",
		},
		{
			name: "pan is less than 16 digits",
			args: args{
				pan: "5300000000000000",
			},
			want: "530000******0000",
		},
		{
			name: "pan is less than 17 digits",
			args: args{
				pan: "53000000000000001",
			},
			want: "530000*******0001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaskPan(tt.args.pan); got != tt.want {
				t.Errorf("maskPan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maskPanWithCharacter(t *testing.T) {
	type args struct {
		pan           string
		maskCharacter uint8
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "pan is less than 10 digits",
			args: args{
				pan: "123456789",
			},
			want: "123456789",
		},
		{
			name: "pan is less than 16 digits",
			args: args{
				pan:           "5300000000000000",
				maskCharacter: 'X',
			},
			want: "530000XXXXXX0000",
		},
		{
			name: "pan is less than 17 digits",
			args: args{
				pan:           "53000000000000001",
				maskCharacter: 'Z',
			},
			want: "530000ZZZZZZZ0001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaskPanWithCharacter(tt.args.pan, tt.args.maskCharacter); got != tt.want {
				t.Errorf("maskPanWithCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidityCheck(t *testing.T) {
	type args struct {
		pan string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "pan is less than 14 digits",
			args: args{
				pan: "1234567890123",
			},
			want: false,
		},
		{
			name: "pan is greater than 19 digits",
			args: args{
				pan: "12345678901234567890",
			},
			want: false,
		},
		{
			name: "pan is not all digits",
			args: args{
				pan: "530000XXXXXX0000",
			},
			want: false,
		},
		{
			name: "pan is invalid",
			args: args{
				pan: "5300000000000000",
			},
			want: false,
		},
		{
			name: "pan is invalid 2",
			args: args{
				pan: "5300000000000001",
			},
			want: false,
		},
		{
			name: "pan is invalid",
			args: args{
				pan: "5300000000000002",
			},
			want: false,
		},
		{
			name: "pan is invalid",
			args: args{
				pan: "5300000000000003",
			},
			want: false,
		},
		{
			name: "pan is invalid",
			args: args{
				pan: "5300000000000004",
			},
			want: false,
		},
		{
			name: "pan is invalid",
			args: args{
				pan: "5300000000000005",
			},
			want: false,
		},
		{
			name: "pan is valid",
			args: args{
				pan: "5300000000000006",
			},
			want: true,
		},
		{
			name: "pan is invalid",
			args: args{
				pan: "5300000000000007",
			},
			want: false,
		},
		{
			name: "pan is invalid",
			args: args{
				pan: "5300000000000008",
			},
			want: false,
		},
		{
			name: "pan is invalid",
			args: args{
				pan: "5300000000000009",
			},
			want: false,
		},
		{
			name: "pan is valid",
			args: args{
				pan: "5300000000000055",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidityCheck(tt.args.pan); got != tt.want {
				t.Errorf("ValidityCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateRandomPanOfLength(t *testing.T) {
	s := GenerateRandomPanOfLength(16)
	fmt.Println("Generated pan", s)
	if len(s) != 16 {
		t.Errorf("Generated pan is not of length 16")
	}
	if !LuhnCheck(s) {
		t.Errorf("GenerateRandomPanOfLength() = %v, want a valid pan", s)
	}

	s = GenerateRandomPanOfLength(19)
	fmt.Println("Generated pan", s)
	if len(s) != 19 {
		t.Errorf("Generated pan is not of length 16")
	}
	if !LuhnCheck(s) {
		t.Errorf("GenerateRandomPanOfLength() = %v, want a valid pan", s)
	}
}

func TestGetCardIssuer(t *testing.T) {
	type args struct {
		pan string
	}
	tests := []struct {
		name string
		args args
		want Scheme.Scheme
	}{
		{
			name: "Unknown issuer",
			args: args{
				pan: "1234567890123",
			},
			want: Scheme.Unknown,
		},
		{
			name: "Mastercard",
			args: args{
				pan: "5100000000000000",
			},
			want: Scheme.MasterCard,
		},
		{
			name: "Not Mastercard",
			args: args{
				pan: "5600000000000000",
			},
			want: Scheme.Unknown,
		},
		{
			name: "Visa",
			args: args{
				pan: "4300000000000000",
			},
			want: Scheme.Visa,
		}, {
			name: "UnionPay",
			args: args{
				pan: "6200000000000000",
			},
			want: Scheme.UnionPay,
		},
		{
			name: "UnionPay 2",
			args: args{
				pan: "6250000000000000",
			},
			want: Scheme.UnionPay,
		},
		{
			name: "Not Union Pay",
			args: args{
				pan: "6251000000000000",
			},
			want: Scheme.Unknown,
		},
		{
			name: "Amex",
			args: args{
				pan: "370000000000000",
			},
			want: Scheme.AmericanExpress,
		},
		{
			name: "JCB",
			args: args{
				pan: "3528000000000000",
			},
			want: Scheme.JCB,
		},
		{
			name: "BP Card",
			args: args{
				pan: "70529000000000000",
			},
			want: Scheme.BPCard,
		},
		{
			name: "Diners Club",
			args: args{
				pan: "36000000000000",
			},
			want: Scheme.DinersClub,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCardScheme(tt.args.pan); got != tt.want {
				t.Errorf("GetCardIssuer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateRandomBPCardPan(t *testing.T) {
	s := generateRandomBPCardPan()
	if !LuhnCheck(s) {
		t.Errorf("generateRandomBPCardPan() = %v, want a valid pan", s)
	}
	if !ValidityCheck(s) {
		t.Errorf("generateRandomBPCardPan() = %v, want a valid pan", s)
	}

	if GetCardScheme(s) != Scheme.BPCard {
		t.Errorf("generateRandomBPCardPan() = %v, want a valid pan", s)
	}
}

func Test_generateRandomUnionPayPan(t *testing.T) {
	s := generateRandomUnionPayPan()
	if !LuhnCheck(s) {
		t.Errorf("generateRandomUnionPayPan() = %v, want a valid pan", s)
	}
	if !ValidityCheck(s) {
		t.Errorf("generateRandomUnionPayPan() = %v, want a valid pan", s)
	}

	if GetCardScheme(s) != Scheme.UnionPay {
		t.Errorf("generateRandomUnionPayPan() = %v, want a valid pan", s)
	}
}

func TestGenerateRandomPanOfScheme(t *testing.T) {
	for i := 0; i < 1_000; i++ {
		var schemes []Scheme.Scheme = []Scheme.Scheme{Scheme.MasterCard,
			Scheme.Visa,
			Scheme.AmericanExpress,
			Scheme.JCB,
			Scheme.DinersClub,
			Scheme.BPCard,
			Scheme.UnionPay,
			Scheme.Discover,
		}
		for _, s := range schemes {
			pan := GenerateRandomPanOfScheme(s)
			if !LuhnCheck(pan) {
				t.Errorf("GenerateRandomPanOfScheme() = %v, want a valid pan", pan)
			}
			if !ValidityCheck(pan) {
				t.Errorf("GenerateRandomPanOfScheme() = %v, want a valid pan", pan)
			}
			if GetCardScheme(pan) != s {
				t.Errorf("GenerateRandomPanOfScheme() = %v, want a valid pan", pan)
			}
		}

		s := GenerateRandomPanOfScheme(Scheme.Unknown)
		if !LuhnCheck(s) {
			t.Errorf("GenerateRandomPanOfScheme() = %v, want a valid pan", s)
		}
	}
}

func TestValidityCheck1(t *testing.T) {
	type args struct {
		pan string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "invalid visa",
			args: args{
				pan: "4111111111111112",
			},
			want: false,
		},
		{
			name: "invalid visa 2",
			args: args{
				pan: "41111111111111",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidityCheck(tt.args.pan); got != tt.want {
				t.Errorf("ValidityCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
