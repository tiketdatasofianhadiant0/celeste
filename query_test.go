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
			rawJSON: `{"source":"people"}`,
			want:    "SELECT * FROM `people`",
		},
		{
			rawJSON: `{"source":"people","columns":["name","age","gender"]}`,
			want:    "SELECT name, age, gender FROM `people`",
		},
		{
			rawJSON: `{"source":"people","find":{"$and":[{"name":{"$like":"\"%james%\""}}]}}`,
			want:    "SELECT * FROM `people` WHERE name LIKE \"%james%\"",
		},
		{
			rawJSON: `{"source":"people","find":{"$and":[{"name":{"$not_like":"\"%james%\""}}]}}`,
			want:    "SELECT * FROM `people` WHERE name NOT LIKE \"%james%\"",
		},
		{
			rawJSON: `{"source":"people","find":{"$and":[{"name":{"$in":["\"james\"","\"bond\""]}}]}}`,
			want:    "SELECT * FROM `people` WHERE name IN (\"james\", \"bond\")",
		},
		{
			rawJSON: `{"source":"people","find":{"$and":[{"name":{"$not_in":["\"james\"","\"bond\""]}}]}}`,
			want:    "SELECT * FROM `people` WHERE name NOT IN (\"james\", \"bond\")",
		},
		{
			rawJSON: `{"source":"people","find":{"$and":[{"name":{"$eq":{"$upper":"\"james\""}}}]}}`,
			want:    "SELECT * FROM `people` WHERE name = UPPER(\"james\")",
		},
		{
			rawJSON: `{"source":"people","find":{"$and":[{"name":{"$eq":{"$lower":"\"james\""}}}]}}`,
			want:    "SELECT * FROM `people` WHERE name = LOWER(\"james\")",
		},
		{
			rawJSON: `{"source":"people","find":{"$and":[{"name":{"$eq":{"$trim":"\"james\""}}}]}}`,
			want:    "SELECT * FROM `people` WHERE name = TRIM(\"james\")",
		},
		{
			rawJSON: `{"source":"people","find":{"$and":[{"$upper":{"name":{"$eq":"\"JAMES\""}}}]}}`,
			want:    "SELECT * FROM `people` WHERE UPPER(name) = \"JAMES\"",
		},
		{
			rawJSON: `{"source":"people","find":{"$and":[{"$lower":{"name":{"$eq":"\"james\""}}}]}}`,
			want:    "SELECT * FROM `people` WHERE LOWER(name) = \"james\"",
		},
		{
			rawJSON: `{"source":"people","find":{"$and":[{"$trim":{"name":{"$eq":"\"james\""}}}]}}`,
			want:    "SELECT * FROM `people` WHERE TRIM(name) = \"james\"",
		},
		{
			rawJSON: `{"source":"people","find":{"$and":[{"name":"\"james\""},{"age":{"$gt":20}}]}}`,
			want:    "SELECT * FROM `people` WHERE name = \"james\" AND age > 20",
		},
		{
			rawJSON: `{"source":"people","find":{"$or":[{"name":"\"james\""},{"age":{"$gt":20}}]}}`,
			want:    "SELECT * FROM `people` WHERE name = \"james\" OR age > 20",
		},
		{
			rawJSON: `{"source":"people","find":{"$and":[{"$is_null":"name"}]}}`,
			want:    "SELECT * FROM `people` WHERE name IS NULL",
		},
		{
			rawJSON: `{"source":"people","find":{"$and":[{"$is_not_null":"name"}]}}`,
			want:    "SELECT * FROM `people` WHERE name IS NOT NULL",
		},
		{
			rawJSON: `{"source":"people","find":{"$or":[{"$is_null":"name"},{"age":{"$gt":20}}]}}`,
			want:    "SELECT * FROM `people` WHERE name IS NULL OR age > 20",
		},
		{
			rawJSON: `{"source":"people","find":{"$and":[{"$is_null":"name"},{"$is_not_null":"birthday"}]}}`,
			want:    "SELECT * FROM `people` WHERE name IS NULL AND birthday IS NOT NULL",
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
