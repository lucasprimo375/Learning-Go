package requests

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

func MakeRequestWithAuthentication(r *http.Request, method, url string, data io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	cookie, err := cookies.Read(r)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
