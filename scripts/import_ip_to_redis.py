# coding=utf-8
import csv
import json
import socket

import ipaddress
import redis


DATA_FOLDER = '../data/'
CITY_BLOCKS_IPv4_FILE = DATA_FOLDER+'GeoLite2-City-Blocks-IPv4.csv'
CITY_BLOCKS_IPv6_FILE = DATA_FOLDER+'GeoLite2-City-Blocks-IPv6.csv'
CITY_LOCATION_EN_FILE = DATA_FOLDER+'GeoLite2-City-Locations-en.csv'
CITY_LOCATION_ZH_CN_FILE = DATA_FOLDER+'GeoLite2-City-Locations-zh-CN.csv'

REDIS_SERVER = '127.0.0.1'


def is_valid_ipv4_address(address):
    """is_valid_ipv4_address will valid ipv4 address """
    try:
        socket.inet_pton(socket.AF_INET, address)
    except AttributeError:  # no inet_pton here, sorry
        try:
            socket.inet_aton(address)
        except socket.error:
            return False
        return address.count('.') == 3
    except socket.error:  # not a valid address
        return False

    return True


def is_valid_ipv6_address(address):
    """is_valid_ipv6_address will valid ipv6 address """
    try:
        socket.inet_pton(socket.AF_INET6, address)
    except socket.error:  # not a valid address
        return False
    return True


def cidr_v4_to_score(cidr):
    """cidr_v4_to_score will convert cidr string to int"""
    net = ipaddress.IPv4Network(unicode(cidr))
    start_ip_address = str(net[0])
    score = 0
    for v in start_ip_address.split('.'):
        score = score*256+int(v, 10)
    return score


def ipv4_to_score(ip):
    score = 0
    for v in ip.split('.'):
        score = score*256+int(v, 10)
    return score


def import_cities_to_redis(connection, location_file, lang):
    count = 0
    key = 'cityid2city:'+lang+':'
    with open(location_file) as csvfile:
        reader = csv.DictReader(csvfile)
        for row in reader:
            count = count+1
            city_id = row['geoname_id']
            continent_code = row['continent_code']
            continent_name = row['continent_name']
            country_code = row['country_iso_code']
            country_name = row['country_name']
            subdivision_code = row['subdivision_1_iso_code']
            subdivision_name = row['subdivision_1_name']
            city_name = row['city_name']
            metro_code = row['metro_code']
            connection.hset(key, city_id, json.dumps([
                continent_code,
                continent_name,
                country_code,
                country_name,
                subdivision_code,
                subdivision_name,
                city_name,
                metro_code]))


def import_ipv4_to_redis(connection, block_file):
    """
    import_ipv4_to_redis will import ip information to redis cache
    """
    with open(block_file) as csvfile:
        reader = csv.DictReader(csvfile)
        count = 0
        for row in reader:
            count = count+1
            start_ip = row['network'] if row else ''
            score = cidr_v4_to_score(start_ip)
            city_id = row['geoname_id']+'_'+str(count)
            connection.zadd('ip2cityid:', city_id, score)


def find_ipv4_info(connection, ipaddress, lang):
    if not is_valid_ipv4_address(ipaddress):
        return None
    ipaddress = ipv4_to_score(ipaddress)
    city_id = connection.zrevrangebyscore(
        'ip2cityid:',
        ipaddress,
        0,
        start=0,
        num=1)
    if not city_id:
        return None
    city_id = city_id[0].partition('_')[0]
    return json.loads(connection.hget('cityid2city:'+lang+":", city_id))


def main():
    r = redis.Redis(host=REDIS_SERVER)

    print "Begin to import ipv4 info to redis db ..."
    import_ipv4_to_redis(r, CITY_BLOCKS_IPv4_FILE)
    print "Import ipv4 info success"
    print "Begin to import city  info to redis db ..."
    import_cities_to_redis(r, CITY_LOCATION_ZH_CN_FILE, 'zh-ch')
    print "Import city info success"
    import_cities_to_redis(r, CITY_LOCATION_EN_FILE, 'en')
    # results = find_ipv4_info(r, '172.217.3.206', 'zh-ch')
    # print results
    # results = find_ipv4_info(r, '172.217.3.206', 'en')
    # print results


if __name__ == '__main__':
    main()
