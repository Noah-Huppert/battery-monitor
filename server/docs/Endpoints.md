# Endpoints
Endpoints are how clients invoke API actions.

# Table Of Contents
- [Overview](#overview)
- [Requests](#requests)
- [Responses](#responses)
	- [Errors](#errors)

# Endpoints
The HTTP API is accessed completely via endpoints.  

Endpoints are categorized by the model they perform actions on. And can be 
accessed at the URL: `/api/v1/<plural model name>`.  

The following model's endpoints are described in more detail on separate pages: 

- [Users Endpoints](/server/docs/endpoints/Users.md)
- [Auth Token Revocations Endpoints](/server/docs/endpoints/Auth-Token-Revocations.md)
- [Devices Endpoints](/server/docs/endpoints/Devices.md)
- [Battery Readings Endpoints](/server/docs/endpoints/Battery-Readings.md)  

# Requests
Requests to endpoints must be made using the HTTPS protocol. Which leverages 
TLS encryption.  

This is needed because sensitive information is sent to and from the API.  

Additionally an auth token must be provide in the `Authorization` HTTP header. 
See the [Authentication Page](/server/docs/Authentication.md#providing-auth-tokens) 
for more details.  

Requests must be made to the `/api/v1/` base path. 

# Responses
Responses will always be in the JSON format. And contain an `errors` field at 
a minimum.  

A response can be considered a success if its HTTP status code equals 200.  
If the HTTP status code is not 200 refer to the `errors` field.  

## Errors
The `errors` field provides details about failed HTTP API requests. It will 
always be returned in every response. Regardless of if there is an error or 
not.  

It is an array of special error objects. These special error objects contain 
the following fields:

- `type` (String): String designating the type of error. Useful for client 
		   side error handling.
	- Valid error types are:
		- TBD
- `technical_msg` (String): Message containing technical details of error
- `user_msg` (String): Message which should be displayed to the user

If the HTTP status code is not 200, and the `errors` array is empty, then 
something went very wrong in handling your HTTP API requests. And you should 
[Open a ticket](https://github.com/Noah-Huppert/battery-monitor/issues/new).
