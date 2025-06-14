package evaluator

import (
	"awesomeProject/ast"
	"awesomeProject/object"
)

func DefineMacros(program *ast.Program, env *object.Environment) {
	definitions := []int{}

	for i, statement := range program.Statements {
		if isMacroDefinition(statement) {
			addMacro(statement, env)
			definitions = append(definitions, i)
		}
	}

	for i := len(definitions) - 1; i >= 0; i-- {
		defIdx := definitions[i]
		program.Statements = append(program.Statements[:defIdx], program.Statements[defIdx+1:]...)
	}
}

func isMacroDefinition(statement ast.Statement) bool {
	letStatement, ok := statement.(*ast.LetStatement)
	if !ok {
		return false
	}
	_, ok = letStatement.Value.(*ast.MacroLiteral)
	if !ok {
		return false
	}
	return true
}

func addMacro(stmt ast.Statement, env *object.Environment) {
	letStmt, _ := stmt.(*ast.LetStatement)
	macroLiteral, _ := letStmt.Value.(*ast.MacroLiteral)

	macro := &object.Macro{
		Parameters: macroLiteral.Parameters,
		Body:       macroLiteral.Body,
		Env:        env,
	}

	env.Set(letStmt.Name.Value, macro)
}

func ExpandMacros(program ast.Node, env *object.Environment) ast.Node {
	return ast.Modify(program, func(node ast.Node) ast.Node {
		callExp, ok := node.(*ast.CallExpression)
		if !ok {
			return node
		}

		macro, ok := isMacroCall(callExp, env)
		if !ok {
			return node
		}

		args := quoteArgs(callExp)
		evalEnv := extendMacroEnv(macro, args)

		evaluated := Eval(macro.Body, evalEnv)

		q, ok := evaluated.(*object.Quote)
		if !ok {
			panic("we only support returning AST-nodes from macros")
		}
		return q.Node
	})
}

func isMacroCall(exp *ast.CallExpression, env *object.Environment) (*object.Macro, bool) {
	identifier, ok := exp.Function.(*ast.Identifier)
	if !ok {
		return nil, false
	}
	obj, ok := env.Get(identifier.Value)
	if !ok {
		return nil, false
	}
	macro, ok := obj.(*object.Macro)
	if !ok {
		return nil, false
	}
	return macro, true
}

func quoteArgs(exp *ast.CallExpression) []*object.Quote {
	var args []*object.Quote

	for _, argument := range exp.Arguments {
		args = append(args, &object.Quote{Node: argument})
	}

	return args
}

func extendMacroEnv(macro *object.Macro, args []*object.Quote) *object.Environment {
	extendedEnv := object.NewEnclosedEnvironment(macro.Env)

	for i, parameter := range macro.Parameters {
		extendedEnv.Set(parameter.Value, args[i])
	}

	return extendedEnv
}
