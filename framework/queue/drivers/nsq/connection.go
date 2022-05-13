package nsq

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/nsqio/go-nsq"
	"github.com/zxdstyle/liey-admin/framework/queue/job"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Connection struct {
	producer *nsq.Producer
	cfg      Config
}

func NewConnection(cfg Config) (*Connection, error) {
	conf := nsq.NewConfig()
	producer, err := nsq.NewProducer(cfg.Host, conf)
	if err != nil {
		return nil, err
	}
	if err := producer.Ping(); err != nil {
		return nil, err
	}
	return &Connection{
		producer: producer,
		cfg:      cfg,
	}, nil
}

func (c *Connection) Produce(name string, payload []byte) error {
	return c.producer.Publish(name, payload)
}

func (c *Connection) Close() {
	c.producer.Stop()
}

func (c *Connection) Consume(j job.Job) error {
	name := job.ResolveJobName(j)

	concurrency := 1
	if val, ok := j.(job.SpecifiedConcurrency); ok {
		concurrency = val.Concurrency()
	}

	channel := "default"
	if val, ok := j.(job.SpecifiedChannel); ok {
		channel = val.Channel()
	}

	config := nsq.NewConfig()
	if err := c.checkTopic(name); err != nil {
		return err
	}
	if err := c.checkChannel(name, channel); err != nil {
		return err
	}

	config.MaxInFlight = concurrency
	consumer, err := nsq.NewConsumer(name, channel, config)
	if err != nil {
		log.Fatal(err)
	}
	consumer.SetLoggerLevel(nsq.LogLevelError)
	consumer.AddConcurrentHandlers(newDefaultHandler(j.Handle), concurrency)

	addr := c.cfg.Host
	if len(c.cfg.Lookupd) > 0 {
		addr = c.cfg.Lookupd
	}
	if len(addr) == 0 {
		return fmt.Errorf("invalid nsq addr")
	}

	if e := consumer.ConnectToNSQLookupd(addr); e != nil {
		return e
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	consumer.Stop()
	return nil
}

func (c *Connection) checkTopic(name string) error {
	url := fmt.Sprintf("%s/topic/create?topic=%s", c.cfg.Lookupd, name)
	resp, err := g.Client().Post(context.Background(), url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create queue %s: ", resp.ReadAllString())
	}
	return nil
}

func (c *Connection) checkChannel(topic, channel string) error {
	url := fmt.Sprintf("%s/channel/create?channel=%s&topic=%s", c.cfg.Lookupd, channel, topic)
	resp, err := g.Client().Post(context.Background(), url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create channel %s: ", resp.ReadAllString())
	}
	return nil
}
