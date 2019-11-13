package tree

import "errors"

func (tree *Tree) DFStep(now, from *Node, res []*Node) ([]*Node, error) {
	if now.Belong != tree {
		return nil, errors.New("[Error]: DFS now node must belong to this tree!")
	} else if from != nil && from.Belong != tree {
		return nil, errors.New("[Error]: DFS from node must belong to this tree!")
	} else {
		res = append(res, now)
		for _, nod := range now.LinkNodes {
			if from == nod {
				continue
			} else {
				var err error
				res, err = tree.DFStep(nod, now, res)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return res, nil
}

func (tree *Tree) DFS(now, from *Node) ([]*Node, error) {
	var res []*Node
	return tree.DFStep(now, from, res)
}
