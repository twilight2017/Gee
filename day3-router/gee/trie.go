package gee

import "strings"

type node struct {
	pattern  string  //待匹配路由
	part     string  //路由中的一部分
	children []*node //子结点
	isWild   bool    //是否精确匹配， part含有：或*时为true
}

//输入一个part，从所有子结点中进行匹配
//结点匹配规则,遍历整棵前缀树，找到第一个匹配的节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

//按照part，从所有子结点中找符合条件的节点，返回一个节点列表
//遍历整棵前缀树，针对待匹配部分，找到所有匹配的节点
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

/*对路由来说，重要的是注册和匹配操作：
1.开发服务时，注册路由，映射到具体的handler
2.访问服务时，匹配路由，查找到对应的handler*/
//height指定从第几层开始匹配，parts将当前路由所有部分存储为一个列表
func (n *node) insert(pattern string, parts []string, height int) {
	// 节点个数等于前缀树高度，给pattern赋予最终匹配路径的值
	if len(parts) == height {
		n.pattern = pattern //如果位于最后一层,pattern是完整的，直接进行返回
		return
	}
	//获得上一次匹配成功的最后一个节点
	part := parts[height]
	//在下一层中查找这个节点
	child := n.matchChild(part)
	//TODO :节点为空时，新建这个节点，插到children域中
	if child == nil {
		child = &node{part: part,
			isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child) //插入这个新建的节点
	}
	//返回本层查找成功，查找不成功时新插入了子结点
	child.insert(pattern, parts, height+1) //去下一层中确认是否需要进行插入
}

//查询功能：递归查询每一层的节点
//将退出规则定义为以下情况：
//1.匹配到了* 2.匹配失败 3.匹配到了第len(parts)层节点
//height是起始进行查找的层数

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		//到了最后一层，但是匹配失败
		if n.pattern == "" {
			return nil
		}
		return n //查找成功，返回当前节点
	}
	part := parts[height]             //获得当前层对应part
	children := n.matchChildren(part) //拿着part去所有子结点中进行查找

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
