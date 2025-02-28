package contract

type HealthCheckResponse struct {
	Resource string `json:"resource,omitempty"`
	Status   string `json:"status,omitempty"`
}
