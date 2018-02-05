package gen

import (
	"context"
	"fmt"
	"github.com/vektah/graphql-go/example/todo"
	"github.com/vektah/graphql-go/introspection"
	"github.com/vektah/graphql-go/jsonw"
	"github.com/vektah/graphql-go/query"
	"github.com/vektah/graphql-go/schema"
	"strconv"
)

type Resolvers interface {
	Mutation_createTodo(ctx context.Context, text string) (todo.Todo, error)
	Mutation_updateTodo(ctx context.Context, id int, done bool) (todo.Todo, error)
	Query_todo(ctx context.Context, id int) (*todo.Todo, error)
	Query_lastTodo(ctx context.Context) (*todo.Todo, error)
	Query_todos(ctx context.Context) ([]todo.Todo, error)
}

var (
	mutationSatisfies     = []string{"Mutation"}
	querySatisfies        = []string{"Query"}
	todoSatisfies         = []string{"Todo"}
	__DirectiveSatisfies  = []string{"__Directive"}
	__EnumValueSatisfies  = []string{"__EnumValue"}
	__FieldSatisfies      = []string{"__Field"}
	__InputValueSatisfies = []string{"__InputValue"}
	__SchemaSatisfies     = []string{"__Schema"}
	__TypeSatisfies       = []string{"__Type"}
)

func _mutation(ec *executionContext, sel []query.Selection, it *interface{}) jsonw.Encodable {
	groupedFieldSet := ec.collectFields(sel, mutationSatisfies, map[string]bool{})
	resultMap := jsonw.Map{}
	for _, field := range groupedFieldSet {
		switch field.Name {
		case "createTodo":
			res, err := ec.resolvers.Mutation_createTodo(
				ec.ctx,
				field.Args["text"].(string),
			)
			if err != nil {
				ec.Error(err)
				continue
			}
			json := _todo(ec, field.Selections, &res)
			resultMap.Set(field.Alias, json)
			continue

		case "updateTodo":
			res, err := ec.resolvers.Mutation_updateTodo(
				ec.ctx,
				field.Args["id"].(int),
				field.Args["done"].(bool),
			)
			if err != nil {
				ec.Error(err)
				continue
			}
			json := _todo(ec, field.Selections, &res)
			resultMap.Set(field.Alias, json)
			continue

		}
		panic("unknown field " + strconv.Quote(field.Name))
	}
	return resultMap
}

func _query(ec *executionContext, sel []query.Selection, it *interface{}) jsonw.Encodable {
	groupedFieldSet := ec.collectFields(sel, querySatisfies, map[string]bool{})
	resultMap := jsonw.Map{}
	for _, field := range groupedFieldSet {
		switch field.Name {
		case "todo":
			res, err := ec.resolvers.Query_todo(
				ec.ctx,
				field.Args["id"].(int),
			)
			if err != nil {
				ec.Error(err)
				continue
			}
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := _todo(ec, field.Selections, res)
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		case "lastTodo":
			res, err := ec.resolvers.Query_lastTodo(
				ec.ctx,
			)
			if err != nil {
				ec.Error(err)
				continue
			}
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := _todo(ec, field.Selections, res)
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		case "todos":
			res, err := ec.resolvers.Query_todos(
				ec.ctx,
			)
			if err != nil {
				ec.Error(err)
				continue
			}
			json := jsonw.Array{}
			for _, val := range res {
				json1 := _todo(ec, field.Selections, &val)
				json = append(json, json1)
			}
			resultMap.Set(field.Alias, json)
			continue

		case "__schema":
			res := ec.introspectSchema()
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := ___Schema(ec, field.Selections, res)
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		case "__type":
			res := ec.introspectType(
				field.Args["name"].(string),
			)
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := ___Type(ec, field.Selections, res)
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		}
		panic("unknown field " + strconv.Quote(field.Name))
	}
	return resultMap
}

func _todo(ec *executionContext, sel []query.Selection, it *todo.Todo) jsonw.Encodable {
	groupedFieldSet := ec.collectFields(sel, todoSatisfies, map[string]bool{})
	resultMap := jsonw.Map{}
	for _, field := range groupedFieldSet {
		switch field.Name {
		case "id":
			res := it.ID
			json := jsonw.Int(res)
			resultMap.Set(field.Alias, json)
			continue

		case "text":
			res := it.Text
			json := jsonw.String(res)
			resultMap.Set(field.Alias, json)
			continue

		case "done":
			res := it.Done
			json := jsonw.Bool(res)
			resultMap.Set(field.Alias, json)
			continue

		}
		panic("unknown field " + strconv.Quote(field.Name))
	}
	return resultMap
}

