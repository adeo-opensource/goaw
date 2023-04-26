package awx

import (
	"fmt"
)

// ProjectUpdatesService implements awx projects apis.
type ProjectUpdatesService interface {
	ProjectUpdateCancel(id int) (*ProjectUpdateCancel, error)
	ProjectUpdateGet(id int) (*Job, error)
}

const projectUpdatesAPIEndpoint = "/api/v2/project_updates/"

// ProjectUpdateCancel cancel of awx projects update.
func (p *awx) ProjectUpdateCancel(id int) (*ProjectUpdateCancel, error) {
	result := new(ProjectUpdateCancel)
	endpoint := fmt.Sprintf("%s%d/cancel", projectUpdatesAPIEndpoint, id)
	resp, err := p.client.Requester.GetJSON(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}
	return result, nil
}

// ProjectUpdateGet get of awx projects update.
func (p *awx) ProjectUpdateGet(id int) (*Job, error) {
	result := new(Job)
	endpoint := fmt.Sprintf("%s%d", projectUpdatesAPIEndpoint, id)
	resp, err := p.client.Requester.GetJSON(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}
	return result, nil
}
