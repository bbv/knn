package bayes

import (
    // "fmt"
    "regexp"
    "math"
    "encoding/json"
    "io/ioutil"
    "strings"
)

var defaultThreshold int = 2
var defaultProb float64 = 0.25
var dataPath string = "."

type WordStat struct {
    Prob float64
    Occurrencies int
}

func (ws WordStat) CorrectedProb() float64 {
    if ws.Occurrencies < defaultThreshold {
        return defaultProb
    }
    if ws.Prob == 1 {
        return 0.99
    }
    if ws.Prob == 0 {
        return 0.01
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
    if good {
        b.DocFrequency = (float64(b.DocNumber - 1) * b.DocFrequency + 1.0) / float64(b.DocNumber)
    } else {
        b.DocFrequency = (float64(b.DocNumber - 1) * b.DocFrequency) / float64(b.DocNumber)
    }

    words := filterWords(splitText(text))
    uniqueWords := make(map[string]int)

    // fmt.Printf("%#v %d\n", words, len(words))
    for _, word := range words {
        if _, ok := uniqueWords[word]; !ok {
            uniqueWords[word] += 1
        }
    }

    for word := range uniqueWords {
        // fmt.Println(word)
        wordStat, ok := b.Words[word]
        if !ok {
            // fmt.Println("Creating stats for word: ", word)
            b.Words[word] = wordStat
        }
        wordStat.Occurrencies += 1
        if good {
            wordStat.Prob = (wordStat.Prob * float64(wordStat.Occurrencies - 1) + 1.0) / float64(wordStat.Occurrencies)
        } else {
            wordStat.Prob = wordStat.Prob * float64(wordStat.Occurrencies - 1) / float64(wordStat.Occurrencies)
        }
        b.Words[word] = wordStat
    }

    return
}

func splitText(text string) []string {
    text = strings.ToLower(text)
    // s := regexp.MustCompile("[^\\p{L}\\-]+").Split(text, -1)
    s := regexp.MustCompile("[^\\p{L}\\-\\p{N}]+").Split(text, -1)
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
        // fmt.Println(wordStat.CorrectedProb())
        eta += math.Log(1.0 - wordStat.CorrectedProb()) - math.Log(wordStat.CorrectedProb())
        // fmt.Println(word, " prob: ", wordStat.CorrectedProb())
    }
    // fmt.Println(eta)

    return 1.0 / (1.0 + math.Exp(eta))
}

func (b *BayesClassifier) ToJSON() ([]byte, error) {
    return json.MarshalIndent(b, "", "    ")
}

func (b *BayesClassifier) Save() error {
    str, err := b.ToJSON()
    // fmt.Println(string(str), err)
    if err != nil {
        return err
    }

    err = ioutil.WriteFile( b.filename(), str, 0666 )
    return err
}

func (b *BayesClassifier) filename() string {
    return dataPath + "/" + b.Name + ".json"
}

func LoadClassifier( filename string ) (BayesClassifier, error) {
    var b BayesClassifier
    str, err := ioutil.ReadFile(dataPath + "/" + filename + ".json")
    if err != nil {
        return NewBayesClassifier(filename), err
    }
    err = json.Unmarshal(str, &b)
    return b, err
}

func SetDataPath(path string) {
    dataPath = path
}

func SetDefaultProb(prob float64) {
    defaultProb = prob
}

func SetDefaultThreshold(threshold int) {
    defaultThreshold = threshold
}
