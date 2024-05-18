package schedule

import (
	"context"
	"time"

	"github.com/zilanlann/acmer-manage-system/server/global"
	"github.com/zilanlann/acmer-manage-system/server/pkg/cf"
)

func StartScheduler(ctx context.Context) {
	scheduleTask()

	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			scheduleTask()
		}
	}
}

func scheduleTask() {
	if err := cf.RefreshContests(); err != nil {
		global.LOG.Error(err.Error())
	}
	if err := cf.RefreshAllUserSubmisions(); err != nil {
		global.LOG.Error(err.Error())
	}
}
