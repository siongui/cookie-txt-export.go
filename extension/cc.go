package main

import (
	. "github.com/siongui/godom"
	"github.com/siongui/gojianfan"
	"strings"
)

var excludedElement = map[string]bool{
	"style":    true,
	"script":   true,
	"noscript": true,
	"iframe":   true,
	"object":   true,
}

func traverse(elm *Object) {
	nodeType := elm.NodeType()

	if nodeType == 1 || nodeType == 9 {
		// element node or document node

		if _, in := excludedElement[strings.ToLower(elm.TagName())]; in {
			return
		}

		for _, node := range elm.ChildNodes() {
			// recursively call to traverse
			traverse(node)
		}
		return
	}

	if nodeType == 3 {
		// text node
		v := strings.TrimSpace(elm.NodeValue())
		if v != "" {
			elm.SetNodeValue(gojianfan.S2T(v))
		}
		return
	}
}

func main() {
	traverse(Document)
}
