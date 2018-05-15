#!/bin/sh

cd ../data/
wget http://geolite.maxmind.com/download/geoip/database/GeoLiteCity_CSV/GeoLiteCity-latest.zip
failed=`echo $?`
if [ $failed == "0" ]
then 
  rm -rf GeoLiteCity_*
  unzip GeoLiteCity-latest.zip 
else
  echo "ERROR: Can't get http://geolite.maxmind.com/download/geoip/database/GeoLiteCity_CSV/GeoLiteCity-latest.zip"
fi
