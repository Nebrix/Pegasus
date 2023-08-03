import hashlib

def hash_string(string, algorithm):
    if algorithm == 'md5':
        return hashlib.md5(string.encode()).hexdigest()
    elif algorithm == 'sha1':
        return hashlib.sha1(string.encode()).hexdigest()
    elif algorithm == 'sha256':
        return hashlib.sha256(string.encode()).hexdigest()
    else:
        return "Error: Unsupported algorithm"
    
if __name__ == "__main__":
    string = input("Enter string: ")
    algorithm = input("Enter algorithm: ")
    print(hash_string(string, algorithm))