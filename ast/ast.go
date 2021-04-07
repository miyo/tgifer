package ast

import (
	"bytes"
	"tgifer/token"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement `json:"statement"`
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type CommentStatement struct {
	Token token.Token `token:"comment"`
}

func (s *CommentStatement) statementNode() {}

func (s *CommentStatement) TokenLiteral() string { return s.Token.Literal }

func (s *CommentStatement) String() string {
	var out bytes.Buffer
	out.WriteString(s.TokenLiteral() + " ")
	out.WriteString("\n")
	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression `json:"expression"`
}

func (s *ExpressionStatement) statementNode() {}

func (s *ExpressionStatement) TokenLiteral() string { return s.Token.Literal }

func (s *ExpressionStatement) String() string {
	if s.Expression != nil {
		return s.Expression.String()
	}
	return ""
}

type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Statement `json:"statement"`
}

func (s *BlockStatement) statementNode() {}

func (s *BlockStatement) TokenLiteral() string { return s.Token.Literal }

func (s *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range s.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string      `json:"ident"`
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string { return i.Value }

type IntegerLiteral struct {
	Token token.Token // the token.IDENT token
	Value int64       `json:"value"`
}

func (i *IntegerLiteral) expressionNode() {}

func (i *IntegerLiteral) TokenLiteral() string { return i.Token.Literal }

func (i *IntegerLiteral) String() string { return i.Token.Literal }

type FloatingLiteral struct {
	Token token.Token // the token.IDENT token
	Value float64     `json:"value"`
}

func (i *FloatingLiteral) expressionNode() {}

func (i *FloatingLiteral) TokenLiteral() string { return i.Token.Literal }

func (i *FloatingLiteral) String() string { return i.Token.Literal }

type PrefixExpression struct {
	Token    token.Token // the prefix token, eg. !
	Operator string
	Right    Expression
}

func (e *PrefixExpression) expressionNode() {}

func (e *PrefixExpression) TokenLiteral() string { return e.Token.Literal }

func (e *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(e.Operator)
	out.WriteString(e.Right.String())
	out.WriteString(")")
	return out.String()
}

type InfixExpression struct {
	Token    token.Token // the prefix token, eg. !
	Left     Expression
	Operator string
	Right    Expression
}

func (e *InfixExpression) expressionNode() {}

func (e *InfixExpression) TokenLiteral() string { return e.Token.Literal }

func (e *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(e.Left.String())
	out.WriteString(" " + e.Operator + " ")
	out.WriteString(e.Right.String())
	out.WriteString(")")
	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

func (b *Boolean) TokenLiteral() string { return b.Token.Literal }

func (b *Boolean) String() string { return b.Token.Literal }

type FunctionLiteral struct {
	Token      token.Token // The 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (f *FunctionLiteral) expressionNode() {}

func (f *FunctionLiteral) TokenLiteral() string { return f.Token.Literal }

func (f *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(f.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(f.Body.String())
	return out.String()
}

type CallExpression struct {
	Token     token.Token  // The '(' token
	Function  Expression   `json:"function"`
	Arguments []Expression `json:"arguments"`
}

func (e *CallExpression) expressionNode() {}

func (e *CallExpression) TokenLiteral() string { return e.Token.Literal }

func (e *CallExpression) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range e.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(e.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}

type StringLiteral struct {
	Token token.Token
	Value string `json:"value"`
}

func (s *StringLiteral) expressionNode() {}

func (s *StringLiteral) TokenLiteral() string { return s.Token.Literal }

func (s *StringLiteral) String() string { return s.Token.Literal }

type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression `json:"expressions"`
}

func (a *ArrayLiteral) expressionNode() {}

func (a *ArrayLiteral) TokenLiteral() string { return a.Token.Literal }

func (a *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range a.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func (e *IndexExpression) expressionNode() {}

func (e *IndexExpression) TokenLiteral() string { return e.Token.Literal }

func (e *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(e.Left.String())
	out.WriteString("[")
	out.WriteString(e.Index.String())
	out.WriteString("]")
	out.WriteString(")")

	return out.String()
}
