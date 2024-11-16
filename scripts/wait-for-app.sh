#!/bin/sh

# Wait for the app to become ready
until curl -s http://app:8000/health; do
	echo "Waiting for app..."
	sleep 2
done

# Start Nginx
exec nginx -g 'daemon off;'
