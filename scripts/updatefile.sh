#!/bin/bash

BASEDIR=$(dirname "$0")
DATA_FOLDER=$(dirname "$0")/../data/
BACKUP_FOLDER=$(dirname "$0")/../data/backup
TEMP_FOLDER=$(dirname "$0")/../data/temp

platform='unknown'
unameStr=`uname`
if [[ "$unameStr" == "Linux" ]];then
  platform="linux"
elif [[ "$unameStr" == "Darwin" ]];then
  platform="macosx"
fi
	
[ ! -d $DATA_FOLDER ] && mkdir -p $DATA_FOLDER
[ ! -d $BACKUP_FOLDER ] && mkdir -p $BACKUP_FOLDER
[ ! -d $TEMP_FOLDER ] && mkdir -p $TEMP_FOLDER


GEOIP2_DOWNLOAD_URL=http://geolite.maxmind.com/download/geoip/database/GeoLite2-City-CSV.zip
GEOIP2_DOWNLOAD_MD5_URL=http://geolite.maxmind.com/download/geoip/database/GeoLite2-City-CSV.zip.md5
GEOIP2_DATA_FOLDER=GeoLite2-City-CSV


GEOIP2_ASN_DOWNLOAD_URL=http://geolite.maxmind.com/download/geoip/database/GeoLite2-ASN-CSV.zip
GEOIP2_ASN_DOWNLOAD_MD5_URL=http://geolite.maxmind.com/download/geoip/database/GeoLite2-ASN-CSV.zip.md5
GEOIP2_ASN_DATA_FOLDER=GeoLite2-ASN-CSV

# check remote file md5 if the same then return 0 
check_md5()
{
  newRemoteMD5=$(wget $1 -q -O -)
  if [ "$2" = "$newRemoteMD5" ];then
    return 0
  else
    return 1
  fi
}

# Download file from url
download_file()
{
  FILE_URL=$1
  filename=`basename $FILE_URL`
  md5filePath=$BACKUP_FOLDER/$filename.md5
  if [ -f $md5filePath ];then
    oldMD5Value=`cat $md5filePath`
    check_md5 $2 $oldMD5Value
    if [ $? -eq 0 ];then 
      return 0
    fi
   
  fi
  # file has changed md5 file value not equal 
  result=$(wget $1 -O $TEMP_FOLDER/$filename)
  if [ $? -eq 0 ];then 
    echo "download sucess"
    rm -rf $DATA_FOLDER/$3
    unzip -o $TEMP_FOLDER/$filename -d $DATA_FOLDER
    # mv $3* $3
    find ./data -type d -name "$3_*" -exec mv {} $DATA_FOLDER/$3 \;
    mv -f $TEMP_FOLDER/$filename $BACKUP_FOLDER/$filename
    if [[ "$platform" == "macosx" ]];then
      newMD5Value=`md5 $BACKUP_FOLDER/$filename | awk -F " = " '{print $2}'`
    else
      newMD5Value=`md5sum $BACKUP_FOLDER/$filename | awk '{ print $1 }'`
    fi
    echo $newMD5Value > $md5filePath
  else
    echo "ERROR: Can't download $filename"
  fi

}

download_file $GEOIP2_ASN_DOWNLOAD_URL $GEOIP2_ASN_DOWNLOAD_MD5_URL $GEOIP2_ASN_DATA_FOLDER
download_file $GEOIP2_DOWNLOAD_URL $GEOIP2_DOWNLOAD_MD5_URL $GEOIP2_DATA_FOLDER

