# Go-MySQL-app

Things we're going to do in this tutorial are:

- Create a go-app
- Dockerize the app
- Deploy the app on k3s

To start the mysql server locally run: 
```bash
docker run -d -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=root mysql
```

This will open mysql in your terminal, and will also create a database named todolist.
```bash
docker exec -it mysql mysql -uroot -proot -e 'CREATE DATABASE todolist'
```
After creating a database, we need to create a table
```sql
CREATE TABLE todos
```

To build your own image run:
```bash
docker build -t go-simple .
```

To run the image & to connect it to database container run: 
```bash
docker run -d --network host --name go-simple go-simple
```

To test the app working run or visit http://localhost:9000/home in your browser.
```bash
curl http://localhost:9000/home
```
