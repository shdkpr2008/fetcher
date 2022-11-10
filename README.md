# About Fetcher
Fetcher is a command line utility to keep record of urls as html and metadata such as number of 
links and images.

A details description of components, libraries, assumptions and task is available at 
https://www.notion.so/shdkpr2008/Autify-Technical-Assignment-f52e09ce30e546bf8ccee532b69402c0  

## BrowserLess
This is a browser-less build, for network please check main/master.

## Prerequisite
- Go (v1.19.3)
- SQLite (v3)
- Docker (optional)
- Make (optional)
- Write permissions of current working directory.

## Environment Variables
Following variables are to be set correctly before proceeding with make:
```
GOOS=<os>
GOARCH=<arch>
```
Please replace the placeholder values along with opening and closing braces.

## Run on Docker
```
docker build -t fetch .  
docker run -it --volume $(pwd):/app fetch:latest "https://www.google.com" "https://www.youtube.com"
docker run -it --volume $(pwd):/app fetch:latest "https://www.google.com" "https://www.youtube.com" --metadata  
```

## Reset
To remove all the local records along with metadata
```
rm -rf *.html
rm -rf database.db
```

## ToDo
Assumptions and tasks that were not done due to time constraint are mentioned at
https://www.notion.so/shdkpr2008/Autify-Technical-Assignment-f52e09ce30e546bf8ccee532b69402c0

### Disclaimer
This repository or peace of code uses other open source libraries and pieces of codes, 
for more information please visit the detailed description section or checkout the plugins 
by visiting the module files.
