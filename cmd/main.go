package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// StringService 服务的定义
type StringService interface {
	hello(string) (string, error)
	count(string) int
}

// 服务的具体实现者
type stringService struct {
}

func (stringService) hello(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return "hello" + s, nil
}

func (stringService) count(string) int {
	return 100
}

//服务数据的结构,为什么需要这个??因为rpc??
type helloRequest struct {
	S string `json:"s"`
}

type helloResponse struct {
	V   string `json:"v"`
	ERR string `json:"err,omitempty"`
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

//服务数据的结构,为什么需要这个??因为rpc??

//相当于对服务的具体实现做了一层封装
//端点??? 参数为实现服务的结构体,返回一个端点,端点是方法类型, 请求转化成结构体,从结构体获取请求数据
func makeHelloEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(helloRequest)
		v, err := svc.hello(req.S)
		if err != nil {
			return helloResponse{v, err.Error()}, nil
		}
		return helloResponse{v, ""}, nil
	}
}

//相当于对服务的具体实现做了一层封装
//端点??? 参数为实现服务的结构体,返回一个端点,端点是方法类型, 请求转化成结构体,从结构体获取请求数据
func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := svc.count(req.S)
		return countResponse{v}, nil
	}
}

// ErrEmpty 空字符串
var ErrEmpty = errors.New("Empty string")

func main() {
	//服务实现者
	svc := stringService{}

	//某个服务的handel
	uppercaseHandler := httptransport.NewServer(
		makeHelloEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	//某个服务的handel
	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)
	//路由注册
	http.Handle("/hello", uppercaseHandler)
	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request helloRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
