// Copyright 2022 The Codefresh Authors.
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

package store

import (
	"fmt"
	"runtime"
	"time"

	"github.com/Masterminds/semver/v3"
)

var s Store

var (
	binaryName               = "cli-v2"
	version                  = "v99.99.99"
	buildDate                = ""
	gitCommit                = ""
	segmentWriteKey          = ""
	maxDefVersion            = "2.0.0"
	RuntimeDefURL            = "manifests/runtime.yaml"
	AddClusterDefURL         = "https://github.com/codefresh-io/csdp-official/add-cluster/kustomize"
	FallbackAddClusterDefURL = "https://github.com/codefresh-io/cli-v2/manifests/add-cluster/kustomize"
)

type Version struct {
	Version    *semver.Version
	BuildDate  string
	GitCommit  string
	GoVersion  string
	GoCompiler string
	Platform   string
}
type Store struct {
	AddClusterJobName                 string
	ArgoCDServerName                  string
	ArgoCDTokenKey                    string
	ArgoCDTokenSecret                 string
	ArgoWFServiceName                 string
	ArgoWFServicePort                 int32
	BinaryName                        string
	Codefresh                         string
	CFComponentType                   string
	CFGitSourceType                   string
	CFRuntimeDefType                  string
	CFRuntimeType                     string
	CFTokenSecret                     string
	CFTokenSecretKey                  string
	CFStoreIVSecretKey                string
	CodefreshCM                       string
	CodefreshSA                       string
	ComponentsReporterName            string
	ComponentsReporterSA              string
	ComponentsReporterURL             string
	DefaultAPI                        string
	EventBusName                      string
	EventReportingEndpoint            string
	EventsReporterName                string
	GitSourceName                     string
	WorkflowsIngressName              string
	WorkflowsIngressPath              string
	AppProxyIngressName               string
	AppProxyIngressPath               string
	AppProxyServicePort               int32
	AppProxyServiceName               string
	DocsLink                          string
	LabelKeyCFType                    string
	LabelKeyCFInternal                string
	AnnotationKeySyncWave             string
	MarketplaceGitSourceName          string
	MarketplaceRepo                   string
	MaxDefVersion                     *semver.Version
	RuntimeDefURL                     string
	Version                           Version
	WaitTimeout                       time.Duration
	WorkflowName                      string
	WorkflowReporterName              string
	WorkflowTriggerServiceAccount     string
	DemoCalendarSensorObjectName      string
	DemoCalendarSensorFileName        string
	DemoCalendarEventSourceFileName   string
	DemoCalendarEventSourceObjectName string
	DemoCalendarEventName             string
	DemoCalendarDependencyName        string
	DemoWorkflowTemplateFileName      string
	DemoWorkflowTemplateName          string
	DemoGitEventSourceFileName        string
	DemoGitEventSourceObjectName      string
	WebhooksRootPath                  string
	DemoGitEventSourceTargetPort      string
	DemoGitEventSourceServicePort     int32
	DemoPipelinesIngressFileName      string
	DemoPipelinesIngressObjectName    string
	DemoGitSensorFileName             string
	DemoGitSensorObjectName           string
	DemoGitEventName                  string
	DemoGitTriggerTemplateName        string
	DemoGitDependencyName             string
	GitTokenSecretObjectName          string
	GitTokenSecretKey                 string
	GithubEventTypeHeader             string
	GitlabEventTypeHeader             string
	ArgoCD                            string
	Silent                            bool
	InsecureIngressHost               bool
	BypassIngressClassCheck           bool
	SetDefaultResources               bool
	MinimumMemorySizeRequired         string
	MinimumCpuRequired                string
	MinimumLocalDiskSizeRequired      string
	ReplicaSetResourceName            string
	AnalysisRunResourceName           string
	WorkflowResourceName              string
	RequirementsLink                  string
	DownloadCliLink                   string
	RolloutReporterName               string
	RolloutResourceName               string
	RolloutReporterServiceAccount     string
	SegmentWriteKey                   string
	DefaultNamespace                  string
	NetworkTesterName                 string
	NetworkTesterGenerateName         string
	NetworkTesterImage                string
	MinKubeVersion                    string
	MaxKubeVersion                    string
	MasterIngressName                 string
	InClusterPath                     string
	SccName                           string
	CFInternalGitSources              []string
	CFInternalReporters               []string
	GsCreateFlow                      string
	InCluster                         string
	IsDownloadRuntimeLogs             bool
	IngressHost                       string
	IscRuntimesDir                    string
}

// Get returns the global store
func Get() *Store {
	return &s
}

