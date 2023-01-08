package main

import (
	"encoding/json"
	"fmt"
	"github.com/tiketdatarisal/celeste"
)

const (
	q1 = `
{
    "source": "people",
    "find": {"$and": [{"name": "\"james\""}, {"age": {"$gt": 20}}]}
}`
	q2 = `
{
    "source": "people",
    "find": {"$and": [{"name": {"$not_like": "\"%james%\""}}]}
}`
	q3 = `
{
    "source": "people",
    "find": {"$and": [{"name": {"$in": ["\"james\"", "\"bond\""]}}]}
}`
)

func main() {
	q := celeste.Query{}
	if err := json.Unmarshal([]byte(q3), &q); err != nil {
		panic(err)
	}

	fmt.Println(q)
}
