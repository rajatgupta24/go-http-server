To start the database server locally run: 

```bash
docker run -d -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=root mysql
```
```bash
docker exec -it mysql mysql -uroot -proot -e 'CREATE DATABASE todolist'
```