func init() {
	s.AddClusterJobName = "csdp-add-cluster-job-"
	s.ArgoCDServerName = "argocd-server"
	s.ArgoCDTokenKey = "token"
	s.ArgoCDTokenSecret = "argocd-token"
	s.ArgoWFServiceName = "argo-server"
	s.ArgoWFServicePort = 2746
	s.BinaryName = binaryName
	s.Codefresh = "codefresh"
	s.GitSourceName = "default-git-source"
	s.CFComponentType = "component"
	s.CFGitSourceType = "git-source"
	s.CFRuntimeDefType = "runtimeDef"
	s.CFRuntimeType = "runtime"
	s.CFTokenSecret = "codefresh-token"
	s.CFTokenSecretKey = "token"
	s.CFStoreIVSecretKey = "encryptionIV"
	s.CodefreshCM = "codefresh-cm"
	s.CodefreshSA = "codefresh-sa"
	s.ComponentsReporterName = "components-reporter"
	s.ComponentsReporterSA = "components-reporter-sa"
	s.DefaultAPI = "https://g.codefresh.io"
	s.EventBusName = "codefresh-eventbus"
	s.EventReportingEndpoint = "/2.0/api/events"
	s.EventsReporterName = "events-reporter"
	s.WorkflowsIngressName = "-workflows-ingress"
	s.WorkflowsIngressPath = "workflows"
	s.AppProxyIngressName = "-cap-app-proxy"
	s.AppProxyIngressPath = "/app-proxy"
	s.AppProxyServicePort = 3017
	s.AppProxyServiceName = "cap-app-proxy"
	s.DocsLink = "https://codefresh.io/csdp-docs/"
	s.LabelKeyCFType = "codefresh.io/entity"
	s.LabelKeyCFInternal = "codefresh.io/internal"
	s.AnnotationKeySyncWave = "argocd.argoproj.io/sync-wave"
	s.MaxDefVersion = semver.MustParse(maxDefVersion)
	s.RuntimeDefURL = RuntimeDefURL
	s.MarketplaceGitSourceName = "marketplace-git-source"
	s.MarketplaceRepo = "https://github.com/codefresh-io/argo-hub.git"
	s.WaitTimeout = 8 * time.Minute
	s.WorkflowName = "workflow"
	s.WorkflowReporterName = "workflow-reporter"
	s.WorkflowTriggerServiceAccount = "argo"
	s.DemoCalendarEventSourceFileName = "calendar.event-source.yaml"
	s.DemoCalendarSensorObjectName = "calendar"
	s.DemoCalendarSensorFileName = "calendar.sensor.yaml"
	s.DemoCalendarEventSourceObjectName = "calendar"
	s.DemoCalendarEventName = "example-with-interval"
	s.DemoCalendarDependencyName = "calendar-dep"
	s.DemoWorkflowTemplateFileName = "echo-message.workflow-template.yaml"
	s.DemoWorkflowTemplateName = "echo-message"
	s.DemoGitEventSourceFileName = "push-commit.event-source.yaml"
	s.DemoGitEventSourceObjectName = "push-commit"
	s.WebhooksRootPath = "/webhooks"
	s.DemoGitEventSourceTargetPort = "80"
	s.DemoGitEventSourceServicePort = 80
	s.DemoPipelinesIngressObjectName = fmt.Sprintf("%s-%s", "cdp", s.GitSourceName)
	s.DemoPipelinesIngressFileName = fmt.Sprintf("%s.ingress.yaml", s.DemoPipelinesIngressObjectName)
	s.DemoGitSensorFileName = "push-commit.sensor.yaml"
	s.DemoGitSensorObjectName = "push-commit"
	s.DemoGitEventName = "push-commit"
	s.DemoGitTriggerTemplateName = "push-commit"
	s.DemoGitDependencyName = "push-commit"
	s.GitTokenSecretObjectName = "autopilot-secret"
	s.GitTokenSecretKey = "git_token"
	s.GithubEventTypeHeader = "X-GitHub-Event"
	s.GitlabEventTypeHeader = "X-Gitlab-Event"
	s.ArgoCD = "argo-cd"
	s.RolloutResourceName = "rollouts"
	s.ReplicaSetResourceName = "replicasets"
	s.AnalysisRunResourceName = "analysisruns"
	s.MinimumMemorySizeRequired = "5000"
	s.MinimumCpuRequired = "2"
	s.WorkflowResourceName = "workflows"
	s.RolloutReporterName = "rollout-reporter"
	s.RolloutReporterServiceAccount = "rollout-reporter-sa"
	s.SegmentWriteKey = segmentWriteKey
	s.RequirementsLink = "https://codefresh.io/csdp-docs/docs/runtime/requirements/"
	s.DownloadCliLink = "https://codefresh.io/csdp-docs/docs/clients/csdp-cli/"
	s.DefaultNamespace = "default"
	s.NetworkTesterName = "cf-network-tester"
	s.NetworkTesterGenerateName = "cf-network-tester-"
	s.NetworkTesterImage = "quay.io/codefresh/cf-venona-network-tester:latest"
	s.MinKubeVersion = "v1.18.0"
	s.MaxKubeVersion = "v1.25.0"
	s.MasterIngressName = "-master"
	s.InClusterPath = "/bootstrap/cluster-resources/in-cluster"
	s.SccName = "cf-scc"
	s.CFInternalGitSources = []string{s.MarketplaceGitSourceName}
	s.CFInternalReporters = []string{s.EventsReporterName, s.WorkflowReporterName, s.RolloutReporterName}
	s.InCluster = "https://kubernetes.default.svc"
	s.IscRuntimesDir = "runtimes"

	initVersion()
}

func initVersion() {
	s.Version.Version = semver.MustParse(version)
	s.Version.BuildDate = buildDate
	s.Version.GitCommit = gitCommit
	s.Version.GoVersion = runtime.Version()
	s.Version.GoCompiler = runtime.Compiler
	s.Version.Platform = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
}
