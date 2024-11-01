package historywatch

import (
	"context"
	"fmt"
	historywatchDom "test_dealls/src/business/domain/history_watch"
	"test_dealls/src/entity"
	"test_dealls/src/utils/log"
	"time"

	"github.com/robfig/cron/v3"
)

type Interface interface{}

type HistoryWatch struct {
	log log.Interface
	dom domain
}

type domain struct {
	historyWatch historywatchDom.Interface
}

func Init(log log.Interface, historyWatchD historywatchDom.Interface) Interface {
	historyWatch := &HistoryWatch{
		log: log,
		dom: domain{
			historyWatch: historyWatchD,
		},
	}
	initScheduller(context.Background(), *historyWatch)

	return historyWatch
}

func initScheduller(ctx context.Context, hw HistoryWatch) {
	fmt.Println("Initializing scheduler")
	c := cron.New(cron.WithLocation(time.FixedZone("Asia/Jakarta", 7*60*60))) // Set timezone to Asia/Jakarta

	c.AddFunc("0 0 * * *", func() { deleteHistoryWatch(ctx, hw) }) // Run daily at 00:00
	c.Start()
}

func deleteHistoryWatch(ctx context.Context, hw HistoryWatch) {
	hw.log.Info(ctx, "Running the task of delete history watch daily")

	listHistoryWatch, err := hw.dom.historyWatch.GetList(ctx, entity.HistoryWatch{})
	if err != nil {
		hw.log.Error(ctx, fmt.Sprintf("Error when getting list of history watch, err is: %v", err))
		return
	}

	for _, v := range listHistoryWatch {
		_, err := hw.dom.historyWatch.Delete(ctx, entity.HistoryWatch{
			ID: v.ID,
		})
		if err != nil {
			hw.log.Error(ctx, err.Error())
			continue
		}
	}

	hw.log.Info(ctx, "Task completed")
}
