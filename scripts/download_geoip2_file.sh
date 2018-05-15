#!/bin/bash

BASEDIR=$(dirname "$0")
DATA_FOLDER=$(dirname "$0")/../data/
GEOIP2_DOWNLOAD_URL=http://geolite.maxmind.com/download/geoip/database/GeoLiteCity_CSV/GeoLiteCity-latest.zip
# check remote file md5 if the same then quit download process

# end

cd $DATA_FOLDER

wget $GEOIP2_DOWNLOAD_URL 
failed=`echo $?`
if [ $failed == "0" ]
then 
  rm -rf GeoLiteCity_*
  unzip GeoLiteCity-latest.zip 
else
  echo "ERROR: Can't get http://geolite.maxmind.com/download/geoip/database/GeoLiteCity_CSV/GeoLiteCity-latest.zip"
fi
