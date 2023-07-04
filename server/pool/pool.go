package pool

import "github.com/panjf2000/ants/v2"

var DOMAIN_SCAN *ants.Pool
var DOMAIN_SCAN2 *ants.PoolWithFunc

func init() {
	DOMAIN_SCAN, _ = ants.NewPool(1000)
	DOMAIN_SCAN2, _ = ants.NewPoolWithFunc(1000, func(i interface{}) {

	})
}