package bayes

import (
    "fmt"
    "regexp"
    "math"
)

const defaultThreshold = 3
const defaultProb = 0.5

type WordStat struct {
    Prob float64
    Occurrencies int
}

func (ws WordStat) CorrectedProb() float64 {
    if ws.Occurrencies < defaultThreshold {
        return defaultProb
    }
    return ws.Prob
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
        if !ok {
            fmt.Println("Creating stats for word: ", word)
            b.Words[word] = wordStat
        }
        if good {
            wordStat.Occurrencies += 1
        }
        wordStat.Prob = float64(wordStat.Occurrencies) / float64(b.DocNumber)
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

func (b *BayesClassifier) Classify(text string) float64 {
    words := filterWords(splitText(text))
    eta := 0.0
    for _, word := range words {
        wordStat, _ := b.Words[word]
        eta += math.Log(1.0 - wordStat.CorrectedProb()) - math.Log(wordStat.CorrectedProb())
        fmt.Println(word, " prob: ", wordStat.CorrectedProb())
    }

    return 1.0 / (1.0 + math.Exp(eta))
}
