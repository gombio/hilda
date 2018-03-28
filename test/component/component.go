package component

import (
	ht "github.com/gombio/hilda/test"
)

//Component test closure
type Component func(c *ht.Context, r *ht.Report)
