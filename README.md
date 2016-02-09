# gotrain

## Tips for installation

192.168.2.147:8000 is my Laptop's Webserver, so please adopt accordingly. 

Windows:
-------

install git, 7zip, golang, liteide from

http://192.168.2.147:8000/windows/

either 32 or 64 bits,

liteide is only extracted, put a link on your desktop


Linux:
------

golang install in Linux 32 bit:

wget http://192.168.2.147:8000/linux/32/go1.5.3.linux-386.tar.gz
sudo tar -C /usr/local -xzf go1.5.3.linux-386.tar.gz

add this line to ~/.profile :

export PATH=$PATH:/usr/local/go/bin

golang install in Linux 64 bit:

wget http://192.168.2.147:8000/linux/64/go1.5.3.linux-amd64.tar.gz.tar.gz
sudo tar -C /usr/local -xzf go1.5.3.linux-amd64.tar.gz

add this line to ~/.profile :

export PATH=$PATH:/usr/local/go/bin


ToBeContinued .....



mkdir C:\gotrain
cd C:\gotrain
set GOPATH=C:\gotrain
go get github.com/iplon/gotrain

cd src\github.com\iplon\gotrain\cmd\L1
go build
.\L1

then start Liteide

add Folder

adopt GOPATH definition in env File (to be explained)

.......
