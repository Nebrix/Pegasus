import argparse
import requests

def check_dir_exists(url):
    response = requests.head(url)
    return response.status_code in (200, 301)

def main():
    parser = argparse.ArgumentParser(description='Directory buster')
    parser.add_argument('base_url', help='Base URL for directory scanning')
    args = parser.parse_args()

    base_url = args.base_url

    try:
        with open('wordlist.txt', 'r') as file:
            directories = file.read().splitlines()
    except FileNotFoundError as e:
        print(f"Error opening wordlist.txt: {e}")
        return

    print("Directory buster started...")

    for dir in directories:
        url = f"{base_url}/{dir}"
        if check_dir_exists(url):
            print(f"[GET] Directory found: {url}")

if __name__ == "__main__":
    main()
