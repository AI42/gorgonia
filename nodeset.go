package gorgonia

// internal interface for optimization purposes
type nodeset interface {
	Len() int
	ToSlice() Nodes
	Contains(*Node) bool

	add(*Node) nodeset
	remove(*Node) nodeset
}

func (ns Nodes) add(n *Node) nodeset {
	return ns.Add(n)
}

func (ns NodeSet) add(n *Node) nodeset {
	ns.Add(n)
	return ns
}

func (ns Nodes) remove(n *Node) nodeset {
	return ns.Remove(n)
}

func (ns NodeSet) remove(n *Node) nodeset {
	ns.Remove(n)
	return ns
}

func (ns NodeSet) Len() int { return len(ns) }

func (ns Nodes) ToSlice() Nodes { return ns }
