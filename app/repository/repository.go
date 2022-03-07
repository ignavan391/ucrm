package repository

import "github.com/ignavan39/ucrm-go/app/models"

type DashboardRepository interface {
	AddDashboard(name string, userId string) (*models.Dashboard, error)
	GetOneDashboard(dashboardId string) (*models.Dashboard, error)
	AddUserToDashboard(dashboardId string, userId string, access string) (*string, error)
	GetOneDashboardWithUserAccess(dashboardId string, userId string, accessType string) (*models.Dashboard, error)
}

type UserRepository interface {
	GetOneUserByEmail(email string, password string) (*models.User, error)
	AddUser(email string, password string) (*models.User, error)
}

type PipelineRepository interface {
	AddPipeline(name string, dashboardId string,order int) (*models.Pipeline, error)
	GetOnePipeline(pipelineId string) (*models.Pipeline, error)
}
