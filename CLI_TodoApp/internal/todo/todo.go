package todo

import (
	"errors"
	"fmt"
	"time"
)

type Item struct {
	Id        int
	Task      string
	Done      bool
	CreatedAt time.Time
}
type TodoList struct {
	Items  []Item
	Store  Store
	NextId int
}

type Store interface {
	Load() ([]Item, error)
	Save([]Item) error
}

func NewTodoList(s Store) (*TodoList, error) {
	list := &TodoList{
		Store: s,
	}
	err := list.LoadItems()
	if err != nil {
		return nil, fmt.Errorf("error loading items: %v", err)
	}
	return list, nil
}

func (list *TodoList) LoadItems() error {
	var err error
	list.Items, err = list.Store.Load()
	if err != nil {
		return fmt.Errorf("error loading items: %v", err)
	}
	list.NextId = 1
	for _, item := range list.Items {
		if item.Id >= list.NextId {
			list.NextId = item.Id + 1
		}
	}
	return nil

}
func (list *TodoList) SaveItem() error {
	return list.Store.Save(list.Items)
}
func (list *TodoList) Add(text string) (Item, error) {
	newItem := Item{
		Id:        list.NextId,
		Task:      text,
		Done:      false,
		CreatedAt: time.Now(),
	}
	list.Items = append(list.Items, newItem)
	list.NextId++
	return newItem, list.SaveItem()
}
func (list *TodoList) Get(id int) (*Item, int) {
	for i, item := range list.Items {
		if item.Id == id {
			return &list.Items[i], i
		}
	}
	return nil, -1
}
func (list *TodoList) MarkDone(id int) error {
	item, _ := list.Get(id)
	if item == nil {
		return errors.New("todo item not found")
	}
	if item.Done {
		return nil
	}
	item.Done = true
	return list.SaveItem()
}
func (list *TodoList) Delete(id int) error {
	_, index := list.Get(id)
	if index == -1 {
		return errors.New("item not found")
	}
	list.Items = append(list.Items[:index], list.Items[index+1:]...)
	return list.SaveItem()
}
func (list *TodoList) Clear() error {
	list.Items = make([]Item, 0)
	list.NextId = 1
	return list.SaveItem()
}
