package binary

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// ShiftLeft node
type ShiftLeft struct {
	Position *position.Position
	Left     node.Node
	Right    node.Node
}

// NewShiftLeft node constructor
func NewShiftLeft(Variable node.Node, Expression node.Node) *ShiftLeft {
	return &ShiftLeft{
		Left:  Variable,
		Right: Expression,
	}
}

// SetPosition sets node position
func (n *ShiftLeft) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *ShiftLeft) GetPosition() *position.Position {
	return n.Position
}

// Attributes returns node attributes as map
func (n *ShiftLeft) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ShiftLeft) Walk(v walker.Visitor) {
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
