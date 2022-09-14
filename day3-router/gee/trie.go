package gee

import "strings"

type node struct {
	pattern  string  //待匹配路由
	part     string  //路由中的一部分
	children []*node //子结点
	isWild   bool    //是否精确匹配， part含有：或*时为true
}

//结点匹配规则,遍历整棵前缀树，找到第一个匹配的节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

//遍历整棵前缀树，针对待匹配部分，找到所有匹配的节点
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child in range n.children{
		if child.part== part || child.isWild{
			nodes.append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int){
	// 节点个数等于前缀树高度，给pattern赋予最终匹配路径的值
	if len(parts) == height{
		n.pattern = pattern
		return
	}
    //获得上一次匹配成功的最后一个节点
	part := parts[height]
	//在本层中查找这个节点
	child := n.matchChild(part)
	//TODO :节点为空时，新建这个节点，插到children域中
	if child == nil{
		child = &{
			part: part
			isWild: part[0] == ":" || part[0] == "*"
		}
		n.children = append(n.children, child)
	}
	//返回本层查找成功，查找不成功时新插入了子结点
	child.insert(pattern, parts, height+1)
}


//查询功能：递归查询每一层的节点
//将退出规则定义为以下情况：
//1.匹配到了* 2.匹配失败 3.匹配到了第len(parts)层节点

func (n *node) search(parts []string, height int) *node{
	if len(parts) == height || strings.HasPreifx(n.part, "*"){
		//到了最后一层，但是匹配失败
		if n.pattern == ""{
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children{
		result := child.search(parts, height+1)
		if result != nil{
			return result
		}
	}
	return nil
}