package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Student struct {
	ID    int    `json: "id"`
	Name  string `json: "name"`
	Age   int    `json: "age"`
	Grade string `json: "grade"`
}

var students []Student

// ambil semua data
func getAllStudents(c echo.Context) error {
	return c.JSON(http.StatusOK, students)
}

// ambil data berdasarkan id
func getStudentByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, student := range students {
		if student.ID == id {
			return c.JSON(http.StatusOK, student)
		}
	}
	return c.JSON(http.StatusNotFound, "Student not found")
}

// membuat data baru
func createNewStudent(c echo.Context) error {
	student := new(Student)
	if err := c.Bind(student); err != nil {
		return err
	}

	student.ID = len(students) + 1
	students = append(students, *student)

	return c.JSON(http.StatusCreated, student)
}

func updateStudentByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i := range students {
		if students[i].ID == id {
			updatedStudent := new(Student)
			if err := c.Bind(updatedStudent); err != nil {
				return err
			}

			students[i].Name = updatedStudent.Name
			students[i].Age = updatedStudent.Age
			students[i].Grade = updatedStudent.Grade

			return c.JSON(http.StatusOK, students[i])

		}
	}
	return c.JSON(http.StatusNotFound, "Student not found")
}

func deleteStudentByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	for i := range students {
		if students[i].ID == id {
			students = append(students[:i], students[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, "Student not found")
}

func main() {
	e := echo.New()

	e.POST("/students", createNewStudent)
	e.GET("/students", getAllStudents)
	e.GET("/students/:id", getStudentByID)
	e.PUT("/students/:id", updateStudentByID)
	e.DELETE("/students/:id", deleteStudentByID)

	e.Start("localhost:8080")
}
