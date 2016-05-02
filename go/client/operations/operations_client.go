package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateHookBySiteID create hook by site Id API
*/
func (a *Client) CreateHookBySiteID(params *CreateHookBySiteIDParams, authInfo runtime.ClientAuthInfoWriter) (*CreateHookBySiteIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateHookBySiteIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createHookBySiteId",
		Method:             "POST",
		PathPattern:        "/hooks",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateHookBySiteIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateHookBySiteIDCreated), nil
}

/*
CreateSite create site API
*/
func (a *Client) CreateSite(params *CreateSiteParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSiteCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSiteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSite",
		Method:             "POST",
		PathPattern:        "/sites",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSiteReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSiteCreated), nil
}

/*
CreateSiteSnippet create site snippet API
*/
func (a *Client) CreateSiteSnippet(params *CreateSiteSnippetParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSiteSnippetCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSiteSnippetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSiteSnippet",
		Method:             "POST",
		PathPattern:        "/sites/{site_id}/snippets",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSiteSnippetReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSiteSnippetCreated), nil
}

/*
DeleteHookBySiteID delete hook by site Id API
*/
func (a *Client) DeleteHookBySiteID(params *DeleteHookBySiteIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteHookBySiteIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteHookBySiteIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteHookBySiteId",
		Method:             "DELETE",
		PathPattern:        "/hooks/{hook_id}",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteHookBySiteIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteHookBySiteIDNoContent), nil
}

/*
DeleteSite delete site API
*/
func (a *Client) DeleteSite(params *DeleteSiteParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSiteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSiteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSite",
		Method:             "DELETE",
		PathPattern:        "/sites/{site_id}",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSiteReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSiteOK), nil
}

/*
DeleteSiteSnippet delete site snippet API
*/
func (a *Client) DeleteSiteSnippet(params *DeleteSiteSnippetParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSiteSnippetNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSiteSnippetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSiteSnippet",
		Method:             "DELETE",
		PathPattern:        "/sites/{site_id}/snippets/{snippet_id}",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSiteSnippetReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSiteSnippetNoContent), nil
}

/*
GetSite get site API
*/
func (a *Client) GetSite(params *GetSiteParams, authInfo runtime.ClientAuthInfoWriter) (*GetSiteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSiteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSite",
		Method:             "GET",
		PathPattern:        "/sites/{site_id}",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSiteReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSiteOK), nil
}

/*
GetSiteDeploy get site deploy API
*/
func (a *Client) GetSiteDeploy(params *GetSiteDeployParams, authInfo runtime.ClientAuthInfoWriter) (*GetSiteDeployOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSiteDeployParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSiteDeploy",
		Method:             "GET",
		PathPattern:        "/sites/{site_id}/deploys/{deploy_id}",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSiteDeployReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSiteDeployOK), nil
}

/*
GetSiteFileByPathName get site file by path name API
*/
func (a *Client) GetSiteFileByPathName(params *GetSiteFileByPathNameParams, authInfo runtime.ClientAuthInfoWriter) (*GetSiteFileByPathNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSiteFileByPathNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSiteFileByPathName",
		Method:             "GET",
		PathPattern:        "/sites/{site_id}/files/{file_path}",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSiteFileByPathNameReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSiteFileByPathNameOK), nil
}

/*
GetSiteMetadata get site metadata API
*/
func (a *Client) GetSiteMetadata(params *GetSiteMetadataParams, authInfo runtime.ClientAuthInfoWriter) (*GetSiteMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSiteMetadataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSiteMetadata",
		Method:             "GET",
		PathPattern:        "/sites/{site_id}/metadata",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSiteMetadataReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSiteMetadataOK), nil
}

/*
GetSiteSnippet get site snippet API
*/
func (a *Client) GetSiteSnippet(params *GetSiteSnippetParams, authInfo runtime.ClientAuthInfoWriter) (*GetSiteSnippetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSiteSnippetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSiteSnippet",
		Method:             "GET",
		PathPattern:        "/sites/{site_id}/snippets/{snippet_id}",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSiteSnippetReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSiteSnippetOK), nil
}

/*
ListFormSubmissions list form submissions API
*/
func (a *Client) ListFormSubmissions(params *ListFormSubmissionsParams, authInfo runtime.ClientAuthInfoWriter) (*ListFormSubmissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListFormSubmissionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listFormSubmissions",
		Method:             "GET",
		PathPattern:        "/forms/{form_id}/submissions",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListFormSubmissionsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListFormSubmissionsOK), nil
}

/*
ListForms list forms API
*/
func (a *Client) ListForms(params *ListFormsParams, authInfo runtime.ClientAuthInfoWriter) (*ListFormsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListFormsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listForms",
		Method:             "GET",
		PathPattern:        "/forms",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListFormsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListFormsOK), nil
}

