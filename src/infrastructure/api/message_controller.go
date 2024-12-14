package api

import (
	"encoding/json"
	"net/http"
	"maker_checker_sqlite/src/application"
)

type MessageController struct {
	Handler *application.MessageHandler
}

func (c *MessageController) HandleSendMessage(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Content     string `json:"content"`
		Recipient   string `json:"recipient"`
		MaxApprovers int    `json:"max_approvers"`
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := c.Handler.SendMessage(req.Content, req.Recipient, req.MaxApprovers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func (c *MessageController) HandleApproveMessage(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		MessageID string `json:"message_id"`
		Approver  string `json:"approver"`
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.Handler.ApproveMessage(req.MessageID, req.Approver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
