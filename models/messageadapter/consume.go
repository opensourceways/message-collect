package messageadapter

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
)

func ConsumeGroup(cfg ConsumeConfig, handler sarama.ConsumerGroupHandler) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Return.Errors = true
	fmt.Println("start connect kafka")
	// 开始连接kafka服务器
	group, err := sarama.NewConsumerGroup([]string{cfg.Address}, cfg.Group, config)

	if err != nil {
		fmt.Println("connect kafka failed; err", err)
		return
	}
	// 检查错误
	go func() {
		for err := range group.Errors() {
			fmt.Println("group errors : ", err)
		}
	}()

	ctx := context.Background()
	fmt.Println("start get msg")
	// for 是应对 consumer rebalance
	for {
		// 需要监听的主题
		topics := []string{cfg.Topic}
		// 启动kafka消费组模式，消费的逻辑在上面的 ConsumeClaim 这个方法里
		err := group.Consume(ctx, topics, handler)
		if err != nil {
			fmt.Println("consume failed; err : ", err)
			return
		}
	}
}
