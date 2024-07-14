package presenters

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/takumi616/ielts-vocabularies-api/adapters/dto"
	"github.com/takumi616/ielts-vocabularies-api/adapters/presenters/utils"
)

type ErrPresenter struct {
}

// Write error message
func (p *ErrPresenter) WriteErrResp(ctx context.Context, errMsg error, w http.ResponseWriter, statusCode int) {
	utils.CreateHeader(w, statusCode)

	errMsgDto := &dto.ErrMsgDto{
		Message: errMsg.Error(),
	}

	if err := json.NewEncoder(w).Encode(errMsgDto); err != nil {
		log.Printf("failed to write error message using response writer: %v", err)
		log.Printf("received error message as a parameter is: %v", errMsg)
		return
	}
}
