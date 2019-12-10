
ConsumeLoop:
	for i := 0; ; {

		if c == nil {
			c, err = connect(kafkaServer)
			if err != nil {
				logger.Error("Can not connect to kafka: %s", err)
				time.Sleep(5 * time.Second)
				continue
			}
		}

		select {
		case ev := <-c.Events():
			switch e := ev.(type) {
			case kafka.AssignedPartitions:
				logger.Info("%% %v\n", e)
				c.Assign(e.Partitions)
			case kafka.RevokedPartitions:
				logger.Info("%% %v\n", e)
				c.Unassign()
			case *kafka.Message:
				dataflow_chan <- e.Value
				i += 1

				if i%LOG_CONSUME_COUNT == 0 {
					i = 0
					logger.Debug("Consumed %d Kafka Messages", LOG_CONSUME_COUNT)
				}

			case kafka.PartitionEOF:
				logger.Debug("%% Reached %v\n", e)
			case kafka.Error:
				logger.Error("%% Error: %v\n", e)
				c.Close()
				c = nil
			}
		case <-ctx.Done():
			c.Close()
			c = nil
			logger.Info("Shutdown %s consumer successfully.", kafkaServer.Name)
			break ConsumeLoop
		}
	}
