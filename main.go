package main

import "log"

func main() {
	trace := `{
 "project_id": "d7f3df5e-1729-43e1-96b4-424fed3d9050",
	"worker_id": "f31679c9-965a-46cc-b473-d3f7a718c63d",
	"tenant_id": "8ffeed0b-00f4-43cb-848a-4804ca7a0f7e",
	"type": "http-log",
	"timestamp": "2021-03-04T08:19:52.453+0000",
	"level": "info",
	"detail": {
		"operation": {
 		"request_mode": "single",
			"query_execution_time": 6.392837e-3,
			"user_vars": { "x-hasura-role":"admin" },
			"request_id": "531820d5-4b4a-43f4-b97b-6db48b5964f5",
     "parameterized_query_hash": "3516614133fe2eafbff9ec3d42e3ad1a3884c5df",
			"response_size": 41,
			"request_size": 99,
			"request_read_time": 1.7153e-5,
			"query": {
				"operationName": "MyQuery",
				"query":"query MyQuery { users { id } }"
			}
		},
		"http_info": {
			"status":200,
			"request_headers": {
				"Hasura-Client-Name": "hasura-console"
			},
			"http_version":"HTTP/1.0",
			"url": "/v1/graphql",
			"ip":"172.17.0.2",
			"method":"POST",
			"content_encoding":"gzip"
		}
	}
}`
	log.Println(trace)
}
