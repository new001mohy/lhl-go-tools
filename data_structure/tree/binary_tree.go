package tree

type Node[T comparable] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

type BinaryTree[T comparable] struct {
	root *Node[T]
}

func NewNode[T comparable](value T) *Node[T] {
	return &Node[T]{
		Value: value,
	}
}

// CalculateDepth 计算二叉树的深度
func (b *BinaryTree[T]) CalculateDepth() int {
	if b == nil {
		return 0
	}
	return b.root.calculateDepth()
}

func (b *Node[T]) calculateDepth() int {
	return max(b.Left.calculateDepth(), b.calculateDepth()) + 1
}

// IsFullBinaryTree 是否是一棵满二叉树
func (b *BinaryTree[T]) IsFullBinaryTree() bool {
	if b == nil {
		return true
	}
	return b.root.isFullBinaryTree(b.CalculateDepth(), 0)
}

// 满二叉树的判断条件是：
// 除了叶子节点外，所有节点均有两个子节点。节点数到最大值去，所有的叶子节点必须在同一层
// 　　1) 一颗树深度为h，最大层数为k，深度与最大层数相同，k=h;
// 　　2) 叶子数为2h;
// 　　3) 第k层的结点数是：2k-1;
// 　　4) 总结点数是：2k-1，且总节点数一定是奇数。
func (b *Node[T]) isFullBinaryTree(height, level int) bool {
	if b == nil {
		return true
	}
	if b.Left == nil && b.Right == nil {
		return height == level+1
	}
	if (b.Left == nil) != (b.Right == nil) {
		return false
	}
	return b.Left.isFullBinaryTree(height, level+1) && b.Right.isFullBinaryTree(height, level+1)
}

func (b *BinaryTree[T]) IsCompleteBinaryTree() bool {
	return false
}
