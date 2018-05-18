# coding=utf-8
import csv
import json
import redis

from utils import is_valid_ipv4_address, ipv4_to_score, update_progress

DATA_FOLDER = '../data/'
CITY_BLOCKS_IPv4_FILE = DATA_FOLDER+'china.csv'
REDIS_SERVER = '127.0.0.1'


def import_info_to_redis(connection, location_file):
    count = 0
    key = 'ip2chinainfo:'
    with open(location_file) as csvfile:
        reader = csv.reader(csvfile)
        row_count = sum(1 for row in reader)
        print "Total row counter is %d" % row_count
        csvfile.seek(0)
        update_progress("Loading IP Info Data", 0)
        for row in reader:
            count = count+1
            if count % 1000 == 0:
                update_progress("Loading IP Info Data", count/float(row_count))
            country_name = row[4].strip('"')
            provience_name = row[5].strip('"')
            city_name = row[6].strip('"')
            isp_name = row[8].strip('"')
            connection.hset(key, str(count), json.dumps([
                country_name,
                provience_name,
                city_name,
                isp_name]))
        update_progress("Loading IP Info Data", 1)


def import_ipv4_to_redis(connection, block_file):
    """
    import_ipv4_to_redis will import ip information to redis cache
    """
    with open(block_file) as csvfile:
        reader = csv.reader(csvfile)
        row_count = sum(1 for row in reader)
        count = 0
        csvfile.seek(0)
        update_progress("Loading IPv4 ID Data", 0)
        for row in reader:
            count = count+1
            if count % 1000 == 0:
                update_progress("Loading IPv4 ID Data", count/float(row_count))
            start_ip = row[0].strip('"') if row else ''
            score = ipv4_to_score(start_ip)
            _id = str(count)
            connection.zadd('ip2cnid:', _id, score)
        update_progress("Loading IPv4 ID Data", 1)


def find_ipv4_info(connection, ipaddress):
    print "Test ip address is %s " % (ipaddress)
    if not is_valid_ipv4_address(ipaddress):
        return None
    ipaddress = ipv4_to_score(ipaddress)
    _id = connection.zrevrangebyscore(
        'ip2cnid:',
        ipaddress,
        0,
        start=0,
        num=1)
    if not _id:
        return None
    _id = _id[0].partition('_')[0]
    return json.loads(connection.hget('ip2chinainfo:', _id))


def main():
    r = redis.Redis(host=REDIS_SERVER)
    print "Begin to load cn ipv4 id info to redis db ..."
    import_ipv4_to_redis(r, CITY_BLOCKS_IPv4_FILE)
    print "Load Success"

    print "Begin to load cn ip info to redis db ..."
    import_info_to_redis(r, CITY_BLOCKS_IPv4_FILE)
    print "Load Success"
    results = find_ipv4_info(r, '1.2.4.8',)
    print results


if __name__ == '__main__':
    main()
