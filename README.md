# anti-var-_0x-util
That sample utilite made for automaticly remove aftermath of malware which add subscript to your js and php files. 
Typicly malware's addition code format like:
<script type='text/javascript'>var _0x...
  ...if(n==!![]){a();}</script>
# HOW TO USE
## Flags
If you want to autoreplace all malwared parts of files, use flag -w.
If you want to find all files, which used eval function, use flag -h.
If you want to remove all parts with eval function in small files (less then 5 lines), use flag -s and -h and -w.
## Windows
1. Download [util.exe](https://github.com/VladChekunov/anti-var-_0x-util/raw/master/util.exe) and copy that to your folder.
2. Run ulit.exe and wait for auto replace all parts of malware code.
## Other platforms
1. Download [util.go](https://github.com/VladChekunov/anti-var-_0x-util/raw/master/util.go).
2. Compile that and copy executable file to your folder.
3. Run and wait for success.
