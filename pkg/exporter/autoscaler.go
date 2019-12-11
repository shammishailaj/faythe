// Copyright (c) 2019 Kien Nguyen-Tuan <kiennt2609@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package exporter

import "github.com/prometheus/client_golang/prometheus"

// Set of raw Prometheus metrics.
// Do not increment directly, use Report* methods.
var (
	numberOfScalers = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "faythe",
			Subsystem: "autoscaler",
			Name:      "workers_total",
			Help:      "The total number of scalers are currently managed by this cluster member.",
		},
		[]string{"cluster"})

	failureScalerActionCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "faythe",
			Subsystem: "autoscaler",
			Name:      "action_failures_total",
			Help:      "The total number of scaler action failures.",
		},
		[]string{"cluster", "type"})
)

func init() {
	prometheus.MustRegister(numberOfScalers)
	prometheus.MustRegister(failureScalerActionCounter)
}

func ReportNumScalers(clusterID string, val float64) {
	if val == 0 {
		numberOfScalers.WithLabelValues(clusterID).Set(val)
	} else if val < 0 {
		numberOfScalers.WithLabelValues(clusterID).Sub(val)
	} else {
		numberOfScalers.WithLabelValues(clusterID).Add(val)
	}
}

func ReportFailureScalerActionCounter(clusterID, actionType string) {
	failureScalerActionCounter.WithLabelValues(clusterID, actionType).Inc()
}
