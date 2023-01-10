package celeste

import (
	"encoding/json"
	"testing"
)

func TestQuery_String(t *testing.T) {
	var tests = []struct {
		rawJSON string
		want    string
	}{
		{
			rawJSON: `{"source":"people","find":{"$and":[{"name":{"$not_like":"\"%james%\""}}]}}`,
			want:    "SELECT * FROM `people` WHERE name NOT LIKE \"%james%\"",
		},
		{
			rawJSON: `{"source":"people","find":{"$and":[{"name":{"$in":["\"james\"","\"bond\""]}}]}}`,
			want:    "SELECT * FROM `people` WHERE name IN (\"james\", \"bond\")",
		},
		{
			rawJSON: `{"source":"people","find":{"$and":[{"name":"\"james\""},{"age":{"$gt":20}}]}}`,
			want:    "SELECT * FROM `people` WHERE name = \"james\" AND age > 20",
		},
	}

	for _, item := range tests {
		t.Run(item.rawJSON, func(t *testing.T) {
			q := Query{}
			if err := json.Unmarshal([]byte(item.rawJSON), &q); err != nil {
				t.Errorf("want %q, got ERROR %q", item.want, err.Error())
			} else if item.want != q.String() {
				t.Errorf("want %s, got %s", item.want, q.String())
			}
		})
	}
}
