package kafka

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type ConsumeConfig struct {
	Topic    string `yaml:"topic"`
	Address  string `yaml:"address"`
	Group    string `yaml:"group"`
	Offset   int64  `yaml:"offset"`
	UserName string `yaml:"user_name"`
	Password string `yaml:"password"`
	MqCert   string `yaml:"mq_cert"`
}

func ConsumeGroup(cfg ConsumeConfig, handler sarama.ConsumerGroupHandler) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = cfg.Offset
	config.Consumer.Return.Errors = true
	if cfg.UserName != "" {
		config.Net.SASL.Enable = true
		config.Net.SASL.User = cfg.UserName
		config.Net.SASL.Password = cfg.Password
		config.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512

		config.Net.TLS.Enable = true
		tlsConfig := &tls.Config{}

		if cfg.MqCert != "" {
			caCert, err := ioutil.ReadFile(cfg.MqCert)
			if err != nil {
				logrus.Errorf("无法加载证书, %v", err)
				return
			}
			caCertPool := x509.NewCertPool()
			caCertPool.AppendCertsFromPEM(caCert)
			tlsConfig.RootCAs = caCertPool
		}
		config.Net.TLS.Config = tlsConfig
	}
	// 开始连接kafka服务器
	group, err := sarama.NewConsumerGroup(strings.Split(cfg.Address, ","), cfg.Group, config)

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
	fmt.Println("start get msg " + cfg.Topic)
	// for 是应对 consumer rebalance

	// 需要监听的主题
	topics := []string{cfg.Topic}
	// 启动kafka消费组模式，消费的逻辑在上面的 ConsumeClaim 这个方法里
	err = group.Consume(ctx, topics, handler)
	if err != nil {
		fmt.Println("consume failed; err : ", err)
		return
	}
}
