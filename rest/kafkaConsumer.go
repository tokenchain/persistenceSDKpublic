package rest

import (
	"github.com/Shopify/sarama"
	"github.com/commitHub/commitBlockchain/wire"
)

//NewConsumer : is a consumer which is needed to create child consumers to consume topics
func NewConsumer(kafkaPorts []string) sarama.Consumer {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer(kafkaPorts, config)
	if err != nil {
		panic(err)
	}
	return consumer
}

//PartitionConsumers : is a child consumer
func PartitionConsumers(consumer sarama.Consumer, topic string) sarama.PartitionConsumer {
	//partition and offset defined in CONSTANTS.go
	partitionConsumer, err := consumer.ConsumePartition(topic, partition, offset)
	if err != nil {
		panic(err)
	}
	return partitionConsumer
}

//KafkaTopicConsumer : Takes a consumer and makes it consume a topic message at a time
func KafkaTopicConsumer(topic string, consumers map[string]sarama.PartitionConsumer, cdc *wire.Codec) KafkaMsg {

	partitionConsumer := consumers[topic]

	if len(partitionConsumer.Messages()) == 0 {
		var kafkaStore = KafkaMsg{Msg: nil}
		return kafkaStore
	}
	kafkaMsg := <-partitionConsumer.Messages()
	var kafkaStore KafkaMsg
	err := cdc.UnmarshalJSON([]byte(kafkaMsg.Value), &kafkaStore)
	if err != nil {
		panic(err)
	}
	return kafkaStore
}
