package main

import (
	"gofr.dev/pkg/gofr"
)

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	app := gofr.New()
	//get api to check working or not
	app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {

		return "Hello World!", nil
	})

	//post student
	app.POST("/student/{name}", func(ctx *gofr.Context) (interface{}, error) {
		name := ctx.PathParam("name")

		_, err := ctx.DB().ExecContext(ctx, "INSERT INTO students (name) VALUES (?)", name)

		return nil, err
	})

	//get all students
	app.GET("/students", func(ctx *gofr.Context) (interface{}, error) {
		var students []Student

		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM students")
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var student Student
			if err := rows.Scan(&student.ID, &student.Name); err != nil {
				return nil, err
			}

			students = append(students, student)
		}

		// return the customer
		return students, nil
	})

	app.Start()
}
