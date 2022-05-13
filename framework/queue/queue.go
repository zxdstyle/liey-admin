package queue

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/zxdstyle/liey-admin/framework/exception"
	"github.com/zxdstyle/liey-admin/framework/queue/config"
	"github.com/zxdstyle/liey-admin/framework/queue/drivers"
	"github.com/zxdstyle/liey-admin/framework/queue/drivers/contract"
	"github.com/zxdstyle/liey-admin/framework/queue/job"
	"sync"
)

var queues = gmap.NewStrAnyMap(true)

type Queue struct {
	name   string
	cfg    config.Queue
	driver drivers.Driver
	conn   contract.Connection
}

func newQueueWithConfig(name string, cfg config.Queue) (*Queue, error) {
	dr, er := drivers.GetDriver(cfg.Driver)
	if er != nil {
		return nil, er
	}
	conn, err := dr.Connect(cfg.Connection)
	if err != nil {
		return nil, err
	}
	return &Queue{
		name:   name,
		conn:   conn,
		driver: dr,
		cfg:    cfg,
	}, nil
}

// Listen 监听队列
func (q *Queue) Listen() error {
	js := job.GetJobWithQueue(q.name)
	wg := &sync.WaitGroup{}
	for _, v := range js {
		wg.Add(1)

		go func(j job.Job) {
			if err := q.conn.Consume(j); err != nil {
				g.Log().Error(context.Background(), err)
			}
			wg.Done()
		}(v)
	}
	wg.Wait()
	return nil
}

func (q Queue) Close() {
	q.conn.Close()
}

func InitQueueWithConfig() {
	ctx := context.Background()
	qs, er := g.Cfg("queue").Get(ctx, "queues")
	if er != nil {
		g.Log().Error(ctx, er)
	}
	var queue map[string]config.Queue
	if err := qs.Scan(&queue); err != nil {
		g.Log().Error(ctx, err)
	}

	for name, q := range queue {
		qu, err := newQueueWithConfig(name, q)
		if err != nil {
			g.Log().Error(ctx, err)
			continue
		}
		if e := Register(name, qu); e != nil {
			g.Log().Error(ctx, e)
		}
	}
}

// Register 注册队列
func Register(name string, q *Queue) error {
	if queues.Contains(name) {
		return gerror.NewCode(exception.CodeInternalError, fmt.Sprintf("queue `%s` already exists", name))
	}
	queues.Set(name, q)
	return nil
}

// GetQueue 获取队列
func GetQueue(name string) *Queue {
	val, ok := queues.Search(name)
	if !ok || val == nil {
		return nil
	}
	return val.(*Queue)
}

func Dispatch(j job.Job, payload interface{}) error {
	name := job.ResolveQueue(j)
	q := GetQueue(name)
	if q == nil {
		return fmt.Errorf("invalid queue: %s", name)
	}
	p, err := j.Serialize(payload)
	if err != nil {
		return err
	}
	return q.conn.Produce(job.ResolveJobName(j), p)
}
