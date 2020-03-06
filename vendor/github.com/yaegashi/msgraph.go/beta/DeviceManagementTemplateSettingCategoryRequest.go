// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// DeviceManagementTemplateSettingCategoryRequestBuilder is request builder for DeviceManagementTemplateSettingCategory
type DeviceManagementTemplateSettingCategoryRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceManagementTemplateSettingCategoryRequest
func (b *DeviceManagementTemplateSettingCategoryRequestBuilder) Request() *DeviceManagementTemplateSettingCategoryRequest {
	return &DeviceManagementTemplateSettingCategoryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceManagementTemplateSettingCategoryRequest is request for DeviceManagementTemplateSettingCategory
type DeviceManagementTemplateSettingCategoryRequest struct{ BaseRequest }

// Get performs GET request for DeviceManagementTemplateSettingCategory
func (r *DeviceManagementTemplateSettingCategoryRequest) Get(ctx context.Context) (resObj *DeviceManagementTemplateSettingCategory, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceManagementTemplateSettingCategory
func (r *DeviceManagementTemplateSettingCategoryRequest) Update(ctx context.Context, reqObj *DeviceManagementTemplateSettingCategory) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceManagementTemplateSettingCategory
func (r *DeviceManagementTemplateSettingCategoryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RecommendedSettings returns request builder for DeviceManagementSettingInstance collection
func (b *DeviceManagementTemplateSettingCategoryRequestBuilder) RecommendedSettings() *DeviceManagementTemplateSettingCategoryRecommendedSettingsCollectionRequestBuilder {
	bb := &DeviceManagementTemplateSettingCategoryRecommendedSettingsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/recommendedSettings"
	return bb
}

// DeviceManagementTemplateSettingCategoryRecommendedSettingsCollectionRequestBuilder is request builder for DeviceManagementSettingInstance collection
type DeviceManagementTemplateSettingCategoryRecommendedSettingsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceManagementSettingInstance collection
func (b *DeviceManagementTemplateSettingCategoryRecommendedSettingsCollectionRequestBuilder) Request() *DeviceManagementTemplateSettingCategoryRecommendedSettingsCollectionRequest {
	return &DeviceManagementTemplateSettingCategoryRecommendedSettingsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceManagementSettingInstance item
func (b *DeviceManagementTemplateSettingCategoryRecommendedSettingsCollectionRequestBuilder) ID(id string) *DeviceManagementSettingInstanceRequestBuilder {
	bb := &DeviceManagementSettingInstanceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceManagementTemplateSettingCategoryRecommendedSettingsCollectionRequest is request for DeviceManagementSettingInstance collection
type DeviceManagementTemplateSettingCategoryRecommendedSettingsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DeviceManagementSettingInstance collection
func (r *DeviceManagementTemplateSettingCategoryRecommendedSettingsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DeviceManagementSettingInstance, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DeviceManagementSettingInstance
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []DeviceManagementSettingInstance
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for DeviceManagementSettingInstance collection
func (r *DeviceManagementTemplateSettingCategoryRecommendedSettingsCollectionRequest) Get(ctx context.Context) ([]DeviceManagementSettingInstance, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DeviceManagementSettingInstance collection
func (r *DeviceManagementTemplateSettingCategoryRecommendedSettingsCollectionRequest) Add(ctx context.Context, reqObj *DeviceManagementSettingInstance) (resObj *DeviceManagementSettingInstance, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}