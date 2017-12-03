# Errors 
The HTTP API returns errors in a common format.

# Table Of Contents
- [Overview](#overview)
- [Common Error Types](#common-error-types)
- [Error Schema](#error-schema)
- [Error Types](#error-types)
	- [General](#general)
	- [Authentication](#authentication)

# Overview
When an HTTP API request fails, one or more errors is returned.  

These errors will be located in the `errors` field of the response. This field 
is an array of error objects.  

The `errors` field is present in every HTTP API response. When a request is 
successful the `errors` field will be an empty array.

# Common Error Types
All HTTP API requests should be expected to return the following common errors:

- `InvalidRequest`
- `InternalError`
- `NotAuthenticated`

Even if the documentation for an endpoint does not specify these errors, they 
should be assumed. This is because these errors could occur for any endpoint. 
So it would be silly to re-write the same errors 30 times.

# Error Schema
Every error follows the same schema:

- `type` (String): String designating the type of error. Useful for client 
		   side error handling.
- `technical_msg` (String): Message containing technical details of error
- `user_msg` (String): Message which should be displayed to the user

# Error Types
The following types of errors exist:

### General
- `InvalidRequest`: Indicates that the API request was not correct in some way
- `InternalError`: An unspecified error occurred, this is only returned if 
		   something went very wrong and the system could not recover 
		   completely
- `NotFound`: The requested resource was not found

### Authentication
- `NotAuthenticated`: Indicates that the provided Auth Token was not valid
