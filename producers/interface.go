// Copyright 2016 Mesosphere, Inc.
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

package producers

// MetricsProducer defines an interface that the various producers must
// implement in order to receive, process, and present metrics to the caller or
// client. All producers must use the MetricsMessage structure to receive
// metrics, and they must implement their own struct for handling configuration.
//
// Further, although it isn't defined in this interface, it is recommended that
// producers must also create their own MetricsMessage channel to be
// used both in the implementation (e.g., &producerImpl{}) and to be returned
// to the caller. Doing so ensures that, in the future, multiple producers
// can be enabled at once (each producer has a dedicated chan).
type MetricsProducer interface {
	Run() error
}

// MetricsMessage defines the structure of the metrics being sent to the various
// producers.
type MetricsMessage struct {
	// Classification is an arbitrary string that can be used to identify or
	// "tag" the outgoing data.
	//
	//   * For the Kafka producer, this might be the Kafka topic.
	//   * For the HTTP producer, this might be the endpoint that the metrics
	//     will be available at.
	//
	Classification string
	Data           interface{}
}
