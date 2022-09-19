package gee

import (
	"strings"
)

type node struct {
	pattern  string  //待匹配路由, 例如：/p/:lang
	part     string  //路由中的一部分，例如：lang
	children []*node //子结点
	isWild   bool    //是否精确匹配，part含有*或:时为true
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if part == child.part || child.isWild == true {
			return child
		}
	}
	return nil
}

//所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if part == child.part || child.isWild == true {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

//在树中插入节点
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		//插入新的节点
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	//当前层插入完成后直接从下一层进行查找并插入
	child.insert(pattern, parts, height+1)
}

//该函数实际上是使用part去下一层的所有孩子节点中进行查找，查找成功则返回节点
//如果已经到了最后一层，也直接返回当前叶子节点
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
