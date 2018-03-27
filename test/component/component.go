package component

import (
	ht "github.com/gombio/hilda/test"
)

type Component func(c *ht.Context, r *ht.Report)
