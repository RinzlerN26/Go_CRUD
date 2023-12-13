# Go_CRUD
A gofr based application implementing CRUD operations on MySQL based docker database image.

# Instructions to run
*Open terminal and run in the project directory.

```sh
go run main.go 
```

*Go to http://localhost:8080/welcome to see the landing page <br />
*Go to http://localhost:8080/get_cars to see information about the cars.<br />
*Go to http://localhost:8080/cars/delete to delete repaired cars.<br />
*Go to http://localhost:8080/cars/{car_id}/{car_status} replacing car_id and car_status with your values to update the database.<br />
*Go to http://localhost:8080/cars/{car_name}/{car_status} to add new cars to the database.<br />


# Landing Page
![image](https://github.com/RinzlerN26/Go_CRUD/assets/74294802/2cd7b1e9-dd21-4a36-a634-b0f7030670c0)

# Inserting values in database
![image](https://github.com/RinzlerN26/Go_CRUD/assets/74294802/2e15faa1-3d6b-4b45-92dd-99d776c55913)

# Docker mySQL image
![image](https://github.com/RinzlerN26/Go_CRUD/assets/74294802/3d09e085-2979-46b4-9777-5e32de717219)

# Postman POST requests
![image](https://github.com/RinzlerN26/Go_CRUD/assets/74294802/7fa77d45-3c55-4c6a-bbb6-840ce849b534)



