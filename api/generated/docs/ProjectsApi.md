# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectIdGet**](ProjectsApi.md#ProjectIdGet) | **Get** /project/{id} | Get project by UID
[**ProjectIdPatch**](ProjectsApi.md#ProjectIdPatch) | **Patch** /project/{id} | Update project
[**ProjectsGet**](ProjectsApi.md#ProjectsGet) | **Get** /projects | Lists all projects
[**ProjectsPost**](ProjectsApi.md#ProjectsPost) | **Post** /projects | Create project

# **ProjectIdGet**
> ProjectResponse ProjectIdGet(ctx, uid)
Get project by UID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | [**string**](.md)| Unique identifier of . | 

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectIdPatch**
> ProjectResponse ProjectIdPatch(ctx, body, uid)
Update project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateProjectRequest**](UpdateProjectRequest.md)| Project parameters | 
  **uid** | [**string**](.md)| Unique identifier of . | 

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsGet**
> InlineResponse200 ProjectsGet(ctx, optional)
Lists all projects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectsApiProjectsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **length** | **optional.Int32**| Pagination length | 
 **offser** | **optional.Int32**| Pagination offset | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsPost**
> ProjectResponse ProjectsPost(ctx, body)
Create project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CrateProjectRequest**](CrateProjectRequest.md)| Project parameters | 

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

