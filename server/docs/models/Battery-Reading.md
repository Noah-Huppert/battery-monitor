# Battery Reading
The battery reading model is used to record a device's battery level at a 
specific time.

# Table Of Contents
- [Schema](#schema)
- [Relationships](#relationships)

# Schema
The battery reading model contains the following fields:

- `user_id` (uint): The ID of the User who submitted the reading
- `device_id` (uint): The ID of the device which the reading was taken from
- `measure_time` (Time): The date and time the battery reading was taken
- `upload_time` (Time): The date and time the battery reading was uploaded to 
			the server
- `level` (Integer): The battery percentage at the time of reading. From [0, 
    		     100]

# Relationships
A battery reading belongs to a user and device. 

The user who submitted the rating is specified by the `user_id` field. In a One 
to One relationship.  

The device the reading was taken from is specified by the `device_id` field. In 
a One to One relationship.
