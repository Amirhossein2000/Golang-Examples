package stats

import (
	"io"
	"log"
	"time"

	"github.com/uber-go/tally"
)

var TallyStatsManager *TallyStats

type TallyStats struct {
	DiamResponseTimeGuage *tally.Gauge
	gauges                map[string]float64
	closer                io.Closer
}

type StatsReporter struct {
}

func newStatsReporter() tally.StatsReporter {
	return &StatsReporter{}
}

func (r *StatsReporter) ReportCounter(name string, _ map[string]string, value int64) {
}

func (r *StatsReporter) ReportGauge(name string, _ map[string]string, value float64) {
	TallyStatsManager.gauges[name] = value
}

func (r *StatsReporter) ReportTimer(name string, _ map[string]string, interval time.Duration) {
}

func (r *StatsReporter) ReportHistogramValueSamples(
	name string,
	_ map[string]string,
	_ tally.Buckets,
	bucketLowerBound,
	bucketUpperBound float64,
	samples int64,
) {
}

func (r *StatsReporter) ReportHistogramDurationSamples(
	name string,
	_ map[string]string,
	_ tally.Buckets,
	bucketLowerBound,
	bucketUpperBound time.Duration,
	samples int64,
) {
}

func (r *StatsReporter) Capabilities() tally.Capabilities {
	return r
}

func (r *StatsReporter) Reporting() bool {
	return true
}

func (r *StatsReporter) Tagging() bool {
	return false
}

func (r *StatsReporter) Flush() {
}

func InitTallyStats() {
	reporter := newStatsReporter()
	rootScope, closer := tally.NewRootScope(tally.ScopeOptions{
		Reporter: reporter,
	}, 5*time.Second)

	subScope := rootScope.SubScope("diameter")
	diamResponseTimeGuage := subScope.Gauge("diameter_response_time")
	TallyStatsManager = &TallyStats{DiamResponseTimeGuage: &diamResponseTimeGuage,
		gauges: make(map[string]float64),
		closer: closer}

	log.Println("Tally stats Initialized")
}

func ShutdownTallyStats() {
	TallyStatsManager.closer.Close()
}
