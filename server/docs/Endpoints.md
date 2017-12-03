# Endpoints
Endpoints are how clients invoke API actions.

# Table Of Contents
- [Overview](#overview)
- [Requests](#requests)
- [Responses](#responses)
	- [Errors](#errors)

# Overview
The HTTP API is accessed completely via endpoints.  

Endpoints are categorized by the model they perform actions on. And can be 
accessed at the URL: `/api/v1/<plural model name>`.  

The following model's endpoints are described in more detail on separate pages: 

- [Users Endpoints](/server/docs/endpoints/Users.md)
- [Auth Token Endpoints](/server/docs/endpoints/Auth-Tokens.md)
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
