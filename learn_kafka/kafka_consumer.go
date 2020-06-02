// Example function-based high-level Apache Kafka consumer
package learn_kafka

/**
 * Copyright 2016 Confluent Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// consumer_example implements a consumer using the non-channel Poll() API
// to retrieve messages and events.

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"os/signal"
	"syscall"
)

func StartReading() {

	//if len(os.Args) < 4 {
	//	fmt.Fprintf(os.Stderr, "Usage: %s <broker> <group> <topics..>\n",
	//		os.Args[0])
	//	os.Exit(1)
	//}

	broker := "localhost:9092"
	group := "random"
	topics := [] string {"goshak_test_topic"}
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		// Avoid connecting to IPv6 brokers:
		// This is needed for the ErrAllBrokersDown show-case below
		// when using localhost brokers on OSX, since the OSX resolver
		// will return the IPv6 addresses first.
		// You typically don't need to specify this configuration property.
		"broker.address.family": "v4",
		"group.id":              group,
		"session.timeout.ms":    6000,
		"auto.offset.reset":     "earliest"})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}
	err = c.SubscribeTopics(topics, nil)

	for {
		message, err := c.ReadMessage(-1)
		if err != nil {
			fmt.Println("Kafka read messages from kafka failed:", err)
		}

		if message != nil {
			fmt.Printf("%v", message.Value)
		}
	}

}
