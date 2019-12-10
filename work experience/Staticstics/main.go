package main

import (
	"fmt"
	"staticstics/stats"
	"time"
)

func main() {
	stats.InitStats()
	nowTime := time.Now()
	time.Sleep(2 * time.Second)

	fmt.Println(*stats.TallyStatsManager.DiamResponseTimeGuage)

	(*stats.TallyStatsManager.DiamResponseTimeGuage).Update(float64(time.Since(nowTime) / time.Millisecond))

	fmt.Println(*stats.TallyStatsManager.DiamResponseTimeGuage)

	nowTime = time.Now()
	time.Sleep(4 * time.Second)

	(*stats.TallyStatsManager.DiamResponseTimeGuage).Update(float64(time.Since(nowTime) / time.Millisecond))

	fmt.Println(*stats.TallyStatsManager.DiamResponseTimeGuage)
}
