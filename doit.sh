#!/bin/sh
ls /app
/bin/bash -c "/app/wait-for-it.sh --host=postgres-svc --port=5432 --timeout=600 --strict -- /app/godddnews"
 
