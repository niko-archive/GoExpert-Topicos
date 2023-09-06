package filters

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type ParamValue struct {
	Param string
	Value string
}

type ParamDefinition struct {
	Param string
	Op    string
	Value string
}

type Pagination struct {
	Page  int
	Limit int
}

func MapFilterToSQL(filter string) (string, error) {
	switch filter {
	case "eq":
		return "=", nil
	case "gt":
		return ">", nil
	case "lt":
		return "<", nil
	case "like":
		return "like", nil
	case "ilike":
		return "ilike", nil
	case "in":
		return "in", nil
	case "neq":
		return "!=", nil
	case "not":
		return "not", nil
	case "not-in":
		return "not in", nil
	default:
		return "", errors.New("invalid filter")
	}
}

func ParseParamRequest(req *http.Request, param string) *ParamDefinition {
	queryValue := req.URL.Query().Get(param)

	if queryValue == "" {
		return nil
	}

	split := strings.Split(queryValue, ".")
	if len(split) != 2 {
		return nil
	}

	op, _ := MapFilterToSQL(split[0])
	value := split[1]

	definition := &ParamDefinition{
		Param: param,
		Op:    op,
		Value: value,
	}

	return definition
}

func (d *ParamDefinition) PreBuildQuery() (string, string) {
	return d.Param + " " + d.Op + " " + "?", d.Value
}

func ParsePaginationRequest(req *http.Request) *Pagination {
	page := req.URL.Query().Get("page")
	limit := req.URL.Query().Get("limit")

	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "10"
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		customErr := errors.New("page must be a number")
		panic(customErr)
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		customErr := errors.New("limit must be a number")
		panic(customErr)
	}

	return &Pagination{
		Page:  pageInt,
		Limit: limitInt,
	}
}

func (p *Pagination) Offset() int {
	return (p.Page - 1) * p.Limit
}
