package services

import (
    "github.com/HomieWilliam/analyzer-api/models"
    "strings"
    "unicode"
)

func normalizeRune(r rune) rune {
    switch r {
    case 'á', 'à', 'ã', 'â', 'ä', 'Á', 'À', 'Ã', 'Â', 'Ä':
        return 'a'
    case 'é', 'è', 'ê', 'ë', 'É', 'È', 'Ê', 'Ë':
        return 'e'
    case 'í', 'ì', 'î', 'ï', 'Í', 'Ì', 'Î', 'Ï':
        return 'i'
    case 'ó', 'ò', 'õ', 'ô', 'ö', 'Ó', 'Ò', 'Õ', 'Ô', 'Ö':
        return 'o'
    case 'ú', 'ù', 'û', 'ü', 'Ú', 'Ù', 'Û', 'Ü':
        return 'u'
    case 'ç', 'Ç':
        return 'c'
    default:
        return r
    }
}

func Analyze(phrase string) models.AnalyzerResponse {
    phrase = strings.ToLower(phrase)

    words := len(strings.Fields(phrase))
    vowels, consonants := 0, 0

    for _, char := range phrase {
        if unicode.IsLetter(char) {
            normalized := normalizeRune(char)
            if strings.ContainsRune("aeiou", normalized) {
                vowels++
            } else {
                consonants++
            }
        }
    }
    
    return models.AnalyzerResponse{
        Words:      words,
        Vowels:     vowels,
        Consonants: consonants,
    }
}
