
# Voicebot-2

## Description

I want to write a small Discord Bot to listen to events when a user enters a specific voice channel and create a new temporal channel for this user. When the last user leaves the temporal channel, the bot should remove this channel.

## Planned Features
- Create new voice channels
- Move the user into the new channel
- keep track how many users are inside the channel
- remove the channel once everyone is disconnected
- create a matching text channel
- set the permissions for only the people inside the voice channel to access the text channel
- use names from a database for the channels
- API for a web interface for the names
- maybe maintenance (Remove channels, minor fix ups) from a web interface

## Goal

A year ago, I wrote a small bot to create temporal Channels in Discord. At this time I had very limited knowledge of Go and a few very bad design choices like reading the Config from a json file (Why did I choose not to just use Environment Variables?) and I used a sqlite Database and wrote the sql Commands myself to execute the Queries.

with this project I want to completely rewrite my old project an implement a few new features - and I don't want to depend on any kind of Stateful data besides of the Database - to be able to run this inside a Kubernetes cluster (With just one Replica - Maybe I can implement something to fix even that problem)

## Milestones

- [ ] Save channel names inside a Database
- [ ] Listen to Events when a user enters a specific channel
- [ ] Create and move user to the new channel
- [ ] Set permissions
- [ ] keep track how many users are inside a channel
- [ ] cleanup

