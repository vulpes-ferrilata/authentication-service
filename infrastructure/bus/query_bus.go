package bus

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
)

type QueryBus interface {
	Register(queryHandlers ...QueryHandler) error
	Use(queryHandlerWrappers ...QueryHandlerWrapper)
	Execute(ctx context.Context, query interface{}) (interface{}, error)
}

func NewQueryBus() QueryBus {
	queryBus := &queryBus{
		queryHandlerFuncs:    make(map[string]QueryHandlerFunc),
		queryHandlerWrappers: make([]QueryHandlerWrapper, 0),
	}

	return queryBus
}

type queryBus struct {
	queryHandlerFuncs    map[string]QueryHandlerFunc
	queryHandlerWrappers []QueryHandlerWrapper
}

func (q *queryBus) addQueryHandler(queryHandler QueryHandler) error {
	query := queryHandler.GetQuery()
	queryName := reflect.TypeOf(query).String()

	if _, ok := q.queryHandlerFuncs[queryName]; ok {
		return errors.Errorf("query (%s) is already assigned", queryName)
	}

	q.queryHandlerFuncs[queryName] = queryHandler.Handle

	return nil
}

func (q *queryBus) Register(queryHandlers ...QueryHandler) error {
	for _, queryHandler := range queryHandlers {
		if err := q.addQueryHandler(queryHandler); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func (q *queryBus) Use(queryHandlerWrappers ...QueryHandlerWrapper) {
	q.queryHandlerWrappers = append(q.queryHandlerWrappers, queryHandlerWrappers...)
}

func (q queryBus) Execute(ctx context.Context, query interface{}) (interface{}, error) {
	queryName := reflect.TypeOf(query).String()

	queryHandlerFunc, ok := q.queryHandlerFuncs[queryName]
	if !ok {
		return nil, errors.Errorf("handler not found for query (%s)", queryName)
	}

	for i := len(q.queryHandlerWrappers) - 1; i >= 0; i-- {
		queryHandlerFunc = q.queryHandlerWrappers[i](queryHandlerFunc)
	}

	result, err := queryHandlerFunc(ctx, query)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}
