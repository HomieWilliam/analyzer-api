package services
 
import (
    "./models"
    "reflect"
    "testing"
)
 
func TestAnalyze(t *testing.T) {
    tests := []struct {
        name   string
        input  string
        output models.AnalyzerResponse
    }{
        {
            name:  "Frase simples",
            input: "Go é bom",
            output: models.AnalyzerResponse{
                Words:      3,
                Vowels:     3, // o, e, o
                Consonants: 3, // g, b, m
            },
        },
        {
            name:  "Frase com várias palavras",
            input: "Go é uma linguagem poderosa",
            output: models.AnalyzerResponse{
                Words:      5,
                Vowels:     11,
                Consonants: 12,
            },
        },
        {
            name:  "Frase vazia",
            input: "",
            output: models.AnalyzerResponse{
                Words:      0,
                Vowels:     0,
                Consonants: 0,
            },
        },
        {
            name:  "Frase só consoantes",
            input: "bcdfg",
            output: models.AnalyzerResponse{
                Words:      1,
                Vowels:     0,
                Consonants: 5,
            },
        },
        {
            name:  "Frase só vogais",
            input: "aeiou",
            output: models.AnalyzerResponse{
                Words:      1,
                Vowels:     5,
                Consonants: 0,
            },
        },
    }
 
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Analyze(tt.input)
            if !reflect.DeepEqual(got, tt.output) {
                t.Errorf("esperado %+v, obtido %+v", tt.output, got)
            }
        })
    }
}
 