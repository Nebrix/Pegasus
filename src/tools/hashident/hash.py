import hashid

def identify_hash(hash_value):
    hasher = hashid.HashID()
    results = list(hasher.identifyHash(hash_value))
    if results:
        return results[0][0]
    return "Unknown"

if __name__ == '__main__':
    hash_value = input("Enter hash: ")
    hash_algorithm = identify_hash(hash_value)
    print("Hash Algorithm:", hash_algorithm)
