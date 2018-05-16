import socket
import ipaddress

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
