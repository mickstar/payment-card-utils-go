# payment-card-utils-go

Small utility library for validating and formatting payment card numbers.  
This library supports  
* Masking PANs
* Validating PANs
* Getting Issuer
* Generating Random PANs

## Usage

    go get github.com/mickstar/payment-card-utils-go v0.1.0

Masking pans

    pan := "4111111111111111"
    maskedPan := CardUtils.MaskPan(pan)
    fmt.Println(maskedPan) // 411111******1111
    fmt.Println(CardUtils.MaskPanWithCharacter(pan, 'X')) // 411111XXXXXX1111

Validating pans

    pan := "4111111111111111"
    // useful if you just want a Luhn Check
    fmt.Println(CardUtils.LuhnCheck(pan)) // true
    // this also verifies card length
    fmt.Println(CardUtils.VerifyPan(pan)) // true

Getting Scheme

        pan := "4111111111111111"
        fmt.Println(CardUtils.GetScheme(pan)) // VISA
    
        switch CardUtils.GetScheme(pan) {
        case Scheme.VISA:
            fmt.Println("Visa")
        case Scheme.MasterCard:
            fmt.Println("Mastercard")
        // etc
        default:
            fmt.Println("something else.")
        }

Generating Random PANs
    
        fmt.Println(CardUtils.GenerateRandomPan() // 1274574654654654


## Author
Michael Johnston <michael.johnston29@gmail.com>

## Licence
Released under MIT

