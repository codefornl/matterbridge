// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AccessPackageAssignmentRequestObjectCancelRequestParameter undocumented
type AccessPackageAssignmentRequestObjectCancelRequestParameter struct {
}

//
type AccessPackageAssignmentRequestObjectCancelRequestBuilder struct{ BaseRequestBuilder }

// Cancel action undocumented
func (b *AccessPackageAssignmentRequestObjectRequestBuilder) Cancel(reqObj *AccessPackageAssignmentRequestObjectCancelRequestParameter) *AccessPackageAssignmentRequestObjectCancelRequestBuilder {
	bb := &AccessPackageAssignmentRequestObjectCancelRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Cancel"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type AccessPackageAssignmentRequestObjectCancelRequest struct{ BaseRequest }

//
func (b *AccessPackageAssignmentRequestObjectCancelRequestBuilder) Request() *AccessPackageAssignmentRequestObjectCancelRequest {
	return &AccessPackageAssignmentRequestObjectCancelRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *AccessPackageAssignmentRequestObjectCancelRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}