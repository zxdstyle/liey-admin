package commands

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/spf13/cobra"
	"github.com/zxdstyle/liey-admin/framework/queue"
	"sync"
)

var (
	wg           = &sync.WaitGroup{}
	QueueCommand = &cobra.Command{
		Use:   "queue [name]",
		Short: "listen to these given queues",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()

			for _, name := range args {
				wg.Add(1)
				q := queue.GetQueue(name)
				if q == nil {
					g.Log().Fatal(ctx, fmt.Errorf("not found queue: %s", name))
				}
				go func(q *queue.Queue) {
					defer wg.Done()
					if err := q.Listen(); err != nil {
						g.Log().Fatal(ctx, err)
					}
				}(q)
			}

			wg.Wait()
		},
	}
)
