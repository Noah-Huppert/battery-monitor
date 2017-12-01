# Auth Token
The auth token model is used to record the details of an issued auth token.  

# Table Of Contents
- [Schema](#schema)

# Schema
The auth token model contains the following fields:

- `user_id` (uint): The ID of the user to which the auth token authenticates
- `device_id` (uint): The ID of the device which the auth token authenticates
- `issued_at` (Time): Date and time authentication token was issued

# Relationships
The user the auth token belongs to is specified via the `user_id` field. 
Creating a One to One relationship.  

The device the auth token belongs to is specified via the `device_id` field. 
Creating a One to One relationship.
