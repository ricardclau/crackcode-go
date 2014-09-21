package treesgraphs

import (
	"testing"
)

func TestDepths(t *testing.T) {
	Tree := BuildTreeAsExample()

	if d := BinaryTreeDepth(Tree, 0); -1 != d {
		t.Errorf("Wrong depth for inexistant node: %d", d)
	}

	if d := BinaryTreeDepth(Tree, 10); -1 != d {
		t.Errorf("Wrong depth for inexistant node: %d", d)
	}

	if d := BinaryTreeDepth(Tree, 1); 0 != d {
		t.Errorf("Wrong depth for node 1: %d", d)
	}

	if d := BinaryTreeDepth(Tree, 2); 1 != d {
		t.Errorf("Wrong depth for node 2: %d", d)
	}

	if d := BinaryTreeDepth(Tree, 3); 1 != d {
		t.Errorf("Wrong depth for node 3: %d", d)
	}

	if d := BinaryTreeDepth(Tree, 4); 2 != d {
		t.Errorf("Wrong depth for node 4: %d", d)
	}

	if d := BinaryTreeDepth(Tree, 5); 2 != d {
		t.Errorf("Wrong depth for node 5: %d", d)
	}

	if d := BinaryTreeDepth(Tree, 6); 2 != d {
		t.Errorf("Wrong depth for node 6: %d", d)
	}

	if d := BinaryTreeDepth(Tree, 7); 2 != d {
		t.Errorf("Wrong depth for node 7: %d", d)
	}

	if d := BinaryTreeDepth(Tree, 8); 3 != d {
		t.Errorf("Wrong depth for node 8: %d", d)
	}
}

func TestLCA(t *testing.T) {
	Tree := BuildTreeAsExample()

	if lca := TreeLCA(Tree, 7, 8); 3 != lca.Key {
		t.Errorf("Wrong LCA for nodes 7,8: %d", lca.Key)
	}

	if lca := TreeLCA(Tree, 2, 8); 1 != lca.Key {
		t.Errorf("Wrong LCA for nodes 2,8: %d", lca.Key)
	}

	if lca := TreeLCA(Tree, 3, 8); 3 != lca.Key {
		t.Errorf("Wrong LCA for nodes 3,8: %d", lca.Key)
	}

	if lca := TreeLCA(Tree, 4, 5); 2 != lca.Key {
		t.Errorf("Wrong LCA for nodes 4,5: %d", lca.Key)
	}
}

func TestExpectedDistances(t *testing.T) {
	Tree := BuildTreeAsExample()
	if d := BinaryTreeDistance(Tree, 4, 5); 2 != d {
		t.Errorf("Wrong Dist(4,5): %d", d)
	}

	if d := BinaryTreeDistance(Tree, 4, 6); 4 != d {
		t.Errorf("Wrong Dist(4,6): %d", d)
	}

	if d := BinaryTreeDistance(Tree, 3, 4); 3 != d {
		t.Errorf("Wrong Dist(4,6): %d", d)
	}

	if d := BinaryTreeDistance(Tree, 2, 4); 1 != d {
		t.Errorf("Wrong Dist(4,6): %d", d)
	}

	if d := BinaryTreeDistance(Tree, 8, 5); 5 != d {
		t.Errorf("Wrong Dist(4,6): %d", d)
	}

	if d := BinaryTreeDistance(Tree, 10, 5); -1 != d {
		t.Errorf("Wrong Dist(4,6): %d", d)
	}

	if d := BinaryTreeDistance(Tree, 5, 10); -1 != d {
		t.Errorf("Wrong Dist(4,6): %d", d)
	}
}

func BuildTreeAsExample() *BinaryTree {
	root := &BinaryTreeNode{Key: 1}

	root.Left = &BinaryTreeNode{Key: 2}
	root.Left.Left = &BinaryTreeNode{Key: 4}
	root.Left.Right = &BinaryTreeNode{Key: 5}

	root.Right = &BinaryTreeNode{Key: 3}
	root.Right.Left = &BinaryTreeNode{Key: 6}
	root.Right.Right = &BinaryTreeNode{Key: 7}

	root.Right.Left.Right = &BinaryTreeNode{Key: 8}

	return &BinaryTree{Root: root}
}
