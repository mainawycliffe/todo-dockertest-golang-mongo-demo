package todos

import (
	"context"

	"github.com/mainawycliffe/todo-dockertest-golang-mongo-demo/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todos struct {
	client *mongo.Client
}

func (todos *Todos) AddTodo(todo model.Todo) error {
	collection := todos.client.Database("todos").Collection("todos")
	_, err := collection.InsertOne(context.Background(), todo)
	return err
}

func (todos *Todos) DeleteTodo(id string) ([]model.Todo, error) {
	panic("not implemented")
}

func (todos *Todos) GetTodos() ([]model.Todo, error) {
	collection := todos.client.Database("todos").Collection("todos")
	cursor, err := collection.Find(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	var todoList []model.Todo
	if err := cursor.All(context.Background(), &todoList); err != nil {
		return nil, err
	}
	return todoList, nil
}

func (todos *Todos) ToggleTodo(todo model.Todo) error {
	panic("not implemented")
}
