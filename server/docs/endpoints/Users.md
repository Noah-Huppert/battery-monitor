# Users
The [Users](/server/docs/models/User.md) endpoint manages user resources.

# Table Of Contents
- [Overview](#overview)
- [Create](#create)
- [Get](#get)
- [Update](#update)
- [Delete](#delete)

# Overview
The [Users](/server/docs/models/User.md) endpoint is located at the 
`/api/v1/users/` URL.

# Create
The create endpoint makes a new user resource.  

It is located at the `/api/v1/users/` URL.

## Request
### Body
Required body fields:

- `first_name` (String): The new user's given name
- `last_name` (String): The new user's family name
- `email` (String): The new user's email address

## Response
### Body
The newly created [User](/server/docs/models/User.md#schema) will be returned 
in the `user` body field.

## Errors
No additional errors other than the 
[Common Error Types](/server/docs/Errors.md#common-error-types).

# Get
The get endpoint retrieves either the current user, or the specified user.  

It is located at the `/api/v1/users/{id}/` URL.

If the `id` parameter is not provided, the currently authenticated user will 
be retrieved.

## Request
### Parameters
Optional parameters:

- `id` (uint): The ID of the user resource to retrieve

## Response
### Body
If the `id` parameter is provided, the specified user will be returned in the 
`user` body field.  

If the `id` parameter is not provided, the currently authenticated user will be 
returned in the `body` field.

### Errors
- `NotFound`: If the specified user resource is not found

# Update
The update endpoint changes user resource values.  

It is located at the `/api/v1/users/{id}/` URL.

## Request
### Parameters
Required parameters:

- `id` (uint): The ID of the user resource to update

### Body
One or more of the following body fields must be provided:

- `first_name` (String): New first name of the user
- `last_name` (String): New last name of the user
- `email` (String): New email of the user

## Response
### Body
The new [User](/server/docs/model/User.md#schema) will be returned in the 
`user` body field.

### Errors
- `NotFound`: If the specified user resource is not found

# Delete
The delete endpoint destroys a specified user resource.  

It is located at the `/api/v1/users/{id}/` URL.  

All resources owned by the user will also be deleted.  

## Request
### Parameters
Required parameters:

- `id` (uint): The ID of the user resource to delete

## Response
### Body
Not additional fields will be returned.
