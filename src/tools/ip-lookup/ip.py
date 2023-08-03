import requests

def ip_lookup(ip_address):
    response = requests.get(f"https://ipinfo.io/{ip_address}/json")
    data = response.json()
    return data

if __name__ == '__main__':
    ip_address = input("Enter IP: ")
    data = ip_lookup(ip_address)
    print(f"Location: {data['city']}, {data['region']}, {data['country']}")
    print(f"ISP: {data['org']}")
    print(f"Location: {data['loc']}")
    print(f"Postal Code: {data['postal']}")
