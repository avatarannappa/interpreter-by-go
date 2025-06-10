package evaluator

import "awesomeProject/object"

var builtins = map[string]*object.Buildin{
	"len": &object.Buildin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"first": &object.Buildin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}
			array := args[0].(*object.Array)
			if len(array.Elements) == 0 {
				return NULL
			}
			return array.Elements[0]
		},
	},
	"last": &object.Buildin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
			}
			array := args[0].(*object.Array)
			if len(array.Elements) == 0 {
				return NULL
			}
			return array.Elements[len(array.Elements)-1]
		},
	},
	"rest": &object.Buildin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
			}
			array := args[0].(*object.Array)
			if len(array.Elements) == 0 {
				return NULL
			}

			//newElements := make([]object.Object, len(array.Elements)-1, len(array.Elements)-1)
			//copy(newElements, array.Elements)
			var newElements []object.Object
			newElements = append(newElements, array.Elements[1:]...)
			return &object.Array{Elements: newElements}
		},
	},
	"push": &object.Buildin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
			}
			array := args[0].(*object.Array)

			var newElements []object.Object
			newElements = append(newElements, array.Elements...)
			newElements = append(newElements, args[1])
			return &object.Array{Elements: newElements}
		},
	},
}
