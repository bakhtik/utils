package sql

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func SqBool(query sq.SelectBuilder, col, value string) sq.SelectBuilder {
	switch value {
	case "true":
		query = query.Where(sq.Eq{col: 1})
	case "false":
		query = query.Where(sq.Eq{col: 0})
	}
	return query
}

func SqLike(query sq.SelectBuilder, col string, values []string) sq.SelectBuilder {
	switch {
	case len(values) == 1:
		return query.Where(fmt.Sprintf("%s LIKE ?", col), values[0])
	case len(values) > 1:
		var args []interface{}
		pred := "("

		for i, v := range values {
			if i > 0 {
				pred += " OR "
			}
			pred += fmt.Sprintf("%s LIKE ?", col)
			args = append(args, v)
		}
		pred += ")"
		query = query.Where(pred, args...)
		return query
	default:
		return query
	}
}
