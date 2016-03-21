package decisionTree

type Tree struct {
    SubTrees map[string]Tree
    Label string
    Feature string
}

func NewTree() Tree {
    tree := Tree{ SubTrees: make(map[string]Tree) }
    return tree
}

func (t *Tree) Res(values map[string]string) string {
    if t.Label != "" {
        return t.Label
    }
    subtree, ok := t.SubTrees[values[t.Feature]]
    if ok {
        return subtree.Res(values)
    }
    return ""
}
