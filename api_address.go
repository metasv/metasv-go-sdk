/*
 * MetaSV API Spec
 *
 * API definition for MetaSV provided apis
 *
 * API version: 2.1.4
 * Contact: heqiming@metasv.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metasv

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// AddressApiService AddressApi service
type AddressApiService service

type ApiAddressAddressBalanceGetRequest struct {
	ctx        _context.Context
	ApiService *AddressApiService
	address    string
}

func (r ApiAddressAddressBalanceGetRequest) Execute() (AddressBalance, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.AddressAddressBalanceGetExecute(r)
}

/*
 * AddressAddressBalanceGet Get address balance by specific address.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param address the requested address
 * @return ApiAddressAddressBalanceGetRequest
 */
func (a *AddressApiService) AddressAddressBalanceGet(ctx _context.Context, address string) ApiAddressAddressBalanceGetRequest {
	return ApiAddressAddressBalanceGetRequest{
		ApiService: a,
		ctx:        ctx,
		address:    address,
	}
}

/*
 * Execute executes the request
 * @return AddressBalance
 */
func (a *AddressApiService) AddressAddressBalanceGetExecute(r ApiAddressAddressBalanceGetRequest) (AddressBalance, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  AddressBalance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressApiService.AddressAddressBalanceGet")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/address/{address}/balance"
	localVarPath = strings.Replace(localVarPath, "{"+"address"+"}", _neturl.PathEscape(parameterToString(r.address, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, executionError
}

type ApiAddressAddressUtxoGetRequest struct {
	ctx        _context.Context
	ApiService *AddressApiService
	address    string
	flag       *string
}

func (r ApiAddressAddressUtxoGetRequest) Flag(flag string) ApiAddressAddressUtxoGetRequest {
	r.flag = &flag
	return r
}

func (r ApiAddressAddressUtxoGetRequest) Execute() ([]AddressUtxo, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.AddressAddressUtxoGetExecute(r)
}

/*
 * AddressAddressUtxoGet Get address utxos by specific address(100 per page).
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param address the requested address
 * @return ApiAddressAddressUtxoGetRequest
 */
func (a *AddressApiService) AddressAddressUtxoGet(ctx _context.Context, address string) ApiAddressAddressUtxoGetRequest {
	return ApiAddressAddressUtxoGetRequest{
		ApiService: a,
		ctx:        ctx,
		address:    address,
	}
}

/*
 * Execute executes the request
 * @return []AddressUtxo
 */
func (a *AddressApiService) AddressAddressUtxoGetExecute(r ApiAddressAddressUtxoGetRequest) ([]AddressUtxo, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  []AddressUtxo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressApiService.AddressAddressUtxoGet")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/address/{address}/utxo"
	localVarPath = strings.Replace(localVarPath, "{"+"address"+"}", _neturl.PathEscape(parameterToString(r.address, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.flag != nil {
		localVarQueryParams.Add("flag", parameterToString(*r.flag, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, executionError
}
