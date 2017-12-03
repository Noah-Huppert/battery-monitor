# Auth Tokens
The auth tokens endpoints create and manage access to the HTTP API.

# Table Of Contents
- [Overview](#overview)
- [Create](#create)
- [Get](#get)
- [Delete](#delete)

# Overview
The auth token endpoints are available at the `/api/v1/auth_tokens/` URL.  

# Create
The create endpoint issues a new auth token.  

## Request
Requires authentication.  

### Body
The following body fields should be provided in a create token request:

- `id_token` (String): Firebase Authentication ID token used as proof of 
		       authentication
- `device_fingerprint` (String): Fingerprint of device used to login on

The `id_token` can be provided to the Firebase Authentication API to determine 
who is logging in.  

The `device_fingerprint` is used to specify which of the user's devices is 
making the login request. An auth token will be issue specifically for this 
device.  

If the user does not have a device with the specified fingerprint, one will be 
created for them. 

## Response
### Body

- [Auth Token Model Fields](/server/docs/models/Auth-Token.md#schema)
- `value` (String): The actual auth token value. This will only be provided 
		    once. And should be provided in most HTTP API requests via 
		    the `Authentication` header. See the 
		    [Authentication Page](/server/docs/Authentication.md) for 
		    more details.

### Errors
- `NotAuthenticated`: If the provided `id_token` is not valid

# Get
The get endpoint returns either a list or a single auth token.  

Located at the url `/api/v1/auth_tokens/{id}`.  

If the `id` parameter is left blank, a list of all the current user's auth 
tokens is returned.  

## Request
Requires authentication.  

### URL Parameters
Optional fields: 

- `id` (uint): The ID of the token to retrieve

## Response
### Body
If the `id` parameter was provided the requested token will be returned in the 
`token` field.  

If the `id` parameter was **not** provided, a list of the current user's auth 
tokens will be returned in the `tokens` field.  

### Errors
- `NotFound`

# Delete
The delete endpoint destroys the specified auth token.

Located at the URL: `/api/v1/auth_tokens/{id}`

## Request
Authentication required.

### Parameters
Required parameters:

- `id` (uint): The ID of the auth token to delete

## Response
### Body
No extra fields will be returned.

### Errors
- `NotFound`
