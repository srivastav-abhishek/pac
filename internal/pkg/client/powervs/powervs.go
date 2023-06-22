package powervs

import "github.com/IBM-Cloud/power-go-client/power/models"

type PowerVS interface {
	GetAllInstance() (*models.PVMInstances, error)
	GetAllDHCPServers() (models.DHCPServers, error)
	GetDHCPServer(id string) (*models.DHCPServerDetail, error)
	GetImageByName(name string) (*models.ImageReference, error)
	GetNetworkByName(name string) (*models.NetworkReference, error)
}
