package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	lexer := lexer.New(input)
	parser := New(lexer)

	program := parser.ParseProgram()
	checkParserErrors(t, parser)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, tt.expectedIdentifier) {
			return
		}
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;
	`
	lexer := lexer.New(input)
	parser := New(lexer)
	program := parser.ParseProgram()
	checkParserErrors(t, parser)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	for _, statement := range program.Statements {
		returnStatement, ok := statement.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("statement not *ast.ReturnStatement. got=%T", statement)
			continue
		}
		if returnStatement.TokenLiteral() != "return" {
			t.Errorf("returnStatement.TokenLiteral not 'return', got %q",
				returnStatement.TokenLiteral())
		}
	}
}

func testLetStatement(test *testing.T, statement ast.Statement, name string) bool {
	if statement.TokenLiteral() != "let" {
		test.Errorf("s.TokenLiteral not 'let'. got=%q", statement.TokenLiteral())
		return false
	}

	letStatement, ok := statement.(*ast.LetStatement)
	if !ok {
		test.Errorf("s not *ast.LetStatement. got=%T", statement)
		return false
	}

	if letStatement.Name.Value != name {
		test.Errorf("letStatement.Name.Value not '%s'. got=%s", name, letStatement.Name.Value)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		test.Errorf("letStatement.Name.TokenLiteral() not '%s'. got=%s",
			name, letStatement.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
