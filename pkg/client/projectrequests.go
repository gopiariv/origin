package client

import (
	kapi "github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/fields"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/labels"

	projectapi "github.com/openshift/origin/pkg/project/api"
	_ "github.com/openshift/origin/pkg/user/api/v1beta1"
)

// UsersInterface has methods to work with User resources in a namespace
type ProjectRequestsInterface interface {
	ProjectRequests() ProjectRequestInterface
}

// UserInterface exposes methods on user resources.
type ProjectRequestInterface interface {
	Create(p *projectapi.ProjectRequest) (*projectapi.Project, error)
	List(label labels.Selector, field fields.Selector) (*kapi.Status, error)
}

type projectRequests struct {
	r *Client
}

// newUsers returns a users
func newProjectRequests(c *Client) *projectRequests {
	return &projectRequests{
		r: c,
	}
}

// Create creates a new ProjectRequest
func (c *projectRequests) Create(p *projectapi.ProjectRequest) (result *projectapi.Project, err error) {
	result = &projectapi.Project{}
	err = c.r.Post().Resource("projectrequests").Body(p).Do().Into(result)
	return
}

// List returns a status object indicating that a user can call the Create or an error indicating why not
func (c *projectRequests) List(label labels.Selector, field fields.Selector) (result *kapi.Status, err error) {
	result = &kapi.Status{}
	err = c.r.Get().Resource("projectrequests").LabelsSelectorParam(label).FieldsSelectorParam(field).Do().Into(result)
	return result, err
}
