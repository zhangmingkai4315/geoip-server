import socket
import ipaddress
import sys


def update_progress(job_title, progress):
    """Update progress in thread"""
    length = 20
    block = int(round(length*progress))
    msg = "\r{0}: [{1}] {2}%".format(
        job_title,
        "#"*block + "-"*(length-block),
        round(progress*100, 2))
    if progress >= 1:
        msg += " DONE\r\n"
    sys.stdout.write(msg)
    sys.stdout.flush()


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
