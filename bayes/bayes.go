package bayes

import "fmt"

type WordData struct {
    prob float64
    occurrencies int
}

type BayesClassifier struct {
    words map[string]WordData
    docNumber int
}
