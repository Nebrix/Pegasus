import os

def version():
    try:
        with open('.version', 'r') as file:
            lines = file.readlines()
            for line in lines:
                if line.startswith('VERSION='):
                    return line.strip()[8:]
    except FileNotFoundError:
        print("Error: .version file not found.")
    except Exception as e:
        print(f"Error reading version: {e}")
    
    return "Unknown"

def get_username():
    return os.getenv('USER') or os.getenv('LOGNAME') or os.getenv('USERNAME') or 'Unknown'

def RetrieveDistributionName(Distro, DistroSize) -> None:
    try:
        with open("/etc/os-release", "r") as fp:
            for line in fp:
                if line.startswith("NAME="):
                    nameStart = line.find('=')
                    if nameStart != -1:
                        nameStart += 1 
                        if line[nameStart] == '"':
                            nameStart += 1
                            nameEnd = line.find('"', nameStart)
                            if nameEnd != -1:
                                Distro[:nameEnd - nameStart] = line[nameStart:nameEnd]
                                return
        Distro[:len("Unknown")] = "Unknown"
    except Exception:
        Distro[:len("Unknown")] = "Unknown"

def ascii():
    Username = get_username()

    distro = ["" for _ in range(256)]
    RetrieveDistributionName(distro, len(distro))
    distro = ''.join(distro).strip()

    print("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣤⡀⠀⠀⠀⠀⠀")
    print("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⣧⠃⡃⠀⠀⠀⠀⠀")
    print("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⢿⠇⠏⡇⠀⠀⠀⠀")
    print("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣟⠏⡜⠰⢣⠀⠀⠀⠀")
    print("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡴⠟⣡⠊⡠⢁⣎⠀⠀⠀⠀")
    print("⠀⠀⢀⣶⣰⡂⠀⠀⠀⠀⣠⠖⠋⡔⠊⡠⢏⡠⢃⡜⠀⠀⠀⠀")
    print("⠀⠀⡧⣩⠾⢹⣓⣤⠀⠘⣯⡀⠀⣗⣋⣤⢃⠔⢩⠆⠀⠀⠀⠀")
    print("⠀⣸⣅⠁⢁⡞⠨⡍⢷⢂⡾⠁⣰⠡⠤⣊⠥⠒⠁⠀⠀⠀⠀⠀")
    print("⠘⠿⠽⠞⢫⠀⠀⠀⢻⠟⣡⡾⠱⡑⠦⠥⣀⣀⣀⠀⠀⠀⠀⠀")
    print("⠀⠀⠀⠀⡎⢀⠐⠀⠁⠈⠺⠷⠣⠃⠁⡀⠈⢿⡇⠨⢢⠀⠀⠀")
    print("⠀⠀⠀⢀⡃⠀⠐⠀⠂⠐⠂⢀⠂⠂⠐⢀⠀⡾⢔⢄⠡⣃⠀⠀")
    print(f"⠀⢸⡏⠃⠀⠙⣞⡆⠀⠀⠀⠀⠀⠀⣵⡚⠱⣎⣇⠀⠀   username: {Username}")
    print(f"⠀⣿⣷⠄⠀⠀⢹⣽⠀⠀⠀⠀⠀⣼⣳⠃⠀⠹⣾⣤⠀   distro: {distro}")
    print(f"⠀⠉⠁⠀⠀⠰⠿⠟⠀⠀⠀⠀⠘⠛⠛⠀⠀⠀⠹⠋    version: {version()}")