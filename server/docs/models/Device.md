# Device
The device model stores the phones a user owns.  

# Table Of Contents
- [Schema](#schema)
- [Relationships](#relationships)

# Schema
The device model contains the following fields:

- `user_id` (uint): Identifies which user the device belongs to
- `name` (String): Name of device, given by user
- `fingerprint` (String): Unique value used to identify the device 

# Relationships
A device belongs to a user. This user is specified by the `user_id` field. In 
a One to One relationship.
