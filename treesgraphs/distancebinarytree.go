package treesgraphs

import (
	"math"
)

type BinaryTreeNode struct {
	Key   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

// BinaryTreeDistance Gets the distance between 2 keys
// If you don´t have a link to the parent the only way is to find the Lowest Common Ancestor
// and the distance is depth(root, key1) + depth(root, key2) - 2 * depth(root, LCA)
func BinaryTreeDistance(tree *BinaryTree, key1 int, key2 int) int {
	depth1 := BinaryTreeKeyDepth(tree, key1)
	if depth1 == -1 {
		return -1
	}

	depth2 := BinaryTreeKeyDepth(tree, key2)
	if depth2 == -1 {
		return -1
	}

	lca := TreeLCA(tree, key1, key2)

	return depth1 + depth2 - 2*BinaryTreeKeyDepth(tree, lca.Key)
}

func IsTreeBalanced(tree *BinaryTree) bool {
	return math.Abs(float64(BinaryTreeRecursiveMaxDepth(tree.Root)-BinaryTreeRecursiveMinDepth(tree.Root))) <= 1
}

func BinaryTreeRecursiveMaxDepth(Root *BinaryTreeNode) int64 {
	if Root == nil {
		return 0
	}

	return 1 + int64(math.Max(float64(BinaryTreeRecursiveMaxDepth(Root.Left)), float64(BinaryTreeRecursiveMaxDepth(Root.Right))))
}

func BinaryTreeRecursiveMinDepth(Root *BinaryTreeNode) int64 {
	if Root == nil {
		return 0
	}

	return 1 + int64(math.Min(float64(BinaryTreeRecursiveMinDepth(Root.Left)), float64(BinaryTreeRecursiveMinDepth(Root.Right))))
}

func TreeLCA(Tree *BinaryTree, key1 int, key2 int) *BinaryTreeNode {
	return recursiveLCA(Tree.Root, key1, key2)
}

func recursiveLCA(root *BinaryTreeNode, key1 int, key2 int) *BinaryTreeNode {
	// Found an end of the tree
	if root == nil {
		return nil
	}

	// Found the node itself
	if root.Key == key1 || root.Key == key2 {
		return root
	}

	nodeOnLeft := recursiveLCA(root.Left, key1, key2)
	nodeOnRight := recursiveLCA(root.Right, key1, key2)

	// Think about the drawing and you will see it
	if nodeOnLeft != nil && nodeOnRight != nil {
		return root
	}

	if nodeOnLeft != nil {
		return recursiveLCA(root.Left, key1, key2)
	}

	return recursiveLCA(root.Right, key1, key2)
}

func BinaryTreeKeyDepth(Tree *BinaryTree, key int) int {
	if key == Tree.Root.Key {
		return 0
	}

	return recursiveBinaryTreeKeyDepth(Tree.Root, key, 0)
}

func recursiveBinaryTreeKeyDepth(Root *BinaryTreeNode, key int, level int) int {
	if Root == nil {
		return -1
	}

	if Root.Key == key {
		return level
	}

	depth := recursiveBinaryTreeKeyDepth(Root.Left, key, level+1)
	if depth != -1 {
		return depth
	}

	return recursiveBinaryTreeKeyDepth(Root.Right, key, level+1)
}
