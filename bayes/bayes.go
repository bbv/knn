package bayes

import (
    "fmt"
    "regexp"
)

type WordStat struct {
    Prob float64
    Occurrencies int
}

type WordsMap map[string]WordStat

type BayesClassifier struct {
    Words WordsMap
    DocNumber int
    DocFrequency float64
    Name string
}

func NewBayesClassifier(name string) BayesClassifier {
    return BayesClassifier{
        Name: name,
        Words: make(WordsMap),
    }
}

func (b *BayesClassifier) Learn(text string, good bool) {
    b.DocNumber += 1

    words := filterWords(splitText(text))
    uniqueWords := make(map[string]int)

    fmt.Printf("%#v %d\n", words, len(words))
    for _, word := range words {
        if _, ok := uniqueWords[word]; !ok {
            uniqueWords[word] += 1
        }
    }

    for word := range uniqueWords {
        fmt.Println(word)
        wordStat, ok := b.Words[word]
        fmt.Println(wordStat)
        if !ok {
            b.Words[word] = wordStat
        }
        wordStat.Occurrencies += 1
        wordStat.Prob = float64(wordStat.Occurrencies / b.DocNumber)
        b.Words[word] = wordStat
    }

    return
}

func splitText(text string) []string {
    // s := regexp.MustCompile("[^\\p{L}\\-]+").Split(text, -1)
    s := regexp.MustCompile("[^\\p{L}\\-]+").Split(text, -1)
    return s
}

func filterWords(words []string) []string {
    res := make([]string, 0)
    for _, v := range words {
        if len(v) > 2 {
            res = append(res, v)
        }
    }
    return res
}
