package keyvault

import "github.com/Azure/open-service-broker-azure/pkg/service"

type provisioningParameters struct {
	ObjectID string `json:"objectId"`
	ClientID string `json:"clientId"`
}

func (
	s *serviceManager,
) getProvisionParametersSchema() map[string]*service.ParameterSchema {

	p := map[string]*service.ParameterSchema{}

	p["objectId"] = &service.ParameterSchema{
		Type: "string",
		Description: "Object ID for an existing service principal, " +
			"which will be granted access to the new vault.",
		Required: true,
	}

	p["clientId"] = &service.ParameterSchema{
		Type: "string",
		Description: "Client ID (username) for an existing service principal," +
			"which will be granted access to the new vault.",
		Required: true,
	}

	p["clientSecret"] = &service.ParameterSchema{
		Type: "string",
		Description: "Client secret (password) for an existing service " +
			"principal, which will be granted access to the new vault.",
		Required: true,
	}
	return p
}

type secureProvisioningParameters struct {
	ClientSecret string `json:"clientSecret"`
}

type instanceDetails struct {
	ARMDeploymentName string `json:"armDeployment"`
	KeyVaultName      string `json:"keyVaultName"`
	VaultURI          string `json:"vaultUri"`
	ClientID          string `json:"clientId"`
}

type credentials struct {
	VaultURI     string `json:"vaultUri"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

func (s *serviceManager) SplitProvisioningParameters(
	cpp service.CombinedProvisioningParameters,
) (
	service.ProvisioningParameters,
	service.SecureProvisioningParameters,
	error,
) {
	pp := provisioningParameters{}
	if err := service.GetStructFromMap(cpp, &pp); err != nil {
		return nil, nil, err
	}
	spp := secureProvisioningParameters{}
	if err := service.GetStructFromMap(cpp, &spp); err != nil {
		return nil, nil, err
	}
	ppMap, err := service.GetMapFromStruct(pp)
	if err != nil {
		return nil, nil, err
	}
	sppMap, err := service.GetMapFromStruct(spp)
	return ppMap, sppMap, err
}

func (s *serviceManager) SplitBindingParameters(
	params service.CombinedBindingParameters,
) (service.BindingParameters, service.SecureBindingParameters, error) {
	return nil, nil, nil
}
