package decisionTrees

type Tree struct {
    SubTrees map[string]Tree
    label string
}

func (t *Tree) Label(feature, value string) string {
    res, ok := t.labels[feature]
    if ok {
        return res
    }

}
