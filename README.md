# golang-account-simple
A simple Go web app to access mysql.

It's named account because it seemed like a simple enough example to use with an orm.

Make sure you understand go enough to have set your GOPATH. Recommended reading: https://github.com/alco/gostart#faq0

TLDR: 
* Make a directory named go. You must commit to forever more have anything related to go in this directory. Sorry, that's how go works. Set GOPATH= the fully qualified path to the go directory.
```
mkdir ~go && echo "export GOPATH=$HOME/go" >> ~/.bashrc && source ~/.bash_rc
```
     
* Make the following directories in your go folder: bin, build, pkg, src
```
mkdir $GOPATH/bin && mkdir $GOPATH/build && mkdir $GOPATH/pkg && mkdir $GOPATH/src
```
* Change into your src directory, and clone this package:
```
cd $GOPATH/src && git clone https://github.com/fimbulvetr/golang-account-simple.git
```

Make sure you have [glide](http://glide.sh/), a golang package manager installed.

Then, simply:

```
glide update
```

create your database on mysql:
```
CREATE DATABASE `account` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin */;
```

create your table in the newly created account database:
```
CREATE TABLE `account` (
  `Id` int(11) NOT NULL,
  `Username` varchar(50) COLLATE utf8mb4_bin NOT NULL,
  `Type` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin
```

create some sample data:

```
INSERT INTO `account` VALUES (2,'you',1);
```

Run:

```
ACCOUNT_DB_USER=$dbuser ACCOUNT_DB_PASS=$dbpassword ACCOUNT_DB_HOST=127.0.0.1:3306 ACCOUNT_DB_NAME=$db_name go run main.go
```