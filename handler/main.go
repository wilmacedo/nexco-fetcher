package handler

import (
	"context"
	"fmt"
)

func HandleRequest(ctx context.Context) (*string, error) {
	message := fmt.Sprintf("Hello World!")
	return &message, nil
}
