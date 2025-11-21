import os
import shutil

for day in range(1, 26):
    day_folder = f"cmd/day{day}"
    shutil.copytree("cmd/dayX", day_folder)
    
    for root, _, files in os.walk(day_folder):
        for file in files:
            old_file_path = os.path.join(root, file)
            new_file_name = file.replace('dayX', f'day{day}')
            new_file_path = os.path.join(root, new_file_name)
            
            os.rename(old_file_path, new_file_path)
            
            with open(new_file_path, 'r') as f:
                content = f.read()
            new_content = content.replace('dayX', f'day{day}')
            
            with open(new_file_path, 'w') as f:
                f.write(new_content)
