package internal

import (
	"sync"
)

// App is a global singleton GlobalApp instance
func App() *GlobalApp { return appUniqueInstance }

// ---------------------------------------------

// func sel(ss ...string) (ret string) {
// 	for _, s := range ss {
// 		if len(s) > 0 {
// 			ret = s
// 			break
// 		}
// 	}
// 	return
// }

// ---------------------------------------------

// GlobalApp is a general global object
type GlobalApp struct {
	muInit sync.RWMutex

	// dbx    dbl.DB
	// cache  *cache.Hub
	// cron   cron.Jobs
}

var onceForApp sync.Once
var appUniqueInstance *GlobalApp

func init() {
	onceForApp.Do(func() {
		appUniqueInstance = &GlobalApp{}
	})
}
