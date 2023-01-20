package Scheme

type Scheme string

const (
	Visa            Scheme = "Visa"
	MasterCard             = "MasterCard"
	Discover               = "Discover"
	AmericanExpress        = "American Express"
	DinersClub             = "Diners Club"
	JCB                    = "JCB"
	Unknown                = "Unknown"
	BPCard                 = "BP Card"
	UnionPay               = "UnionPay"
)

func LengthCheckForScheme(scheme Scheme, length int) bool {
	switch scheme {
	case AmericanExpress:
		return length == 15
	case Visa:
		return length == 16 || length == 13
	case MasterCard:
		return length == 16
	case Discover, JCB, UnionPay:
		return length >= 16 && length <= 19
	case DinersClub:
		return length >= 14 && length <= 19
	case BPCard:
		return length == 19
	default:
		return false
	}
}