/*
ListHookTypes list hook types API
*/
func (a *Client) ListHookTypes(params *ListHookTypesParams, authInfo runtime.ClientAuthInfoWriter) (*ListHookTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListHookTypesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listHookTypes",
		Method:             "GET",
		PathPattern:        "/hooks/types",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListHookTypesReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListHookTypesOK), nil
}

/*
ListHooksBySiteID list hooks by site Id API
*/
func (a *Client) ListHooksBySiteID(params *ListHooksBySiteIDParams, authInfo runtime.ClientAuthInfoWriter) (*ListHooksBySiteIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListHooksBySiteIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listHooksBySiteId",
		Method:             "GET",
		PathPattern:        "/hooks",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListHooksBySiteIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListHooksBySiteIDOK), nil
}

/*
ListSiteDeploys list site deploys API
*/
func (a *Client) ListSiteDeploys(params *ListSiteDeploysParams, authInfo runtime.ClientAuthInfoWriter) (*ListSiteDeploysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSiteDeploysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSiteDeploys",
		Method:             "GET",
		PathPattern:        "/sites/{site_id}/deploys",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSiteDeploysReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSiteDeploysOK), nil
}

/*
ListSiteFiles list site files API
*/
func (a *Client) ListSiteFiles(params *ListSiteFilesParams, authInfo runtime.ClientAuthInfoWriter) (*ListSiteFilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSiteFilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSiteFiles",
		Method:             "GET",
		PathPattern:        "/sites/{site_id}/files",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSiteFilesReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSiteFilesOK), nil
}

/*
ListSiteForms list site forms API
*/
func (a *Client) ListSiteForms(params *ListSiteFormsParams, authInfo runtime.ClientAuthInfoWriter) (*ListSiteFormsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSiteFormsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSiteForms",
		Method:             "GET",
		PathPattern:        "/sites/{site_id}/forms",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSiteFormsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSiteFormsOK), nil
}

/*
ListSiteSnippets list site snippets API
*/
func (a *Client) ListSiteSnippets(params *ListSiteSnippetsParams, authInfo runtime.ClientAuthInfoWriter) (*ListSiteSnippetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSiteSnippetsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSiteSnippets",
		Method:             "GET",
		PathPattern:        "/sites/{site_id}/snippets",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSiteSnippetsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSiteSnippetsOK), nil
}

/*
ListSiteSubmissions list site submissions API
*/
func (a *Client) ListSiteSubmissions(params *ListSiteSubmissionsParams, authInfo runtime.ClientAuthInfoWriter) (*ListSiteSubmissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSiteSubmissionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSiteSubmissions",
		Method:             "GET",
		PathPattern:        "/sites/{site_id}/submissions",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSiteSubmissionsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSiteSubmissionsOK), nil
}

/*
ListSites list sites API
*/
func (a *Client) ListSites(params *ListSitesParams, authInfo runtime.ClientAuthInfoWriter) (*ListSitesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSitesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSites",
		Method:             "GET",
		PathPattern:        "/sites",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSitesReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSitesOK), nil
}

/*
ProvisionSiteSSL provision site s s l API
*/
func (a *Client) ProvisionSiteSSL(params *ProvisionSiteSSLParams, authInfo runtime.ClientAuthInfoWriter) (*ProvisionSiteSSLOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProvisionSiteSSLParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "provisionSiteSSL",
		Method:             "POST",
		PathPattern:        "/sites/{site_id}/ssl",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProvisionSiteSSLReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ProvisionSiteSSLOK), nil
}

/*
RestoreSiteDeploy restore site deploy API
*/
func (a *Client) RestoreSiteDeploy(params *RestoreSiteDeployParams, authInfo runtime.ClientAuthInfoWriter) (*RestoreSiteDeployCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestoreSiteDeployParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "restoreSiteDeploy",
		Method:             "POST",
		PathPattern:        "/sites/{site_id}/deploys/{deploy_id}/restore",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RestoreSiteDeployReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RestoreSiteDeployCreated), nil
}

/*
UpdateSite update site API
*/
func (a *Client) UpdateSite(params *UpdateSiteParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSiteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSiteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSite",
		Method:             "PATCH",
		PathPattern:        "/sites/{site_id}",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSiteReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateSiteNoContent), nil
}

/*
UpdateSiteMetadata update site metadata API
*/
func (a *Client) UpdateSiteMetadata(params *UpdateSiteMetadataParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSiteMetadataNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSiteMetadataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSiteMetadata",
		Method:             "PUT",
		PathPattern:        "/sites/{site_id}/metadata",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSiteMetadataReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateSiteMetadataNoContent), nil
}

/*
UpdateSiteSnippet update site snippet API
*/
func (a *Client) UpdateSiteSnippet(params *UpdateSiteSnippetParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSiteSnippetNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSiteSnippetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSiteSnippet",
		Method:             "PUT",
		PathPattern:        "/sites/{site_id}/snippets/{snippet_id}",
		ProducesMediaTypes: []string{"application/io.swagger.netlify.v1+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSiteSnippetReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateSiteSnippetNoContent), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
