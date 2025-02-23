// Copyright 2024 Google LLC
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

package history

import (
	"time"

	"github.com/GoogleCloudPlatform/khi/pkg/model/binarychunk"
	"github.com/GoogleCloudPlatform/khi/pkg/model/enum"
)

// The entire inspection data.
type History struct {
	Version   string                 `json:"version"`
	Metadata  map[string]interface{} `json:"metadata"`
	Logs      []*SerializableLog     `json:"logs"`
	Timelines []*ResourceTimeline    `json:"timelines"`
	Resources []*Resource            `json:"resources"`
}

type Resource struct {
	ResourceName     string                  `json:"name"`
	Timeline         string                  `json:"timeline"`
	FullResourcePath string                  `json:"path"`
	Children         []*Resource             `json:"children"`
	Relationship     enum.ParentRelationship `json:"relationship"`
}

type ResourceTimeline struct {
	ID        string              `json:"id"`
	Revisions []*ResourceRevision `json:"revisions"`
	Events    []*ResourceEvent    `json:"events"`
}

type ResourceRevision struct {
	ChangeTime time.Time                    `json:"changeTime"`
	Requestor  *binarychunk.BinaryReference `json:"requestor"`
	Body       *binarychunk.BinaryReference `json:"body"`
	Log        string                       `json:"log"`
	Verb       enum.RevisionVerb            `json:"verb"`
	State      enum.RevisionState           `json:"state"`
	Partial    bool                         `json:"partial"`
}

type ResourceEvent struct {
	Log string `json:"log"`
}

type SerializableLog struct {
	Timestamp   time.Time                    `json:"ts"`
	Body        *binarychunk.BinaryReference `json:"body"`
	Summary     *binarychunk.BinaryReference `json:"summary"`
	ID          string                       `json:"id"`
	DisplayId   string                       `json:"displayId"`
	Annotations []any                        `json:"annotations"`
	Type        enum.LogType                 `json:"type"`
	Severity    enum.Severity                `json:"severity"`
}

func NewHistory() *History {
	return &History{
		Version:   "5",
		Timelines: make([]*ResourceTimeline, 0),
		Logs:      make([]*SerializableLog, 0),
		Resources: make([]*Resource, 0),
	}
}

func newTimeline(tid string) *ResourceTimeline {
	return &ResourceTimeline{
		ID:        tid,
		Revisions: make([]*ResourceRevision, 0),
		Events:    make([]*ResourceEvent, 0),
	}
}
