import requests
import random

from requests import async

channels = [
    "shanson",
    "auto_plus",
    "cam25_300",
    "russia_today_hd",
    "mtv_live_hd",
    "failed_case_1",
    "failed_case_2",
    "failed_case_3"
]

failed_cases = [ 
    "failed_case_1",
    "failed_case_2",
    "failed_case_3"
]

ips = []

for i in range(250):
    ips.append("10.240.230.%s"%i)

for i in range(250):
    ips.append("1.1.3.%s"%i)

ip_requested = []

def test_auth():
    for i in range(1000):
        ip = random.choice(ips)
        channel = random.choice(channels)
        if ip not in ip_requested and channel not in failed_cases:
            ip_requested.append(ip)

        headers = {
            'iptest':ip,
            'channeltest':channel
        }

        r = requests.get("http://10.240.230.55:8080/auth", headers=headers)

        print("case %s channel %s ip %s status %s"%(i, channel, ip, r.status_code))
    
    print("Amount ip in cache %s"%len(ip_requested))

async_list = []

def hook_response(response):
    print(response.status_code)

def test_auth_async():
    for i in range(1000):
        ip = random.choice(ips)
        channel = random.choice(channels)
        if ip not in ip_requested and channel not in failed_cases:
            ip_requested.append(ip)

        headers = {
            'iptest':ip,
            'channeltest':channel
        }

        action_item = async.get("http://10.240.230.55:8080/auth", headers=headers, hooks = {'response': hook_response})

        async_list.append(action_item)

async.map(async_list)
    
    print("Amount ip in cache %s"%len(ip_requested))


def test_archive():
    for i in range(1000):
        ip = random.choice(ips)
        channel = random.choice(channels)
        if ip not in ip_requested:
            ip_requested.append(ip)

        headers = {
            'iptest':ip,
        }

        r = requests.get("http://10.240.230.55:8080/playlist/program/assdfsfd.m3u8", headers=headers)

        print("case %s channel %s ip %s status %s"%(i, channel, ip, r.status_code))
    
    print("Amount ip in cache %s"%len(ip_requested))

test_auth()
test_archive()

