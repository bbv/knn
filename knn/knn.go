package knn

import (
    "fmt"
    "errors"
    "math"
    "sort"
)

type DiffArray struct {
    Values []float64
    Indices []int
}

func NewDiffArray(len int) DiffArray {
    da := DiffArray{
        make([]float64, len),
        make([]int, len),
    }
    for i := 0; i < len; i++ {
        da.Indices[i] = i;
    }
    return da
}

func (da DiffArray) Len() int {
    return len(da.Values)
}

func (da DiffArray) Less(i, j int) bool {
    return da.Values[i] < da.Values[j]
}

func (da DiffArray) Swap(i, j int) {
    da.Values[i], da.Values[j] = da.Values[j], da.Values[i]
    da.Indices[i], da.Indices[j] = da.Indices[j], da.Indices[i]
}

func Classify( inX []float64, dataSet [][]float64, labels []string, k int ) (string, error) {
    if len(dataSet) == 0 {
        return "", errors.New("Empty data set")
    }
    diff := calcDiff(inX, dataSet)
    sort.Sort(diff)
    // fmt.Println(diff)

    return vote(labels, diff, k), nil
}

func calcDiff(inX []float64, dataSet [][]float64) DiffArray{
    res := NewDiffArray(len(dataSet))
    // fmt.Println(res)
    for i, row := range dataSet {
        // fmt.Printf("%#v %d", row, i)
        for j := 0; j < len(inX); j++ {
            res.Values[i] += math.Pow(inX[j]*inX[j] - row[j]*row[j], 2)
        }
        res.Values[i] = math.Sqrt(res.Values[i])
    }
    return res
}

func vote(labels []string, diff DiffArray, k int) string {
    m := make(map[string]int)
    for i := 0; i < int(math.Min( float64(len(diff.Values)), float64(k)) ); i++ {
        m[labels[diff.Indices[i]]] += 1
    }

    var max int
    var res string
    for k, v := range m {
        if v > max {
            max = v
            res = k
        }
    }

    return res
}

func NormalizeData(dataSet [][]float64) ([][]float64, []float64, []float64) {
    fmt.Println(dataSet)
    normalizedArray := make([][]float64, len(dataSet))
    rows := len(dataSet)
    cols := len(dataSet[0])

    mins  := make([]float64, cols)
    maxes := make([]float64, cols)

    for i := 0; i < cols; i++ {
        mins[i] = dataSet[0][i]
        maxes[i] = dataSet[0][i]
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if mins[j] > dataSet[i][j] {
                mins[j] = dataSet[i][j]
            }
            if maxes[j] < dataSet[i][j] {
                maxes[j] = dataSet[i][j]
            }
        }
    }

    fmt.Println(mins)
    fmt.Println(maxes)

    return normalizedArray, mins, maxes;
}
