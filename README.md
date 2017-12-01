# Battery Monitor
Battery Monitor records device battery levels. It can be used to analyse and 
compare different device's battery life.

# Table Of Contents
- [Overview](#overview)
- [App](#app)
- [Server](#server)

# Overview
Battery Monitor is made up of two components. The app (For both iOS and 
Android) and the server.  

The app measures battery levels, and sends them to the server. The server 
stores these readings, and makes them easy to view.  

# App
The Battery Monitor application lets users automatically submit battery 
readings from their devices.  

The mobile app source code is located in the `/app` directory. See the 
[App README](/app/README.md).

# Server
The server provides a web ui and HTTP API for the battery monitor platform.  

The server source code is located in the `/server` directory. See the 
[Server README](/server/README.md).
