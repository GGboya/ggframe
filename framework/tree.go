package framework

import (
	"errors"
	"strings"
)

type Tree struct {
	root *node // 根节点
}

type node struct {
	isLast   bool                // 是否是路由规则的最后节点
	segment  string              // 这个段所属的字符串
	handlers []ControllerHandler // 中间件 + 控制器
	childs   []*node             // 孩子节点
	parent   *node               // 父节点，双向指针
}

func newNode() *node {
	return &node{
		isLast:  false,
		segment: "",
		childs:  []*node{},
	}
}

func NewTree() *Tree {
	root := newNode()
	return &Tree{root: root}
}

// 判断一个segment是否是通用segment，即以:开头
func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

// 过滤下一层满足segment规则的子节点
func (n *node) filterChildNodes(segment string) []*node {
	if len(n.childs) == 0 {
		return nil
	}

	// 如果segment是通配符，则所有下一层节点都满足需求
	if isWildSegment(segment) {
		return n.childs
	}

	nodes := make([]*node, 0, len(n.childs))
	// 过滤下一层节点
	for _, cnode := range n.childs {
		if isWildSegment(cnode.segment) {
			// 如果下一层节点有通配符，则满足需求
			nodes = append(nodes, cnode)
		} else if cnode.segment == segment {
			// 如果下一层节点没有通配符，但是文本匹配完全匹配，则满足需求
			nodes = append(nodes, cnode)
		}
	}
	return nodes
}

// 匹配路由
func (n *node) matchNode(uri string) *node {
	// 拿到一串url，先分割成两个部分。
	segments := strings.SplitN(uri, "/", 2)
	segment := segments[0]
	// 现在去过滤一下有多少节点满足segment
	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}
	cnodes := n.filterChildNodes(segment) // 匹配segment的孩子节点

	if cnodes == nil || len(cnodes) == 0 {
		// 无法匹配
		return nil
	}

	// 如果当前的segment已经是最后一项，并且匹配到了
	if len(segments) == 1 {
		for _, cnode := range cnodes {
			if cnode.isLast {
				return cnode
			}
		}
		// 虽然存在这个segment，但并不是最后一个
		return nil
	}

	// 如果路由还存在下一级，则递归匹配
	for _, cnode := range cnodes {
		match := cnode.matchNode(segments[1])
		if match != nil {
			return match
		}
	}
	// 不匹配
	return nil
}

// AddRouter 增加路由节点，路由节点有先后顺序
/*
/book/list
/book/:id (冲突)
/book/:id/name
/book/:student/age
/:user/name(冲突)
/:user/name/:age
*/
func (tree *Tree) AddRouter(uri string, handlers []ControllerHandler) error {
	n := tree.root
	// 判断路由是否冲突
	if n.matchNode(uri) != nil {
		return errors.New("route exist:" + uri)
	}

	segments := strings.Split(uri, "/")

	for index, segment := range segments {
		// 为了统一，把非通配符都变成大写的
		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}
		isLast := index == len(segments)-1
		var objNode *node
		childNodes := n.filterChildNodes(segment)

		// 判断在当前层是否有匹配的节点
		if len(childNodes) > 0 {
			for _, cnode := range childNodes {
				if cnode.segment == segment {
					objNode = cnode
					break
				}
			}
		}

		if objNode == nil {
			// 如果没有匹配的节点，就需要自己创建一个
			cnode := newNode()
			cnode.segment = segment
			if isLast {
				cnode.isLast = true
				cnode.handlers = handlers
			}
			// 父节点指针修改
			cnode.parent = n
			n.childs = append(n.childs, cnode)
			objNode = cnode
		}

		n = objNode

	}
	return nil

}

// FindHandler 匹配uri
func (tree *Tree) FindHandler(uri string) []ControllerHandler {
	matchNode := tree.root.matchNode(uri)
	if matchNode == nil {
		return nil
	}
	return matchNode.handlers
}

// 将uri解析为parmas
func (n *node) parseParamsFromEndNode(uri string) map[string]string {
	ret := map[string]string{}
	segments := strings.Split(uri, "/")
	cnt := len(segments)
	cur := n
	for i := cnt - 1; i >= 0; i-- {
		if cur.segment == "" {
			break
		}
		// 如果是通配符节点
		if isWildSegment(cur.segment) {
			// 设置params
			ret[cur.segment[1:]] = segments[i]
		}
		cur = cur.parent
	}
	return ret
}
