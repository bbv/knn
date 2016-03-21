package decisionTree

import (
    "fmt"
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
    // fmt.Println("Entropy: ", shannonEnt, "; dataSet: ", dataSet)
    return shannonEnt
}

func SplitDataSet(dataSet [][]int, axis int, value int) [][]int {
    var retDataSet [][]int
    // fmt.Println("splitting: ", dataSet)
    for _, featureVector := range dataSet {
        if featureVector[axis] == value {
            reducedFeatVector := make([]int, len(featureVector) - 1)
            copy(reducedFeatVector, featureVector[:axis])
            copy(reducedFeatVector[axis:], featureVector[axis+1:])
            retDataSet = append(retDataSet, reducedFeatVector)
        }
    }
    // fmt.Println("splitting result: ", retDataSet)
    // fmt.Println("splitting source: ", dataSet)
    return retDataSet
}

func ChooseBestFeatureToSplit(dataSet [][]int) int {
    fmt.Println(dataSet)
    numFeatures := len(dataSet[0]) - 1
    baseEntropy := CalcShannonEntropy(dataSet)
    fmt.Println("base entropy: ", baseEntropy)
    bestInfoGain := 0.0
    bestFeature := -1

    for i := 0; i < numFeatures; i++ {
        featList := make(map[int]bool)
        for row := 0; row < len(dataSet); row++ {
            featList[dataSet[row][i]] = true
        }
        newEntropy := 0.0
        for value := range featList {
            subDataSet := SplitDataSet(dataSet, i, value)
            // fmt.Println(i," ", value," ", subDataSet)
            prob := float64(len(subDataSet)) / float64(len(dataSet))
            newEntropy += prob * CalcShannonEntropy(subDataSet)
            // fmt.Println(newEntropy)
        }

        infoGain := baseEntropy - newEntropy
        fmt.Println(i, ": ", infoGain)
        if infoGain >= bestInfoGain {
            bestInfoGain = infoGain
            bestFeature = i
        }
    }
    return bestFeature
}

func MajorityCnt(classList []int) int {
    classCnt := make(map[int]int)

    for _, v := range classList {
        classCnt[v] += 1
    }
    maxV := -1
    for v := range classCnt {
        if classCnt[v] >= maxV {
            maxV = v
        }
    }
    return maxV
}

func CreateTree(dataSet [][]int, labels []string, valueLabels [][]string) Tree {
    // bestFeature := ChooseBestFeatureToSplit(dataSet)
    valuesColIndex := len(valueLabels) - 1

    tree := NewTree()
    classList := make([]int, len(dataSet))
    for i, v := range dataSet {
        classList[i] = v[len(v)-1]
    }
    if len(dataSet[0]) == 1 {
        tree.Label = valueLabels[valuesColIndex][MajorityCnt(classList)]
        return tree
    }

    bestFeature := ChooseBestFeatureToSplit(dataSet)
    bestFeatLabel := labels[bestFeature]
    uniqueFeatures := make(map[int]int)
    for _, v := range dataSet {
        uniqueFeatures[v[bestFeature]] += 1
    }
    for value := range uniqueFeatures {
        subLabels := make([]string, len(labels)-1)
        copy(subLabels, labels[:bestFeature])
        copy(subLabels, labels[bestFeature+1:])

        subValueLabels := make([][]string, len(valueLabels)-1)
        copy(subValueLabels, valueLabels[:bestFeature])
        copy(subValueLabels, valueLabels[bestFeature+1:])

        subDataSet := SplitDataSet(dataSet, bestFeature, value)
        tree.Feature = bestFeatLabel
        tree.SubTrees[valueLabels[bestFeature][value]] = CreateTree(subDataSet, subLabels, subValueLabels)
    }


    return tree
}
