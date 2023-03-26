#!/bin/bash

# Menentukan variabel
tegar=$1
facebook_url="https://www.facebook.com/$2"
linkedin_url="https://www.linkedin.com/in/$3"

# Membuat folder
mkdir "$tegar at $(date)"

# Membuat file facebook.txt dan menambahkan url ke dalamnya
touch "$tegar at $(date)/about_me/personal/facebook.txt"
echo "$facebook_url" >> "$tegar at $(date)/about_me/personal/facebook.txt"

# Membuat file linkedin.txt dan menambahkan url ke dalamnya
touch "$tegar at $(date)/about_me/professional/linkedin.txt"
echo "$linkedin_url" >> "$tegar at $(date)/about_me/professional/linkedin.txt"

# Membuat file list_of_my_friends.txt dan mengambil daftar nama teman-teman
touch "$tegar at $(date)/my_friends/list_of_my_friends.txt"
curl https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/fd8b1ad6e4e7ee4d06c4c4e2c6d9d6f22b6f1ba6/list_of_my_friends.txt >> "$tegar at $(date)/my_friends/list_of_my_friends.txt"

# Membuat file about_this_laptop.txt dan menambahkan informasi tentang laptop
touch "$tegar at $(date)/my_system_info/about_this_laptop.txt"
echo "My tegar: $USER" >> "$tegar at $(date)/my_system_info/about_this_laptop.txt"
echo "With host: $(uname -a)" >> "$tegar at $(date)/my_system_info/about_this_laptop.txt"

# Membuat file internet_connection.txt dan menambahkan informasi tentang koneksi internet
touch "$tegar at $(date)/my_system_info/internet_connection.txt"
echo "Connection to google:" >> "$tegar at $(date)/my_system_info/internet_connection.txt"
ping -c 3 google.com >> "$tegar at $(date)/my_system_info/internet_connection.txt"
