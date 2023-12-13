package main

import "gofr.dev/pkg/gofr"

type cars struct {
	Id    int    `json:"id"`
	Cname string `json:"car_name"`
	Cstat string `json:"car_status"`
}

func main() {

	app := gofr.New()

	//Create Route
	app.POST("/cars/{car_name}/{car_status}", func(ctx *gofr.Context) (interface{}, error) {
		caname := ctx.PathParam("car_name")
		castat := ctx.PathParam("car_status")
		_, err := ctx.DB().ExecContext(ctx, "INSERT INTO cars (car_name,status) VALUES (?,?)", caname, castat)

		return nil, err
	})

	//Landing Page Route
	app.GET("/welcome", func(ctx *gofr.Context) (interface{}, error) {

		return "Welcome to CARS management GO_CRUD application!", nil
	})

	//Get Route
	app.GET("/get_cars", func(ctx *gofr.Context) (interface{}, error) {
		var car_d []cars

		// Getting the customer from the database using SQL
		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM cars")
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var car_i cars
			if err := rows.Scan(&car_i.Id, &car_i.Cname, &car_i.Cstat); err != nil {
				return nil, err
			}

			car_d = append(car_d, car_i)
		}

		return car_d, nil
	})

	//Update Route
	app.PUT("/cars/{car_id}/{car_status}", func(ctx *gofr.Context) (interface{}, error) {
		caid := ctx.PathParam("car_id")
		castat := ctx.PathParam("car_status")
		_, err := ctx.DB().ExecContext(ctx, "Update cars set status=? where id=?", castat, caid)

		return nil, err
	})

	//Delete route
	app.DELETE("/cars/delete", func(ctx *gofr.Context) (interface{}, error) {
		_, err := ctx.DB().ExecContext(ctx, "Delete from cars where status='RepairingDone'")

		return nil, err
	})

	app.Start()
}
