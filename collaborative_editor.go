// collaborative_editor.go
package main

type CollaborativeEditor struct {
	document *Document
}

// NewCollaborativeEditor constructor
func NewCollaborativeEditor(document *Document) *CollaborativeEditor {
	return &CollaborativeEditor{document: document}
}

// AddCollaborator adds a user to the document
func (ce *CollaborativeEditor) AddCollaborator(user *User) {
	ce.document.AddCollaborator(user)
}

// RemoveCollaborator removes a user from the document
func (ce *CollaborativeEditor) RemoveCollaborator(user *User) {
	ce.document.RemoveCollaborator(user)
}

// EditDocument modifies the document content
func (ce *CollaborativeEditor) EditDocument(newText string) {
	ce.document.EditDocument(newText)
}
