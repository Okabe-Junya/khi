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

package types

import (
	"github.com/GoogleCloudPlatform/khi/pkg/log"
	"github.com/GoogleCloudPlatform/khi/pkg/log/structure"
	"github.com/GoogleCloudPlatform/khi/pkg/model"
	"github.com/GoogleCloudPlatform/khi/pkg/source/gcp/task/gke/k8s_audit/rtype"
)

// ResourceSpecificParserInput is a type passed to ResourceSpecificParser from the prestep parser.
type ResourceSpecificParserInput struct {
	Request                                *structure.Reader
	ResourceBodyReader                     *structure.Reader
	Response                               *structure.Reader
	Log                                    *log.LogEntity
	Operation                              *model.KubernetesObjectOperation
	PrincipalEmail                         string
	MethodName                             string
	ResourceName                           string
	ResourceBodyYaml                       string
	RequestType                            rtype.Type
	ResponseType                           rtype.Type
	Code                                   int
	GeneratedFromDeleteCollectionOperation bool
}

type TimelineGrouperResult struct {
	TimelineResourcePath string
	PreParsedLogs        []*ResourceSpecificParserInput
}
