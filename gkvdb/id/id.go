package id

import (
	"github.com/sanrentai/snowflake"
)

var s *snowflake.Snowflake

func init() {
	s, _ = snowflake.New(1)
}

func ID() int64 {
	return s.Generate().Int64()
}
