package bayes

import (
    "fmt"
    "regexp"
    "encoding/json"
    "io/ioutil"
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
    if good {
        b.DocFrequency = (float64(b.DocNumber - 1) * b.DocFrequency + 1.0) / float64(b.DocNumber)
    } else {
        b.DocFrequency = (float64(b.DocNumber - 1) * b.DocFrequency) / float64(b.DocNumber)
    }

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

func (b *BayesClassifier) ToJSON() ([]byte, error) {
    return json.MarshalIndent(b, "", "    ")
}

func (b *BayesClassifier) Save( path string ) error {
    str, err := b.ToJSON()
    fmt.Println(string(str), err)
    if err != nil {
        return err
    }

    err = ioutil.WriteFile( b.filename(path), str, 0666 )
    return err
}

func (b *BayesClassifier) filename(path string) string {
    return path + "/" + b.Name + ".json"
}

func Load( filename string ) (BayesClassifier, error) {
    var b BayesClassifier
    str, err := ioutil.ReadFile(filename)
    if err != nil {
        return b, err
    }
    err = json.Unmarshal(str, &b)
    return b, err
}
