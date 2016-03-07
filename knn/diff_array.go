package knn

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
