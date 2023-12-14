package generator

import (
    "crypto/rand"
    "math/big"
    "strings"
)

// Opciones de caracteres para la generación de contraseñas
const (
    letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    lettersEasy   = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
    numbers       = "0123456789"
    symbols       = "!@#$%^&*()-_=+[]{}|;:,.<>/?"
    numbersAndSyms = numbers + symbols
)

// GeneratePassword genera una contraseña basada en los parámetros dados.
func GeneratePassword(length int, includeNums bool, includeSyms bool, easyToSay bool, easyToRead bool) string {
    var charset string

    if easyToSay || easyToRead {
        charset = lettersEasy // Usa un conjunto de caracteres sin los confusos.
    } else {
        charset = letters
    }

    if includeNums {
        charset += numbers
    }

    if includeSyms {
        charset += symbols
    }

    var password strings.Builder
    for i := 0; i < length; i++ {
        charIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
        if err != nil {
            continue // Si hay un error, simplemente omite este carácter.
        }
        password.WriteRune(rune(charset[charIndex.Int64()]))
    }
    return password.String()
}
