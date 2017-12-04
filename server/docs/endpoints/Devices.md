# Devices
The devices endpoint manges device resources.

# Table Of Contents
- [Overview](#overview)
- [Create](#create)
- [Get](#get)
- [Get Battery Readings](#get-battery-readings)
- [Get Auth Tokens](#get-auth-tokens)
- [Update](#update)
- [Delete](#delete)

# Overview
The devices endpoint is located at the `/api/v1/devices/` URL.

# Create
The create endpoint saves a new device resource.  

It is located at the `/api/v1/devices/` URL.  

## Request
### Body
Required fields:

- `name` (String): Name of new device
- `fingerprint` (String): Unique value used to identifying device

## Response
### Body
The newly created [Device](/server/docs/models/Device.md#Schema) in the 
`device` body field.

### Errors
No extra errors other than the 
[Common Error Types](/server/docs/Errors.md#common-error-types).


# Get
The get endpoint retrieves one or more devices.  

It is located at the `/api/v1/devices/{id}/` URL.  

If the `id` parameter is not provided, all the devices the user owns are 
retrieved.

## Request
### Parameters
Optional parameters:

- `id` (uint): The ID of the device resource to retrieve

## Response
### Body
If the `id` parameter is provided, the requested device resource is returned 
via the `device` body field.  

If no `id` parameter was provided, all the user's devices are returned in the 
`devices` body field.

### Errors
- `NotFound`: If the device specified is not found

# Get Battery Readings
The get battery readings endpoint retrieves all a device's battery readings.  

It is located at the `/api/v1/devices/{id}/battery_readings` URL.  

## Request
### Parameters
Required parameters:

- `id` (uint): The ID of the device to retrieve battery readings for

## Response
### Body
The device's [Battery Readings](/server/docs/models/Battery-Reading.md#schema) 
will be returned in the `battery_readings` body field.  

### Errors
- `NotFound`: If the specified device is not found

# Get Auth Tokens
The get auth tokens endpoint retrieves all auth token's issued for the 
specified device.  

It is located at the `/api/v1/devices/{id}/auth_tokens` URL.  

## Request
### Parameters
Required parameters:

- `id` (uint): The ID of the device to retrieve related auth tokens for

## Response
### Body
[Auth Tokens](/server/docs/models/Auth-Token.md#schema) issued for the 
specified device will be returned in the `auth_tokens` body field.

### Errors
- `NotFound`: If the specified device resource is not found

# Update
The update endpoint changes device resource values.  

It is located at the `/api/v1/devices/{id}/` URL.

## Request
### Parameters
Require parameters:

- `id` (uint): The ID of the device resource to update

### Body
One or more of the following body fields must be provided:

- `name` (String): The new name of the device
- `fingerprint` (String): New unique identifying value, used to identify the 
			  device

## Response
### Body
The new [Device](/server/docs/models/Device.md#schema) resource is returned in the `device` body field.

### Errors
- `NotFound`: If the specified device resource to update was not found

# Delete
The delete endpoint destroys a specified device resource.  

It is located at the `/api/v1/devices/{id}/` URL.  

## Request
### Parameters
Required parameters:

- `id` (uint): The ID of the device resource to destroy

## Response
### Body
No extra body fields are returned.
