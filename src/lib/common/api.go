package common

import (
	"fmt"
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
		relationList[i] = strcase.ToCamel(relation)
	}
	return relationList
}

func ConditionsFromQuery(c *fiber.Ctx) SQLConditions {
	filters := c.Query("filters", "")
	if filters == "" {
		return SQLConditions{}
	}
	fmt.Printf("filters: %s\n", filters)
	filterList := strings.Split(filters, ",")

	var conditions []SQLCondition
	for _, filter := range filterList {
		filterSplit := strings.Split(filter, ":")
		condition := SQLCondition{}
		condition.Field = strcase.ToSnake(filterSplit[0])
		condition.Value = filterSplit[2]
		switch filterSplit[1] {
		case "like":
			condition.Comparator = Like
		case "eq":
			condition.Comparator = Equal
		case "ilike":
			condition.Comparator = ILike
		}
		fmt.Printf("condition: %v\n", helpers.PrettyStruct(condition))
		conditions = append(conditions, condition)
	}
	return conditions
}
