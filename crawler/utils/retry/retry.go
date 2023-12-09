package retry

import (
	"novel_crawler/global/consts"
	"time"
)

func Retry(task func() error, count int) error {
	if err := task(); err == nil {
		return nil
	} else if count > 1 {
		time.Sleep(consts.RetrySleep)
		return Retry(task, count-1)
	} else {
		return err
	}
}
