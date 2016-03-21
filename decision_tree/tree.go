package decisionTree

type Tree struct {
    SubTrees map[string]Tree
    label string
}

func NewTree() Tree {
    tree := Tree{}
    return tree
}

func (t *Tree) Label(feature, value string) string {
    if t.label != "" {
        return t.label
    }
    return ""
}
