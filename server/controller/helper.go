package controller

import "fmt"

func prepareErrorResponse(err error, m string) string {
	return fmt.Sprintf(`{"ok":false,"message":"%s","description":"%s"}`, m, err.Error())
}
