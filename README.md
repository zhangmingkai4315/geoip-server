# geoip-server
Using Free GeoIP databases, transfer to redis database for api query

#### Download

when you checkout the git repo, please follow the next instructions to install database to local 
server.

```
./scripts/updatefile.sh
``` 
> you can put this code to crontab and download new files any time

#### Save To Redis

make sure you redis server is running(not local? change the **REDIS_SERVER** in scripts/*.py).

```
pip install -r requirements.txt
cd scripts
python import_asn_to_redis.py
python import_ip_to_redis.py
```
the first command will install deps for python scripts. then import asn data and ip data to redis.after install it will print one test result for test if it success.

##### RunServer

The server is a golang progrom. You can download the binary file from release page or install golang development env for change code as you wish.


#### Databases

GeoLite2 databases are free IP geolocation databases comparable to, but less accurate than, MaxMind’s GeoIP2 databases. The GeoLite2 Country and City databases are updated on the first Tuesday of each month. The GeoLite2 ASN database is updated every Tuesday.

[download](https://dev.maxmind.com/geoip/geoip2/geolite2/)

#### License
This product includes GeoLite2 data created by MaxMind, available from [http://www.maxmind.com](http://www.maxmind.com)


