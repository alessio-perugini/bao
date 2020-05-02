package bao

import "encoding/json"

func UnmarshalAbuseIPAPI(data []byte) (AbuseIPAPI, error) {
	var r AbuseIPAPI
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AbuseIPAPI) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AbuseIPAPI struct {
	Data Data `json:"data"`
}

type Data struct {
	IPAddress            string   `json:"ipAddress"`
	IsPublic             bool     `json:"isPublic"`
	IPVersion            int64    `json:"ipVersion"`
	IsWhitelisted        bool     `json:"isWhitelisted"`
	AbuseConfidenceScore int64    `json:"abuseConfidenceScore"`
	CountryCode          string   `json:"countryCode"`
	CountryName          string   `json:"countryName"`
	UsageType            string   `json:"usageType"`
	ISP                  string   `json:"isp"`
	Domain               string   `json:"domain"`
	TotalReports         int64    `json:"totalReports"`
	NumDistinctUsers     int64    `json:"numDistinctUsers"`
	LastReportedAt       string   `json:"lastReportedAt"`
	Reports              []Report `json:"reports"`
}

type Report struct {
	ReportedAt          string  `json:"reportedAt"`
	Comment             string  `json:"comment"`
	Categories          []int64 `json:"categories"`
	ReporterID          int64   `json:"reporterId"`
	ReporterCountryCode string  `json:"reporterCountryCode"`
	ReporterCountryName string  `json:"reporterCountryName"`
}
