package commands

import (
	"fmt"
	"strings"
	"tgifer/ast"
)

func GetStrings(program *ast.Program) []string {
	ret := searchStrings(program)
	return ret
}

func searchStrings(node ast.Node) []string {
	result := []string{}
	switch n := node.(type) {
	case *ast.Program:
		for _, s := range n.Statements {
			ret := searchStrings(s)
			result = append(result, ret...)
		}
	case *ast.CommentStatement:
		// skip
	case *ast.ExpressionStatement:
		return searchStrings(n.Expression)
	case *ast.CallExpression:
		if n.Function.TokenLiteral() == "text" {
			return parseText(n)
		} else if n.Function.TokenLiteral() == "minilines" {
			return parseMiniLines(n)
		} else if n.Function.TokenLiteral() == "mini_line" {
			return parseMiniLine(n)
		} else if n.Function.TokenLiteral() == "str_block" {
			return parseStrBlock(n)
		} else if n.Function.TokenLiteral() == "str_seg" {
			return parseStrSeg(n)
		} else {
			//fmt.Println("func:", n.Function.TokenLiteral())
		}
	case *ast.IntegerLiteral:
		// skip
	case *ast.FloatingLiteral:
		// skip
	case *ast.ArrayLiteral:
		for _, a := range n.Elements {
			ret := searchStrings(a)
			result = append(result, ret...)
		}
	default:
		fmt.Printf("unsupported: %v[%T]\n", n, n)
	}
	return result
}

func parseText(text *ast.CallExpression) []string {
	s := ""
	for _, a := range text.Arguments {
		switch n := a.(type) {
		case *ast.ArrayLiteral:
			s += strings.Join(searchStrings(n), " ")
		}
	}
	return []string{s}
}

func parseMiniLines(minilines *ast.CallExpression) []string {
	result := []string{}
	for _, a := range minilines.Arguments {
		switch n := a.(type) {
		case *ast.ArrayLiteral:
			ret := searchStrings(n)
			result = append(result, ret...)
		}
	}
	return result
}

func parseMiniLine(miniline *ast.CallExpression) []string {
	result := []string{}
	for _, a := range miniline.Arguments {
		switch n := a.(type) {
		case *ast.ArrayLiteral:
			ret := searchStrings(n)
			result = append(result, ret...)
		}
	}
	return result
}
func parseStrBlock(n *ast.CallExpression) []string {
	result := []string{}
	for _, a := range n.Arguments {
		switch n := a.(type) {
		case *ast.ArrayLiteral:
			ret := searchStrings(n)
			result = append(result, ret...)
		}
	}
	return result
}
func parseStrSeg(n *ast.CallExpression) []string {
	args := n.Arguments
	str := args[len(args)-1].String()
	return []string{str}
}
