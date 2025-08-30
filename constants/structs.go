package constants

type DeploymentRequest struct {
	FullPath   string     `json:"fullPath"`
	ScriptType ScriptType `json:"scriptType"`
}

type DeploymentResponse struct {
	Message string `json:"message"`
	Reason  string `json:"reason,omitempty"`
}
