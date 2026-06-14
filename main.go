package main

import (
	"github.com/charmbracelet/bubbles/list"
)

// Identifiable interface allows us to track items across updates.
type Identifiable interface {
	ID() string
}

// UpdateListPreservingSelection updates the list items while attempting to keep the selection on the same item.
func UpdateListPreservingSelection(m *list.Model, newItems []list.Item) {
	var selectedID string
	prevIndex := m.Index()

	// Capture ID of currently selected item if it implements Identifiable
	if selected := m.SelectedItem(); selected != nil {
		if idItem, ok := selected.(Identifiable); ok {
			selectedID = idItem.ID()
		}
	}

	m.SetItems(newItems)

	if selectedID != "" {
		// Try to find the item by ID
		for i, item := range newItems {
			if idItem, ok := item.(Identifiable); ok && idItem.ID() == selectedID {
				m.Select(i)
				return
			}
		}
	}

	// Fallback: clamp to the new list bounds
	if prevIndex >= len(newItems) {
		if len(newItems) > 0 {
			m.Select(len(newItems) - 1)
		} else {
			m.Select(0)
		}
	} else {
		m.Select(prevIndex)
	}
}