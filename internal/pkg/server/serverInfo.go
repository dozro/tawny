package server

import (
	"net/http"

	"github.com/dozro/tawny/internal/pkg/server_config"
	"github.com/gin-gonic/gin"
)

type serverInfoRespType struct {
	TawnyVersion      string                                `json:"tawny_version" xml:"TawnyVersion"`
	TawnyRevision     string                                `json:"tawny_revision" xml:"TawnyRevision"`
	SourceRepository  string                                `json:"source_repository" yaml:"source_repository" xml:"SourceRepository"`
	GinVersion        string                                `json:"gin_version" yaml:"gin_version" xml:"GinVersion"`
	Operator          ServerInfoOp                          `json:"operator" yaml:"operator" xml:"Operator"`
	DisabledEndpoints server_config.ServerDisabledEndpoints `json:"disabled_endpoints" xml:"DisabledEndpoints"`
}

type ServerInfoOp struct {
	Name             string `json:"name" xml:"Name"`
	Contact          string `json:"contact" xml:"Contact"`
	ImprintURL       string `json:"imprint_url" xml:"ImprintURL"`
	PrivacyPolicyURL string `json:"privacy_policy_url" xml:"PrivacyPolicyURL"`
}

func serverInfo(c *gin.Context) {
	resp := serverInfoRespType{}
	resp.TawnyVersion = proxyConfig.ExtendedServerConfig.TawnyVersion
	resp.TawnyRevision = proxyConfig.ExtendedServerConfig.TawnyRevision
	if proxyConfig.ExtendedServerConfig.DisableGinVersionPublished {
		resp.GinVersion = "hidden"
	} else {
		resp.GinVersion = gin.Version
	}
	resp.SourceRepository = proxyConfig.ExtendedServerConfig.SourceCodeURL
	resp.Operator = ServerInfoOp{
		Name:             proxyConfig.ServerOperator.OperatorName,
		Contact:          proxyConfig.ServerOperator.OperatorContact,
		ImprintURL:       proxyConfig.ServerOperator.ImprintURL,
		PrivacyPolicyURL: proxyConfig.ServerOperator.PrivacyPolicyURL,
	}

	resp.DisabledEndpoints = proxyConfig.DisabledEndpoints
	render(c, http.StatusOK, resp)
}
