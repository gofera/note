# client: go-nsq

Consume
new config
	set default
new consumer (topic, channel, config)
	go ready loop
		for ticker (per 5s)
			for each connection
				update RDY with count 0
					conn send RDY command
consumer add handler
consumer connect to NSQD (addr)
	new connection (addr, config, delegate => consumer)
	connection connect
		tcp dial addr
		write MagicV2
		write command identify
		
