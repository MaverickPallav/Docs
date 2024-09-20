// main.go
package main

import "fmt"

func main() {
	// Create users
	user1 := NewUser("user1")
	user2 := NewUser("user2")

	// Create a document
	document := NewDocument()

	// Create a collaborative editor
	collaborativeEditor := NewCollaborativeEditor(document)

	// Add collaborators to the document
	collaborativeEditor.AddCollaborator(user1)
	collaborativeEditor.AddCollaborator(user2)

	// Edit the document
	collaborativeEditor.EditDocument("This is the edited document content.")

	// Update cursor position for user1
	document.UpdateCursorPosition(user1, 10)

	// Update cursor position for user2
	document.UpdateCursorPosition(user2, 15)

	// Show final output
	fmt.Printf("All collaborators: %s, %s\n", user1.GetUsername(), user2.GetUsername())
}
