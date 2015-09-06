## How to start up

Download the code

```
go get -u github.com/theplant/device_management
```


Go to the code directory

```
cd $GOPATH/src/github.com/theplant/device_management
```


Setup the environment variables

```
source dev_env
```


Create database

```
mysql -u root -e "create database device_management"
```


Start the app

```
go run main.go
```


Go to URL

```
http://localhost:9000/admin/

```
# hybris_qor_cms
