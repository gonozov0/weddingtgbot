package lambda

// Request represents struct that is passed to cloud function in Yandex Cloud.
type Request struct {
	HttpMethod                      string              `json:"httpMethod"`
	Headers                         map[string]string   `json:"headers"`
	URL                             string              `json:"url"`
	Params                          map[string]string   `json:"params"`
	MultiValueParams                map[string][]string `json:"multiValueParams"`
	PathParams                      map[string]string   `json:"pathParams"`
	MultiValueHeaders               map[string][]string `json:"multiValueHeaders"`
	QueryStringParameters           map[string]string   `json:"queryStringParameters"`
	MultiValueQueryStringParameters map[string][]string `json:"multiValueQueryStringParameters"`
	RequestContext                  requestContext      `json:"requestContext"`
	Body                            string              `json:"body"`
	IsBase64Encoded                 bool                `json:"isBase64Encoded"`
}

type requestContext struct {
	Identity struct {
		SourceIP  string `json:"sourceIp"`
		UserAgent string `json:"userAgent"`
	} `json:"identity"`
	HttpMethod       string `json:"httpMethod"`
	RequestID        string `json:"requestId"`
	RequestTime      string `json:"requestTime"`
	RequestTimeEpoch int64  `json:"requestTimeEpoch"`
}

// Response represents struct that is required to be returned from cloud function in Yandex Cloud.
type Response struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
}
