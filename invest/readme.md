The root by default listens only to localhost connections. need to create a new user, who can listen to all IP connections.
create user if not exists `duck`@`%`
 % wildcard indicates that the duck user should listen to all IPs.

Your default MySQL config only accepts connections from 127.0.0.1 and not from any other IP address. To resolve this, we must update the config to allow binding (accepting) connections from all IP addresses. You can set specific IP addresses, but WSL’s IP addresses vary, so this is an easier alternative.
Update the bind-address property from 127.0.0.1 to 0.0.0.0 in the MySQL Config file located at /etc/mysql/mysql.conf.d/mysqld.cnf and restart the MySQL server


## DB 설정

```sh
$ docker run -v /invest/db:/var/lib/mysql --name investDb -e MYSQL_ROOT_PASSWORD=root -d -p 3306:3306 mysql 
```


## Mock 설정
```sh
go install github.com/vektra/mockery/v2@v2.44.1
go generate ./...
```

//go:generate mockery --name {interface name} --case underscore --inpackage

mockery는  mock의 메소드별 행동을 명시적으로 정의해야 함.
모든 메소드에 대해서 행동을 지정하기 위해서는 reflect를 사용하여 모든 메소드들을 loop 돌면서 지정 가능
```go
ref := reflect.TypeOf(MyInterface)
for i :=0; i < ref.NumMethod(); i++{
    method := ref.Method(i)
    mockObj.On(method.Name).Return(nil)
}

```