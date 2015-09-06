## How to start up

Download the code

```
go get -u github.com/theplant/hybris_qor_cms
```


Go to the code directory

```
cd $GOPATH/src/github.com/theplant/hybris_qor_cms
```


Setup the environment variables

```
source dev_env
```


Create database

```
mysql -u root -e "create database hybris_qor_cms"
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