func ___Directive(ec *executionContext, sel []query.Selection, it *introspection.Directive) jsonw.Encodable {
	groupedFieldSet := ec.collectFields(sel, __DirectiveSatisfies, map[string]bool{})
	resultMap := jsonw.Map{}
	for _, field := range groupedFieldSet {
		switch field.Name {
		case "name":
			res := it.Name()
			json := jsonw.String(res)
			resultMap.Set(field.Alias, json)
			continue

		case "description":
			res := it.Description()
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := jsonw.String(*res)
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		case "locations":
			res := it.Locations()
			json := jsonw.Array{}
			for _, val := range res {
				json1 := jsonw.String(val)
				json = append(json, json1)
			}
			resultMap.Set(field.Alias, json)
			continue

		case "args":
			res := it.Args()
			json := jsonw.Array{}
			for _, val := range res {
				var json1 jsonw.Encodable = jsonw.Null
				if val != nil {
					json11 := ___InputValue(ec, field.Selections, val)
					json1 = json11
				}
				json = append(json, json1)
			}
			resultMap.Set(field.Alias, json)
			continue

		}
		panic("unknown field " + strconv.Quote(field.Name))
	}
	return resultMap
}

func ___EnumValue(ec *executionContext, sel []query.Selection, it *introspection.EnumValue) jsonw.Encodable {
	groupedFieldSet := ec.collectFields(sel, __EnumValueSatisfies, map[string]bool{})
	resultMap := jsonw.Map{}
	for _, field := range groupedFieldSet {
		switch field.Name {
		case "name":
			res := it.Name()
			json := jsonw.String(res)
			resultMap.Set(field.Alias, json)
			continue

		case "description":
			res := it.Description()
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := jsonw.String(*res)
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		case "isDeprecated":
			res := it.IsDeprecated()
			json := jsonw.Bool(res)
			resultMap.Set(field.Alias, json)
			continue

		case "deprecationReason":
			res := it.DeprecationReason()
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := jsonw.String(*res)
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		}
		panic("unknown field " + strconv.Quote(field.Name))
	}
	return resultMap
}

func ___Field(ec *executionContext, sel []query.Selection, it *introspection.Field) jsonw.Encodable {
	groupedFieldSet := ec.collectFields(sel, __FieldSatisfies, map[string]bool{})
	resultMap := jsonw.Map{}
	for _, field := range groupedFieldSet {
		switch field.Name {
		case "name":
			res := it.Name()
			json := jsonw.String(res)
			resultMap.Set(field.Alias, json)
			continue

		case "description":
			res := it.Description()
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := jsonw.String(*res)
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		case "args":
			res := it.Args()
			json := jsonw.Array{}
			for _, val := range res {
				var json1 jsonw.Encodable = jsonw.Null
				if val != nil {
					json11 := ___InputValue(ec, field.Selections, val)
					json1 = json11
				}
				json = append(json, json1)
			}
			resultMap.Set(field.Alias, json)
			continue

		case "type":
			res := it.Type()
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := ___Type(ec, field.Selections, res)
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		case "isDeprecated":
			res := it.IsDeprecated()
			json := jsonw.Bool(res)
			resultMap.Set(field.Alias, json)
			continue

		case "deprecationReason":
			res := it.DeprecationReason()
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := jsonw.String(*res)
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		}
		panic("unknown field " + strconv.Quote(field.Name))
	}
	return resultMap
}

func ___InputValue(ec *executionContext, sel []query.Selection, it *introspection.InputValue) jsonw.Encodable {
	groupedFieldSet := ec.collectFields(sel, __InputValueSatisfies, map[string]bool{})
	resultMap := jsonw.Map{}
	for _, field := range groupedFieldSet {
		switch field.Name {
		case "name":
			res := it.Name()
			json := jsonw.String(res)
			resultMap.Set(field.Alias, json)
			continue

		case "description":
			res := it.Description()
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := jsonw.String(*res)
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		case "type":
			res := it.Type()
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := ___Type(ec, field.Selections, res)
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		case "defaultValue":
			res := it.DefaultValue()
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := jsonw.String(*res)
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		}
		panic("unknown field " + strconv.Quote(field.Name))
	}
	return resultMap
}

