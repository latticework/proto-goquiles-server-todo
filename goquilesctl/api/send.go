package api

import "github.com/latticework/proto-goquiles-server-todo/go-jali/jalicore"

type SendRequest struct {
	RequestURL  string
	RequestJSON string
}

type SendResponse struct {
	ResponseJSON string
}

func (client *Client) Send(request *SendRequest) (*SendResponse, error) {
	//	body := map[string]interface{}{
	//		"data": request.RequestJSON,
	//	}

	return _, jalicore.NotImplementedError{}.Init(_, _)
}
