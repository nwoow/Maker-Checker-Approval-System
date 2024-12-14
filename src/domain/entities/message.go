package entities

import (
	"errors"
	"strings" // Add this import
)

// Message represents the domain entity for a message
type Message struct {
	ID           string
	Content      string
	Recipient    string
	Status       string // Pending, Approved, Rejected
	Approvers    string // Comma-separated list of approvers
	MaxApprovers int
}

// Approve marks the message as approved if all approvers validate it
func (m *Message) Approve(approver string) error {
	if m.Status != "Pending" {
		return errors.New("message is not in pending state")
	}

	// Check if approver has already approved
	approvers := map[string]bool{}
	for _, a := range splitApprovers(m.Approvers) {
		approvers[a] = true
	}

	if approvers[approver] {
		return errors.New("approver has already approved")
	}

	m.Approvers = appendApprover(m.Approvers, approver)
	if len(splitApprovers(m.Approvers)) >= m.MaxApprovers {
		m.Status = "Approved"
	}

	return nil
}

// Helper functions to handle comma-separated approvers
func splitApprovers(approvers string) []string {
	if approvers == "" {
		return []string{}
	}
	return strings.Split(approvers, ",")
}

func appendApprover(approvers, approver string) string {
	if approvers == "" {
		return approver
	}
	return approvers + "," + approver
}