func ___Schema(ec *executionContext, sel []query.Selection, it *introspection.Schema) jsonw.Encodable {
	groupedFieldSet := ec.collectFields(sel, __SchemaSatisfies, map[string]bool{})
	resultMap := jsonw.Map{}
	for _, field := range groupedFieldSet {
		switch field.Name {
		case "types":
			res := it.Types()
			json := jsonw.Array{}
			for _, val := range res {
				var json1 jsonw.Encodable = jsonw.Null
				if val != nil {
					json11 := ___Type(ec, field.Selections, val)
					json1 = json11
				}
				json = append(json, json1)
			}
			resultMap.Set(field.Alias, json)
			continue

		case "queryType":
			res := it.QueryType()
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := ___Type(ec, field.Selections, res)
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		case "mutationType":
			res := it.MutationType()
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := ___Type(ec, field.Selections, res)
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		case "subscriptionType":
			res := it.SubscriptionType()
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := ___Type(ec, field.Selections, res)
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		case "directives":
			res := it.Directives()
			json := jsonw.Array{}
			for _, val := range res {
				var json1 jsonw.Encodable = jsonw.Null
				if val != nil {
					json11 := ___Directive(ec, field.Selections, val)
					json1 = json11
				}
				json = append(json, json1)
			}
			resultMap.Set(field.Alias, json)
			continue

		}
		panic("unknown field " + strconv.Quote(field.Name))
	}
	return resultMap
}

func ___Type(ec *executionContext, sel []query.Selection, it *introspection.Type) jsonw.Encodable {
	groupedFieldSet := ec.collectFields(sel, __TypeSatisfies, map[string]bool{})
	resultMap := jsonw.Map{}
	for _, field := range groupedFieldSet {
		switch field.Name {
		case "kind":
			res := it.Kind()
			json := jsonw.String(res)
			resultMap.Set(field.Alias, json)
			continue

		case "name":
			res := it.Name()
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := jsonw.String(*res)
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		case "description":
			res := it.Description()
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := jsonw.String(*res)
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		case "fields":
			res := it.Fields(
				field.Args["includeDeprecated"].(bool),
			)
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := jsonw.Array{}
				for _, val := range *res {
					var json11 jsonw.Encodable = jsonw.Null
					if val != nil {
						json111 := ___Field(ec, field.Selections, val)
						json11 = json111
					}
					json1 = append(json1, json11)
				}
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		case "interfaces":
			res := it.Interfaces()
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := jsonw.Array{}
				for _, val := range *res {
					var json11 jsonw.Encodable = jsonw.Null
					if val != nil {
						json111 := ___Type(ec, field.Selections, val)
						json11 = json111
					}
					json1 = append(json1, json11)
				}
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		case "possibleTypes":
			res := it.PossibleTypes()
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := jsonw.Array{}
				for _, val := range *res {
					var json11 jsonw.Encodable = jsonw.Null
					if val != nil {
						json111 := ___Type(ec, field.Selections, val)
						json11 = json111
					}
					json1 = append(json1, json11)
				}
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		case "enumValues":
			res := it.EnumValues(
				field.Args["includeDeprecated"].(bool),
			)
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := jsonw.Array{}
				for _, val := range *res {
					var json11 jsonw.Encodable = jsonw.Null
					if val != nil {
						json111 := ___EnumValue(ec, field.Selections, val)
						json11 = json111
					}
					json1 = append(json1, json11)
				}
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		case "inputFields":
			res := it.InputFields()
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := jsonw.Array{}
				for _, val := range *res {
					var json11 jsonw.Encodable = jsonw.Null
					if val != nil {
						json111 := ___InputValue(ec, field.Selections, val)
						json11 = json111
					}
					json1 = append(json1, json11)
				}
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		case "ofType":
			res := it.OfType()
			var json jsonw.Encodable = jsonw.Null
			if res != nil {
				json1 := ___Type(ec, field.Selections, res)
				json = json1
			}
			resultMap.Set(field.Alias, json)
			continue

		}
		panic("unknown field " + strconv.Quote(field.Name))
	}
	return resultMap
}

var parsedSchema = schema.MustParse("\nschema {\n\tquery: Query\n\tmutation: Mutation\n}\n\ntype Query {\n\ttodo(id: Int!): Todo\n\tlastTodo: Todo\n\ttodos: [Todo!]!\n}\n\ntype Mutation {\n\tcreateTodo(text: String!): Todo!\n\tupdateTodo(id: Int!, done: Boolean!): Todo!\n}\n\ntype Todo {\n\tid: Int!\n\ttext: String!\n\tdone: Boolean!\n}\n")
var _ = fmt.Print
