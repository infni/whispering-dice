#!/bin/bash

. ./local.sh # sets the bot token, app_id and guild_id

docker build --network host . -t whisperingdice --build-arg BOT_TOKEN --build-arg APP_ID --build-arg GUILD_ID