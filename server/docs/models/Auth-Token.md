# Auth Token
The auth token model is used to record the details of an issued auth token.  

# Table Of Contents
- [Schema](#schema)

# Schema
The auth token model contains the following fields:

- `user_id` (uint): The ID of the user to which the auth token authenticates
- `device_id` (uint): The ID of the device which the auth token authenticates
- `revoked` (Boolean): Specifies if the auth token is revoked, defaults to 
		       false, true if the auth token is revoked

*Note: The time the token was issued at, is recorded in the [created at field](/server/docs/Models.md#base-model).*

# Relationships
The user the auth token belongs to is specified via the `user_id` field. 
Creating a One to One relationship.  

The device the auth token belongs to is specified via the `device_id` field. 
Creating a One to One relationship.
