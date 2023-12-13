package main

import "gofr.dev/pkg/gofr"

type cars struct {
	Id    int    `json:"id"`
	Cname string `json:"car_name"`
	Cstat string `json:"car_status"`
}

func main() {
	// initialise gofr object
	app := gofr.New()

	app.GET("/welcome", func(ctx *gofr.Context) (interface{}, error) {

		return "Welcome to CARS management GO_CRUD application!", nil
	})

	app.POST("/cars/{car_name}/{car_status}", func(ctx *gofr.Context) (interface{}, error) {
		caname := ctx.PathParam("car_name")
		castat := ctx.PathParam("car_status")
		_, err := ctx.DB().ExecContext(ctx, "INSERT INTO cars (car_name,status) VALUES (?,?)", caname, castat)

		return nil, err
	})

	app.PUT("/cars/{car_id}/{car_status}", func(ctx *gofr.Context) (interface{}, error) {
		caid := ctx.PathParam("car_id")
		castat := ctx.PathParam("car_status")
		_, err := ctx.DB().ExecContext(ctx, "Update table cars set status=? where id=?", castat, caid)

		return nil, err
	})

	app.DELETE("/cars/{car_id}", func(ctx *gofr.Context) (interface{}, error) {
		caid := ctx.PathParam("car_id")
		_, err := ctx.DB().ExecContext(ctx, "Delete from cars where status='RepairingDone' and id=?", caid)

		return nil, err
	})

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

		// return the customer
		return car_d, nil
	})

	// Starts the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Start()
}
