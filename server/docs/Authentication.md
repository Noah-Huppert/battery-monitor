# Authentication
Clients must authenticate with the HTTP API to make requests.

# Table Of Contents
- [Overview](#overview)
- [Auth Tokens](#auth-tokens)
- [Providing Auth Tokens](#providing-auth-tokens)
- [Retrieving Auth Tokens](#retrieving-auth-tokens)

# Overview
The HTTP API uses Auth Tokens, in the form of JSON Web Tokens (JWTs), to 
authenticate clients. These tokens hold data signed by a private key, 
held by the server.  

They must be included in most HTTP API request via the `Authorization` HTTP 
Header.  

Auth Tokens can be obtained by exchanging a valid 
[Firebase Authentication](https://firebase.google.com/docs/auth/) ID Token at 
the [New Auth Token Endpoint](/server/docs/endpoints/auth/tokens.md#create).

# Auth Tokens
The HTTP API uses Auth Tokens for authentication. These auth tokens use the 
JWT format.  

JWTs have 2 main parts, the header and the payload. When transmitted over the 
wire, JWTs take on a more compact string form.

## Header
The JWT header holds information about the type of JWT and the algorithm used 
to sign the data.  

The type of all JWTs issued by the HTTP API will be `JWT`. This information 
will be stored in the `typ` field of the JWT header.  

The algorithm used to sign JWTs will be `HS256`. This will be stored in the 
`alg` field of the JWT header.  

## Payload
A JWT holds a set of claims. JWTs issued by the HTTP API will have the 
following claims:

- `iss` (String): Issuer of token, will always be 
                  `github.com/Noah-Huppert/battery-monitor/server`
- `sub` (String): Who JWT refers to. Will be the ID of the device followed by 
		  the ID of the user the auth token was obtained for. Separated 
		  by a colon
- `aud` (String): Who can receive the token, will always be 
                  `github.com/Noah-Huppert/battery-monitor/server`
- `iat` (Date Time): When token was issued
- `jti` (String): ID of token, used revoke specific tokens

# Providing Auth Tokens
Auth tokens must be provided in all requests, with a few exceptions.  

By placing the auth token in string form in the `Authorization` HTTP header.

# Retrieving Auth Tokens
To authenticate with the HTTP API you need an auth token. To retrieve an auth 
token you must login.  

The HTTP API uses Firebase Authentication to verify user's identity. First 
obtain an ID token from the Firebase Authentication platform. Then send this 
ID token to the [New Auth Token Endpoint](/server/docs/endpoints/auth/tokens.md#create).  

This endpoint will return a HTTP API auth token if the provided ID token is 
valid.  

