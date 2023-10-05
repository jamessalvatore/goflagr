# \ExportApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExportEvalCacheJSON**](ExportApi.md#GetExportEvalCacheJSON) | **Get** /export/eval_cache/json | 
[**GetExportSqlite**](ExportApi.md#GetExportSqlite) | **Get** /export/sqlite | 


# **GetExportEvalCacheJSON**
> interface{} GetExportEvalCacheJSON(ctx, )


Export JSON format of the eval cache dump

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExportSqlite**
> *os.File GetExportSqlite(ctx, optional)


Export sqlite3 format of the db dump, which is converted from the main database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetExportSqliteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetExportSqliteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **excludeSnapshots** | **optional.Bool**| export without snapshots data - useful for smaller db without snapshots  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

