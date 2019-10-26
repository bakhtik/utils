package sql

import (
	"fmt"
	"net/http"
	"strconv"

	sq "github.com/Masterminds/squirrel"
)

const emptyString = ""

func ApplyLimit(query sq.SelectBuilder, r *http.Request, limit int) (sq.SelectBuilder, error) {
	var err error
	limitStr := r.FormValue("limit")
	if limitStr != emptyString {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit == 0 {
			return query, err
		}
	}
	

	return query.Suffix(fmt.Sprintf("FETCH NEXT %d ROWS ONLY", limit)), nil
}

func ApplyOffset(query sq.SelectBuilder, r *http.Request) (sq.SelectBuilder, error) {
	var (
		err    error
		offset int
	)
	offsetStr := r.FormValue("offset")
	if offsetStr != emptyString {
		if offset, err = strconv.Atoi(offsetStr); err != nil {
			return query, err
		}
	}

	return query.Suffix(fmt.Sprintf("OFFSET %d ROWS", offset)), nil
}
