package utils

import (
	"encoding/json"
	"encoding/xml"
)

type MinioError struct {
	XMLName    xml.Name `xml:"Error"`
	Text       string   `xml:",chardata"`
	Code       string   `xml:"Code"`
	Message    string   `xml:"Message"`
	Key        string   `xml:"Key"`
	BucketName string   `xml:"BucketName"`
	Resource   string   `xml:"Resource"`
	RequestId  string   `xml:"RequestId"`
	HostId     string   `xml:"HostId"`
}

type ScanSbomRequest struct {
	SbomParameters
	SbomBody
}

type SbomParameters struct {
	ImageName             string `json:"image_name"`
	ImageId               string `json:"image_id"`
	ScanId                string `json:"scan_id" required:"true"`
	KubernetesClusterName string `json:"kubernetes_cluster_name"`
	HostName              string `json:"host_name"`
	NodeId                string `json:"node_id"`
	NodeType              string `json:"node_type"`
	ScanType              string `json:"scan_type"`
	ContainerName         string `json:"container_name"`
	SBOMFilePath          string `json:"sbom_file_path"`
	Mode                  string `json:"mode,omitempty"`
	RegistryId            string `json:"registry_id,omitempty"`
	SkipScan              bool   `json:"skip_scan,omitempty"`
}

type SbomBody struct {
	SBOM string `json:"sbom" required:"true"`
}

type SecretScanParameters struct {
	ImageName             string `json:"image_name"`
	ImageId               string `json:"image_id"`
	ScanId                string `json:"scan_id" required:"true"`
	KubernetesClusterName string `json:"kubernetes_cluster_name"`
	HostName              string `json:"host_name"`
	NodeId                string `json:"node_id"`
	NodeType              string `json:"node_type"`
	ScanType              string `json:"scan_type"`
	ContainerName         string `json:"container_name"`
	Mode                  string `json:"mode,omitempty"`
	RegistryId            string `json:"registry_id,omitempty"`
}

type MalwareScanParameters struct {
	ImageName             string `json:"image_name"`
	ImageId               string `json:"image_id"`
	ScanId                string `json:"scan_id" required:"true"`
	KubernetesClusterName string `json:"kubernetes_cluster_name"`
	HostName              string `json:"host_name"`
	NodeId                string `json:"node_id"`
	NodeType              string `json:"node_type"`
	ScanType              string `json:"scan_type"`
	ContainerName         string `json:"container_name"`
	Mode                  string `json:"mode,omitempty"`
	RegistryId            string `json:"registry_id,omitempty"`
}

type ReportParams struct {
	ReportID   string        `json:"report_id"`
	ReportType string        `json:"report_type"`
	Duration   int           `json:"duration"`
	Filters    ReportFilters `json:"filters"`
}

type ReportFilters struct {
	ScanId                string                `json:"scan_id"`
	ScanType              string                `json:"scan_type" validate:"required" required:"true" enum:"vulnerability,secret,malware,compliance,cloud_compliance"`
	NodeType              string                `json:"node_type" validate:"required" required:"true" enum:"host,container,container_image,linux,cluster,aws,gcp,azure"`
	SeverityOrCheckType   []string              `json:"severity_or_check_type" enum:"critical,high,medium,low,cis,gdpr,nist,hipaa,pci,soc_2"`
	IncludeDeadNode       bool                  `json:"include_dead_nodes"`
	AdvancedReportFilters AdvancedReportFilters `json:"advanced_report_filters,omitempty"`
}

type RegistrySyncParams struct {
	PgID int32 `json:"pg_id"`
}

type AdvancedReportFilters struct {
	Masked                []bool   `json:"masked,omitempty"`
	ScanStatus            []string `json:"scan_status,omitempty"`
	PodName               []string `json:"pod_name,omitempty"`
	ContainerName         []string `json:"container_name,omitempty"`
	ImageName             []string `json:"image_name,omitempty"`
	HostName              []string `json:"host_name,omitempty"`
	AccountId             []string `json:"account_id,omitempty"`
	KubernetesClusterName []string `json:"kubernetes_cluster_name,omitempty"`
}

func (r ReportFilters) String() string {
	if b, err := json.Marshal(r); err != nil {
		return ""
	} else {
		return string(b)
	}
}
