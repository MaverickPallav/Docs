// observer.go
package main

import "fmt"

// DocumentObserver interface
type DocumentObserver interface {
	Update(document *Document)
}

// CursorPositionObserver struct
type CursorPositionObserver struct {
	user *User
}

// NewCursorPositionObserver constructor
func NewCursorPositionObserver(user *User) *CursorPositionObserver {
	return &CursorPositionObserver{user: user}
}

// Update method for CursorPositionObserver
func (o *CursorPositionObserver) Update(document *Document) {
	// Update cursor position for the user
	fmt.Printf("Cursor position updated for user %s: %d\n", o.user.GetUsername(), document.GetCursorPosition(o.user))
}
