#!/usr/bin/env bash
if [ -f .env ]
then
  echo -e "[INFO]Loading environment from local .env file"

  export $(cat .env | sed 's/#.*//g' | xargs)
fi

echo -e "[INFO]Starting Go\n"
go $@
