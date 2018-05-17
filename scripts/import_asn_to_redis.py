# coding=utf-8
import csv
import redis

from utils import cidr_v4_to_score, is_valid_ipv4_address, ipv4_to_score

DATA_FOLDER = '../data/GeoLite2-ASN-CSV/'
ASN_BLOCKS_IPv4_FILE = DATA_FOLDER+'GeoLite2-ASN-Blocks-IPv4.csv'
# not use ipv6
# ASN_BLOCKS_IPv6_FILE = DATA_FOLDER+'GeoLite2-ASN-Blocks-IPv6.csv'

REDIS_SERVER = '127.0.0.1'


def import_asn_to_redis(connection, asn_file):
    count = 0
    with open(asn_file) as csvfile:
        reader = csv.DictReader(csvfile)
        for row in reader:
            count = count+1
            try:
                start_ip = row['network']
                score = cidr_v4_to_score(start_ip)
                asn_id = row['autonomous_system_number']
                organization = row['autonomous_system_organization']
                connection.hset('asnid:', asn_id, organization)
                asn_id_with_count = asn_id+'_'+str(count)
                connection.zadd('ip2asnid:', asn_id_with_count, score)
            except Exception:
                continue


def find_ipv4_asn_info(connection, ipaddress):
    if not is_valid_ipv4_address(ipaddress):
        return None
    ipaddress = ipv4_to_score(ipaddress)
    asn_id = connection.zrevrangebyscore(
        'ip2asnid:',
        ipaddress,
        0,
        start=0,
        num=1)
    print "Test ipaddress is %s\n" % ipaddress
    if not asn_id:
        return None
    asn_id = asn_id[0].partition('_')[0]
    return connection.hget("asnid:", asn_id)


def main():
    r = redis.Redis(host=REDIS_SERVER)
    print "Begin to import asn data to redis db ..."
    import_asn_to_redis(r, ASN_BLOCKS_IPv4_FILE)
    print "Save data success"
    print(find_ipv4_asn_info(r, '1.2.4.8'))


if __name__ == '__main__':
    main()
