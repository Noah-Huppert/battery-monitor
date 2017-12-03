# Battery Readings
The battery readings endpoint manages battery reading resources.

# Table Of Contents
- [Overview](#overview)
- [Create](#create)
- [Get](#get)
- [Update](#update)
- [Delete](#delete)

# Overview
The battery readings endpoints is located at the `/api/v1/battery_readings/` 
URL.

# Create
The create endpoint saves one or more new battery readings.  

Located at the `/api/v1/battery_readings/` URL.  

A single battery reading can be provided in the `battery_reading` body field. 
Or an array of battery readings can be provided via the `battery_readings` body 
field.

## Request
Authentication required.

### Body
Required fields:

- `battery_reading` (Battery Reading): Holds the battery reading resource to 
		                        create

OR

- `battery_readings` ([]Battery Reading): Holds the array of battery reading 
					  resources to create

Both these fields only require the following fields of a battery reading 
resource be specified:

- `measure_time` (Time): The time the measurement was taken
- `level` (Integer): The battery percentage at the time of reading. In the 
 		     range of [0, 100]

All other properties of the battery reading resources (`user_id`, `device_id`, 
and `upload_time`) can be determined from other detail of the request.

## Response
### Body
If just a single battery reading was specified, this reading will be returned 
in the `battery_reading` field.  

If multiple battery reading resources were specified, the readings will be 
returned in the `battery_readings` field.

### Errors
Only [Common Errors](/server/docs/Errors.md#common-errors).

# Get
The get endpoint retrieves a single, or multiple battery readings.  

Located at the `/api/v1/battery_readings/{id}/` URL.

If the `id` property is not provided, a list of all the user's battery readings 
will be returned.

## Request
Authentication required.  

### 
