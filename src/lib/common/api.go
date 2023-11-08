package common

import (
	"folly/src/lib/helpers"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/iancoleman/strcase"
)

type ResponseStatus string

const (
	Success ResponseStatus = "success"
	Error   ResponseStatus = "error"
)

type ApiResponse[T any] struct {
	Status  ResponseStatus `json:"status"`
	Message string         `json:"message"`
	Data    T              `json:"data,omitempty"`
	Error   string         `json:"error,omitempty"`
}

type Page[T any] struct {
	Items    []T   `json:"items"`
	Page     int   `json:"page"`
	Size     int   `json:"size"`
	Total    int64 `json:"total"`
	Filtered int64 `json:"filtered"`
}

func NewPage[T any](items []T, page int, size int, total int64, filtered int64) Page[T] {
	return Page[T]{
		Items:    items,
		Page:     page,
		Size:     size,
		Total:    total,
		Filtered: filtered,
	}
}

func NewSuccessResponse[T any](data T, message string) ApiResponse[T] {
	return ApiResponse[T]{
		Status:  Success,
		Message: message,
		Data:    data,
	}
}

func NewErrorResponse(err error, message string) ApiResponse[any] {
	return ApiResponse[any]{
		Status:  Error,
		Message: message,
		Error:   err.Error(),
	}
}

func NewValidationErrorResponse(errors []*helpers.ValidationErrors, message string) ApiResponse[any] {
	return ApiResponse[any]{
		Status:  Error,
		Message: message,
		Error:   "Validation failed",
		Data:    errors,
	}
}

func OrderBysFromQuery(c *fiber.Ctx) OrderBys {
	orders := c.Query("orders", "")
	if orders == "" {
		return OrderBys{}
	}
	orderList := strings.Split(orders, ",")

	var orderBys OrderBys
	for _, order := range orderList {
		orderSplit := strings.Split(order, ":")
		orderBy := OrderBy{}
		var field, direction string
		if len(orderSplit) == 1 {
			field = orderSplit[0]
			direction = "asc"
		} else if len(orderSplit) == 2 {
			field = orderSplit[0]
			direction = orderSplit[1]
		} else {
			continue
		}
		orderBy.Field = strcase.ToSnake(field)
		switch direction {
		case "asc":
			orderBy.Direction = Asc
		case "desc":
			orderBy.Direction = Desc
		default:
			orderBy.Direction = Asc
		}
		orderBys = append(orderBys, orderBy)
	}
	return orderBys
}

func PageableFromQuery(c *fiber.Ctx) (Pageable, error) {
	page, err := strconv.Atoi(c.Query("page", "0"))
	if err != nil {
		return Pageable{}, err
	}
	size, err := strconv.Atoi(c.Query("size", "10"))
	if err != nil {
		return Pageable{}, err
	}
	return Pageable{
		Page: page,
		Size: size,
	}, nil
}

func RelationsFromQuery(c *fiber.Ctx) []string {
	relations := c.Query("relations", "")
	if relations == "" {
		return []string{}
	}
	relationList := strings.Split(relations, ",")
	for i, relation := range relationList {
		path := strings.Split(relation, ".")
		for j, part := range path {
			path[j] = strcase.ToCamel(part)
		}
		relation = strings.Join(path, ".")

		relationList[i] = relation
	}
	return relationList
}

func ConditionsFromQuery(c *fiber.Ctx) SQLConditions {
	filters := c.Query("filters", "")
	if filters == "" {
		return SQLConditions{}
	}
	filterList := strings.Split(filters, ",")

	var conditions []SQLCondition
	for _, filter := range filterList {
		filterSplit := strings.Split(filter, ":")
		condition := SQLCondition{}
		var field, value, operator, conditionType string
		if len(filterSplit) == 3 {
			field = filterSplit[0]
			operator = filterSplit[1]
			value = filterSplit[2]
			conditionType = "or"
		} else if len(filterSplit) == 4 {
			conditionType = filterSplit[0]
			field = filterSplit[1]
			operator = filterSplit[2]
			value = filterSplit[3]
		} else {
			continue
		}
		condition.Field = strcase.ToSnake(field)
		condition.Value = value
		switch operator {
		case "like":
			condition.Comparator = Like
		case "ilike":
			condition.Comparator = ILike
		case "gt":
			condition.Comparator = GreaterThan
		case "lt":
			condition.Comparator = LessThan
		case "gte":
			condition.Comparator = GreaterEqualThan
		case "lte":
			condition.Comparator = LessEqualThan
		case "in":
			condition.Comparator = In
		default:
			condition.Comparator = Equal
		}
		condition.Type = SQLConditionType(conditionType)
		conditions = append(conditions, condition)
	}
	return conditions
}
