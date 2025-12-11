package todo

import (
	"reflect"
	"testing"
)

// MockStore implements the todo.Store interface for testing.
type MockStore struct {
	// Fields to control behavior and capture calls
	MockLoadData []Item
	MockLoadErr  error
	MockSaveErr  error

	// Data captured on Save call
	SavedData  []Item
	SaveCalled bool
}

// Load returns the mock data/error configured in the MockStore.
func (m *MockStore) Load() ([]Item, error) {
	return m.MockLoadData, m.MockLoadErr
}

// Save captures the data passed to it and returns the mock error.
func (m *MockStore) Save(items []Item) error {
	m.SavedData = items
	m.SaveCalled = true
	return m.MockSaveErr
}

// --- Logic Function Tests ---

func TestNewTodoList_Load(t *testing.T) {
	mockItems := []Item{
		{Id: 1, Task: "Existing Task 1", Done: false},
		{Id: 5, Task: "Existing Task 2", Done: true},
	}
	mockStore := &MockStore{
		MockLoadData: mockItems,
	}

	list, err := NewTodoList(mockStore)

	if err != nil {
		t.Fatalf("NewTodoList failed: %v", err)
	}

	// Test if items were loaded correctly
	if !reflect.DeepEqual(list.Items, mockItems) {
		t.Errorf("Loaded items mismatch. Got %v, want %v", list.Items, mockItems)
	}

	// Test if nextId is correctly calculated (should be max Id + 1 = 6)
	if list.NextId != 6 {
		t.Errorf("nextId incorrect. Got %d, want %d", list.NextId, 6)
	}
}

func TestTodoList_Add(t *testing.T) {
	// Start with an empty list
	mockStore := &MockStore{}
	list, _ := NewTodoList(mockStore)

	task := "Buy milk"
	addedItem, err := list.Add(task)

	if err != nil {
		t.Fatalf("Add failed: %v", err)
	}
	if addedItem.Task != task {
		t.Errorf("Added item task mismatch. Got %s, want %s", addedItem.Task, task)
	}
	if addedItem.Id != 1 {
		t.Errorf("Added item Id mismatch. Got %d, want %d", addedItem.Id, 1)
	}
	if !mockStore.SaveCalled {
		t.Error("Save was not called after Add")
	}
	// Check if the saved data contains the new item
	if len(mockStore.SavedData) != 1 || mockStore.SavedData[0].Task != task {
		t.Errorf("Saved data mismatch. Got %v", mockStore.SavedData)
	}
}

func TestTodoList_MarkDone(t *testing.T) {
	mockItems := []Item{
		{Id: 10, Task: "Need to be done", Done: false},
		{Id: 11, Task: "Already done", Done: true},
	}
	mockStore := &MockStore{MockLoadData: mockItems}
	list, _ := NewTodoList(mockStore)

	// 1. Mark an undone item as done
	err := list.MarkDone(10)
	if err != nil {
		t.Fatalf("MarkDone(10) failed: %v", err)
	}
	if !list.Items[0].Done {
		t.Error("Item 10 was not marked as done")
	}
	if !mockStore.SaveCalled {
		t.Error("Save was not called after MarkDone")
	}

	// 2. Try to mark a non-existent item
	err = list.MarkDone(99)
	if err == nil {
		t.Error("MarkDone(99) should have failed but dIdn't")
	}
}

func TestTodoList_Delete(t *testing.T) {
	mockItems := []Item{
		{Id: 1, Task: "A"},
		{Id: 2, Task: "B"},
		{Id: 3, Task: "C"},
	}
	mockStore := &MockStore{MockLoadData: mockItems}
	list, _ := NewTodoList(mockStore)

	// 1. Delete the mIddle item (Id 2)
	err := list.Delete(2)
	if err != nil {
		t.Fatalf("Delete(2) failed: %v", err)
	}

	// Verify the size and remaining items
	if len(list.Items) != 2 {
		t.Errorf("Expected 2 items after delete, got %d", len(list.Items))
	}
	if list.Items[0].Id != 1 || list.Items[1].Id != 3 {
		t.Errorf("Items remaining are incorrect: %v", list.Items)
	}

	// 2. Check if Save was called and saved the correct data
	if len(mockStore.SavedData) != 2 || mockStore.SavedData[1].Id != 3 {
		t.Error("Saved data after Delete is incorrect")
	}

	// 3. Try to delete a non-existent item
	err = list.Delete(99)
	if err == nil {
		t.Error("Delete(99) should have failed but dIdn't")
	}
}

func TestTodoList_Clear(t *testing.T) {
	mockItems := []Item{
		{Id: 1, Task: "A"},
		{Id: 2, Task: "B"},
	}
	mockStore := &MockStore{MockLoadData: mockItems}
	list, _ := NewTodoList(mockStore)

	err := list.Clear()
	if err != nil {
		t.Fatalf("Clear failed: %v", err)
	}

	if len(list.Items) != 0 {
		t.Errorf("Expected 0 items after clear, got %d", len(list.Items))
	}
	if list.NextId != 1 {
		t.Errorf("nextId should be 1 after clear, got %d", list.NextId)
	}
	if !mockStore.SaveCalled || len(mockStore.SavedData) != 0 {
		t.Error("Save was not called or dId not save an empty list after Clear")
	}
}
