# Go_CRUD
A gofr based application implementing CRUD operations on MySQL based docker database image.

# Instructions to run
* Open terminal and run in the project directory.

```sh
go run main.go 
```

* Go to http://localhost:8080/welcome to see the welcome page <br />
* Go to http://localhost:8080/get_cars to see information about the cars.<br />
* Go to http://localhost:8080/cars/delete to delete repaired cars.<br />
* Go to http://localhost:8080/cars/{car_id}/{car_status} replacing car_id and car_status with your values to update the database.<br />
* Go to http://localhost:8080/cars/{car_name}/{car_status} to add new cars to the database.<br />

The .env should have DB_PORT=3306 and HTTP_PORT=8080

# Docker Desktop  
![image](https://github.com/RinzlerN26/Go_CRUD/assets/74294802/3d09e085-2979-46b4-9777-5e32de717219)

# GET
![image](https://github.com/RinzlerN26/Go_CRUD/assets/74294802/6f1531a7-e223-486a-ae02-0258941a22c1)

# POST 
![image](https://github.com/RinzlerN26/Go_CRUD/assets/74294802/7fa77d45-3c55-4c6a-bbb6-840ce849b534)

# PUT 
![image](https://github.com/RinzlerN26/Go_CRUD/assets/74294802/d2750172-c017-45e0-996a-b224a3aec02e)

# DELETE 
![image](https://github.com/RinzlerN26/Go_CRUD/assets/74294802/d78329e9-90b6-4ac8-bab2-d169879278c9)

# Testing

Testing done using main_test.go.

![image](https://github.com/RinzlerN26/Go_CRUD/assets/74294802/ae25a296-c493-45d6-9e0d-ecf94a89ef92)

![image](https://github.com/RinzlerN26/Go_CRUD/assets/74294802/d4bdf40e-21ad-42e1-b0bc-5394ab158e71)







