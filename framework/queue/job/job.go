package job

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/nsqio/go-nsq"
	"github.com/zxdstyle/liey-admin/framework/exception"
	"github.com/zxdstyle/liey-admin/framework/queue/serializer"
	"log"
	"reflect"
)

var (
	jobs = gmap.NewStrAnyMap(true)
)

type (
	Job interface {
		serializer.Serializer
		Handle(ctx context.Context, payload []byte) error
	}

	SpecifiedName interface {
		Name() string
	}

	// SpecifiedQueue 指定队列
	SpecifiedQueue interface {
		QueueName() string
	}

	SpecifiedConcurrency interface {
		Concurrency() int
	}

	SpecifiedChannel interface {
		Channel() string
	}

	Handler func(ctx context.Context, payload []byte) error
)

func RegisterJob(ctx context.Context, job ...Job) error {
	for _, j := range job {
		name := ResolveJobName(j)
		if !nsq.IsValidTopicName(name) {
			return fmt.Errorf("invalid queue name: %s", name)
		}
		if jobs.Contains(name) {
			return gerror.NewCode(exception.CodeInternalError, fmt.Sprintf("queue `%s` already exists", name))
		}
		jobs.Set(name, j)
	}
	return nil
}

func GetJob(name string) Job {
	val, ok := jobs.Search(name)
	if !ok || val == nil {
		return nil
	}
	return val.(Job)
}

func GetJobWithQueue(queues ...string) (js []Job) {
	qs := garray.NewStrArrayFrom(queues)
	jobs.Iterator(func(k string, v interface{}) bool {
		j := v.(Job)
		queue := ResolveQueue(j)
		if qs.Contains(queue) {
			js = append(js, j)
		}
		return true
	})
	return
}

func ResolveJobName(j interface{}) string {
	if val, ok := j.(SpecifiedName); ok {
		return val.Name()
	}
	t := reflect.TypeOf(j)
	return t.Name()
}

func ResolveQueue(j Job) string {
	name := "default"
	if val, ok := j.(SpecifiedQueue); ok {
		name = val.QueueName()
	}
	return name
}

func Connect() {
	cfg := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", cfg)
	if err != nil {
		log.Fatal(err)
	}

	messageBody := []byte("hello")
	topicName := "topic"

	// Synchronously publish a single message to the specified topic.
	// Messages can also be sent asynchronously and/or in batches.
	err = producer.Publish(topicName, messageBody)
	if err != nil {
		log.Fatal(err)
	}

	// Gracefully stop the producer when appropriate (e.g. before shutting down the service)
	producer.Stop()

}
