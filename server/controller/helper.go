package controller

import "fmt"

func prepareErrorResponse(err error, m string) string {
	if err == nil {
		return fmt.Sprintf(`{"ok":false,"message":"%s"}`, m)
	}
	return fmt.Sprintf(`{"ok":false,"message":"%s","description":"%s"}`, m, err.Error())
}
