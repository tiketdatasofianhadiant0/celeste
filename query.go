package celeste

import (
	"strings"
)

type Query struct {
	Source  string     `json:"source"`
	Columns []string   `json:"columns,omitempty"`
	Find    *FindToken `json:"find,omitempty"`
}

func (q Query) String() string {
	cols := "SELECT * "
	if len(q.Columns) > 0 {
		cols = "SELECT " + strings.Join(q.Columns, ", ") + " "
	}

	src := ""
	if q.Source != "" {
		src = "FROM `" + q.Source + "` "
	}

	filter := ""
	if q.Find != nil {
		filter = "WHERE " + q.Find.String()
	}

	if src == "" {
		return ""
	}

	return strings.TrimSpace(cols + src + filter)
}
