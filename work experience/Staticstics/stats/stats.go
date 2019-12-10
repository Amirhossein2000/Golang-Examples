package stats

func InitStats() {
	InitAgentx()
	InitTallyStats()
}

func Shutdown() {
	ShutdownAgentx()
	ShutdownTallyStats()
}
