The root by default listens only to localhost connections. need to create a new user, who can listen to all IP connections.
create user if not exists `duck`@`%`
 % wildcard indicates that the duck user should listen to all IPs.

Your default MySQL config only accepts connections from 127.0.0.1 and not from any other IP address. To resolve this, we must update the config to allow binding (accepting) connections from all IP addresses. You can set specific IP addresses, but WSL’s IP addresses vary, so this is an easier alternative.
Update the bind-address property from 127.0.0.1 to 0.0.0.0 in the MySQL Config file located at /etc/mysql/mysql.conf.d/mysqld.cnf and restart the MySQL server

go get -u github.com/chromedp/chromedp // 삭제


