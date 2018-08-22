package config

type ProducerConfig struct {
	PolicyName    string
	PolicyContext ProducerPolicy
}

type ProducerPolicy struct {
	// for solo
	Timer uint64
	// reserve
	Num uint64
}
