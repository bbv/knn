package decisionTrees

import (
    // "fmt"
    "math"
)

func CalcShannonEntropy(dataSet [][]int) float64 {
    numEntries := len(dataSet)
    labelCounts := make(map[int]int)
    for i := 0; i < numEntries; i++ {
        currentLabel := dataSet[i][len(dataSet[i]) - 1]
        labelCounts[currentLabel] += 1
    }

    shannonEnt := 0.0
    for _, count := range labelCounts {
        prob := float64(count) / float64(numEntries)
        shannonEnt -= prob * math.Log2(prob)
    }
    return shannonEnt
}

func SplitDataSet(dataSet [][]int, axis int, value int) [][]int {
    var retDataSet [][]int
    for _, featureVector := range dataSet {
        if featureVector[axis] == value {
            reducedFeatVector := append(featureVector[:axis], featureVector[axis+1:]...)
            retDataSet = append(retDataSet, reducedFeatVector)
        }
    }
    return retDataSet
}
