package services
 
import (
    "github.com/HomieWilliam/analyzer-api/models"
    "strings"
    "unicode"
)
 
func Analyze(phrase string) models.AnalyzerResponse {
    phrase = strings.ToLower(phrase)
 
    // Contar palavras
    words := len(strings.Fields(phrase))
 
    // Contar vogais e consoantes
    vowels, consonants := 0, 0
    for _, char := range phrase {
        if unicode.IsLetter(char) {
            if strings.ContainsRune("aeiou", char) {
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