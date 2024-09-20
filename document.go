// document.go
package main

type Document struct {
	content         string
	collaborators   []*User
	versionHistory  []string
	cursorPositions map[*User]int
	observers       []DocumentObserver
}

// NewDocument constructor
func NewDocument() *Document {
	return &Document{
		content:         "",
		collaborators:   []*User{},
		versionHistory:  []string{},
		cursorPositions: make(map[*User]int),
		observers:       []DocumentObserver{},
	}
}

// EditDocument modifies the content and notifies observers
func (d *Document) EditDocument(newText string) {
	d.content = newText
	d.notifyObservers()
}

// AddCollaborator adds a user to the document
func (d *Document) AddCollaborator(user *User) {
	d.collaborators = append(d.collaborators, user)
	d.cursorPositions[user] = 0

	// Add cursor position observer for the new collaborator
	observer := NewCursorPositionObserver(user)
	d.observers = append(d.observers, observer)
}

// RemoveCollaborator removes a user from the document
func (d *Document) RemoveCollaborator(user *User) {
	for i, collaborator := range d.collaborators {
		if collaborator == user {
			d.collaborators = append(d.collaborators[:i], d.collaborators[i+1:]...)
			break
		}
	}
	delete(d.cursorPositions, user)
}

// UpdateCursorPosition updates the cursor position for a user
func (d *Document) UpdateCursorPosition(user *User, newPosition int) {
	d.cursorPositions[user] = newPosition
}

// GetCursorPosition retrieves the cursor position for a user
func (d *Document) GetCursorPosition(user *User) int {
	if position, exists := d.cursorPositions[user]; exists {
		return position
	}
	return 0
}

// NotifyObservers notifies all observers of a change
func (d *Document) notifyObservers() {
	for _, observer := range d.observers {
		observer.Update(d)
	}
}
