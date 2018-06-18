package binary

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// GreaterOrEqual node
type GreaterOrEqual struct {
	Left  node.Node
	Right node.Node
}

// NewGreaterOrEqual node constructor
func NewGreaterOrEqual(Variable node.Node, Expression node.Node) *GreaterOrEqual {
	return &GreaterOrEqual{
		Variable,
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *GreaterOrEqual) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *GreaterOrEqual) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Left != nil {
		v.EnterChildNode("Left", n)
		n.Left.Walk(v)
		v.LeaveChildNode("Left", n)
	}

	if n.Right != nil {
		v.EnterChildNode("Right", n)
		n.Right.Walk(v)
		v.LeaveChildNode("Right", n)
	}

	v.LeaveNode(n)
}
