package gc

import (
	"github.com/xingdl2007/doozerd/store"
	"time"
)

func Clean(st *store.Store, keep int64, ticker <-chan time.Time) {
	for _ = range ticker {
		last := (<-st.Seqns) - keep
		st.Clean(last)
	}
}
