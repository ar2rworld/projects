package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ContactForm struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func (c *Controller) Contact(w http.ResponseWriter, r *http.Request) {
	var form ContactForm
	// unmarshall the request body into the form struct
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &form); err != nil {
		fmt.Fprintf(w, "Error unmarshalling request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	methodName := "sendMessage"
	url := fmt.Sprintf("%s%s/%s", "https://api.telegram.org/bot", c.NotificatorToken, methodName)
	requestBody := fmt.Sprintf(`{"chat_id":"%s","text":"Name: %s\nEmail: %s\nMessage: %s"}`, c.NotificatorChatID, form.Name, form.Email, form.Message)
	requestData := bytes.NewBuffer([]byte(requestBody))

	request, err := http.NewRequest(http.MethodPost, url, requestData)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Fprintf(w, "Error creating request %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Fprintf(w, "Error sending request: %v", err.Error())
	}
	defer response.Body.Close()
	responseBody, _ := io.ReadAll(response.Body)
	fmt.Println(string(responseBody))
}
