# Models
The HTTP API stores data defined by schemas. 

# Table Of Contents
- [Overview](#overview)
- [Base Model](#base-model)
- [Storage](#storage)

# Overview
The HTTP API defines the following models:

- [User](/server/docs/models/User.md)
- [Auth Token](/server/docs/models/Auth-Token.md)
- [Device](/server/docs/models/Device.md)
- [Battery Reading](/server/docs/models/Battery-Reading.md)

The schemas of these models are defined by Go `struct`s. Every single model 
contains some common fields, described in the [Base Model](#base-model) section 
of this page.

# Base Model
Models all have a few fields in common:

- `id` (uint): Primary key of model in database, unique identifier
- `created_at` (Time): Date and time model was created
- `updated_at` (Time): Date and time model was last updated

# Storage
Models will be stored in PostgreSQL. Managed by the 
[GORM](http://jinzhu.me/gorm/) ORM.
