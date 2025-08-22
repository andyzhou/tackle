# tackle
Anonymous open tool web service

#tips
* how build cross platform
  CC=/usr/local/gcc-4.8.1-for-linux64/bin/x86_64-pc-linux-gcc CGO_ENABLED=1 GOARCH=amd64 GOOS=linux go build


#video2gif
- upload video and convert to gif
- file data storage in `pond` big file, include snap and gif file
- doc and count data storage in sqlite db