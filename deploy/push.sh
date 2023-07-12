#!/bin/bash

export GITSHA=`git rev-parse --short=8 HEAD`

docker tag whisperingdice 171055152598.dkr.ecr.us-east-1.amazonaws.com/whisperingdice:$GITSHA

aws ecr get-login-password | docker login -u AWS --password-stdin https://171055152598.dkr.ecr.us-east-1.amazonaws.com

docker push 171055152598.dkr.ecr.us-east-1.amazonaws.com/whisperingdice:$GITSHA

sed "s/\$REVISION_HASH/$GITSHA/g" deploy/task-template.json > deploy/task.json

aws ecs register-task-definition --family "Whispering-Dice" --requires-compatibilities FARGATE --cli-input-json file://deploy/task.json

aws ecs update-service --cluster Whispering-Dice --service Whispering-Dice --task-definition Whispering-Dice
