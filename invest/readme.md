

## DB 설정

```sh
$ docker run -v /invest/db:/var/lib/mysql --name investDb -e MYSQL_ROOT_PASSWORD=root -d -p 3306:3306 mysql 
```