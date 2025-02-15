package gowaku

import "encoding/json"

type jsonResponseError struct {
	Error *string `json:"error"`
}

type jsonResponseSuccess struct {
	Result interface{} `json:"result"`
}

func PrepareJSONResponse(result interface{}, err error) string {

	if err != nil {
		errStr := err.Error()
		errResponse := jsonResponseError{
			Error: &errStr,
		}
		response, _ := json.Marshal(&errResponse)
		return string(response)
	}

	data, err := json.Marshal(jsonResponseSuccess{Result: result})
	if err != nil {
		return PrepareJSONResponse(nil, err)
	}
	return string(data)
}

func MakeJSONResponse(err error) string {
	if err != nil {
		errStr := err.Error()
		outBytes, _ := json.Marshal(jsonResponseError{Error: &errStr})
		return string(outBytes)
	}

	out := jsonResponseSuccess{
		Result: true,
	}
	outBytes, _ := json.Marshal(out)

	return string(outBytes)
}
