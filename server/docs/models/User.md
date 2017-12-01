# User
The user model holds information about people using the Battery Monitor 
platform.  

# Table Of Contents
- [Schema](#schema)
- [Relationships](#relationships)

# Schema
The user model has the following fields:

- `first_name` (String): Peron's given name
- `last_name` (String): Person's family name
- `email` (String): Person's email address
- `email_verified` (Boolean): Indicates if the email should be trusted as 
			      belonging to the user

# Relationships
Users can own multiple devices. This is represented via a One to many 
relationship between a User and multiple Devices.
