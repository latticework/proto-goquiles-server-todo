package api

type SendRequest struct {
	RequestURL string
	RequestJSON string
}

type SendResponse struct {
	ResponseJSON string
}

func (client *Client) Send(request *SendRequest) (*SendResponse, error) {
	body := map[string]interface{}{
		"data": request.RequestJSON,
	}


}