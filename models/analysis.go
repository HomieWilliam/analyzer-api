package models
 
type AnalyzerRequest struct {
    Phrase string `json:"phrase"`
}
 
type AnalyzerResponse struct {
    Words      int `json:"words"`
    Vowels     int `json:"vowels"`
    Consonants int `json:"consonants"`
}