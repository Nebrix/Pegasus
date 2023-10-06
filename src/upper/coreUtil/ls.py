import os
import stat
import time

def human_readable_size(size):
    if size >= 1024 * 1024:
        return f"{size / (1024 * 1024):.1f} Mib"
    elif size >= 1024:
        return f"{size / 1024:.1f} KiB"
    else:
        return f"{size} B"
    
def human_readable_time(mod_time):
    now = time.time()
    seconds_ago = now - mod_time

    if seconds_ago < 60:
        return f"{int(seconds_ago)} seconds ago"
    elif seconds_ago < 3600:
        return f"{int(seconds_ago / 60)} minutes ago"
    elif seconds_ago < 86400:
        return f"{int(seconds_ago / 3600)} hours ago"
    elif seconds_ago < 604800:
        return f"{int(seconds_ago / 86400)} days ago"
    else:
        return time.strftime("%b %d, %Y %H:%M", time.localtime(mod_time))
    
def print_table_row(number, name, file_type, size, modified):
    print(f"| {number:3d} | {name:16s} | {file_type:4s} | {human_readable_size(size):10s} | {human_readable_time(modified):20s} |")

def print_table_header():
    print("╭─────┬──────────────────┬──────┬────────────┬──────────────────────╮")
    print("│ %-3s │ %-16s │ %-4s │ %-10s │ %-20s │" % ("#", "Name", "Type", "Size", "Modified"))
    print("├─────┼──────────────────┼──────┼────────────┼──────────────────────┤")

def print_table_footer():
    print("╰─────┴──────────────────┴──────┴────────────┴──────────────────────╯")

def list_command():
    number = 0

    print_table_header()

    for entry in os.listdir("."):
        filestat = os.stat(entry)
        file_type = "dir" if stat.S_ISDIR(filestat.st_mode) else "file"
        print_table_row(number, entry, file_type, filestat.st_size, filestat.st_mtime)
        number += 1

    print_table_footer()

if __name__ == "__main__":
    list_command()