/*
 * Copyright (c) 2021 Enix, SAS
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 *
 * Authors:
 * Paul Laffitte <paul.laffitte@enix.fr>
 * Alexandre Buisine <alexandre.buisine@enix.fr>
 */

package dothill

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	APICallMetric = "dothill_api_appliance_api_call"
	APICallHelp   = "How many API calls have been executed"

	APICallDurationMetric = "dothill_api_appliance_api_call_duration"
	APICallDurationHelp   = "The total duration of API calls"
)

type Collector struct {
	apiCall         *prometheus.CounterVec
	apiCallDuration *prometheus.CounterVec
}

func newCollector() *Collector {
	return &Collector{
		apiCall: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: APICallMetric,
				Help: APICallHelp,
			},
			[]string{"endpoint", "success"},
		),
		apiCallDuration: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: APICallDurationMetric,
				Help: APICallDurationHelp,
			},
			[]string{"endpoint"},
		),
	}
}

func (collector *Collector) Describe(ch chan<- *prometheus.Desc) {
	collector.apiCall.Describe(ch)
	collector.apiCallDuration.Describe(ch)
}

func (collector *Collector) Collect(ch chan<- prometheus.Metric) {
	collector.apiCall.Collect(ch)
	collector.apiCallDuration.Collect(ch)
}

func (collector *Collector) trackAPICall(endpoint string) func(bool) {
	start := time.Now()

	return func(success bool) {
		duration := float64(time.Since(start).Nanoseconds()) / 1000 / 1000 / 1000
		collector.apiCallDuration.WithLabelValues(endpoint).Add(duration)
		collector.apiCall.WithLabelValues(endpoint, fmt.Sprintf("%t", success)).Inc()
	}
}
