package ast

import (
	"bytes"
	"monkey/token"
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
	Statements []Statement
}

func (program *Program) TokenLiteral() string {
	if len(program.Statements) > 0 {
		return program.Statements[0].TokenLiteral()
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

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (statement *LetStatement) statementNode()       {}
func (statement *LetStatement) TokenLiteral() string { return statement.Token.Literal }

func (statement *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(statement.TokenLiteral() + " ")
	out.WriteString(statement.Name.String())
	out.WriteString(" = ")
	if statement.Value != nil {
		out.WriteString(statement.Value.String())
	}
	out.WriteString(";")

	return out.String()
}

type Identifier struct {
	Token token.Token // the token.IDENTIFIER token
	Value string
}

func (identifier *Identifier) String() string { return identifier.Value }

func (identifier *Identifier) expressionNode()      {}
func (identifier *Identifier) TokenLiteral() string { return identifier.Token.Literal }

type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (statement *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(statement.TokenLiteral() + " ")
	if statement.ReturnValue != nil {
		out.WriteString(statement.ReturnValue.String())
	}
	out.WriteString(";")

	return out.String()
}

func (statement *ReturnStatement) statementNode()       {}
func (statement *ReturnStatement) TokenLiteral() string { return statement.Token.Literal }

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (statement *ExpressionStatement) String() string {
	if statement.Expression != nil {
		return statement.Expression.String()
	}
	return ""
}

func (statement *ExpressionStatement) statementNode()       {}
func (statement *ExpressionStatement) TokenLiteral() string { return statement.Token.Literal }
