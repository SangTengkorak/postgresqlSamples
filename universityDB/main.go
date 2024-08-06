package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"math/rand"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 50; i++ {
		firstName := fmt.Sprintf("FirstName%d", i)
		lastName := fmt.Sprintf("LastName%d", i)
		dateOfBirth := fmt.Sprintf("%d-%02d-%02d", rand.Intn(10)+1990, rand.Intn(12)+1, rand.Intn(28)+1)
		gender := "M"
		if i%2 == 0 {
			gender = "F"
		}
		major := fmt.Sprintf("Major%d", rand.Intn(10)+1)

		_, err := conn.Exec(context.Background(), "INSERT INTO student (first_name, last_name, date_of_birth, gender, major) VALUES ($1, $2, $3, $4, $5)",
			firstName, lastName, dateOfBirth, gender, major)
		if err != nil {
			log.Fatalf("Unable to insert student: %v\n", err)
		}

		courseName := fmt.Sprintf("Course%d", i)
		department := fmt.Sprintf("Department%d", rand.Intn(5)+1)

		_, err = conn.Exec(context.Background(), "INSERT INTO course (course_name, department) VALUES ($1, $2)", courseName, department)
		if err != nil {
			log.Fatalf("Unable to insert course: %v\n", err)
		}

		courseId := i
		sectionNumber := rand.Intn(3) + 1
		semester := "Fall"
		if i%2 == 0 {
			semester = "Spring"
		}
		year := rand.Intn(10) + 2010

		_, err = conn.Exec(context.Background(), "INSERT INTO section (course_id, section_number, semester, year) VALUES ($1, $2, $3, $4)",
			courseId, sectionNumber, semester, year)
		if err != nil {
			log.Fatalf("Unable to insert section: %v\n", err)
		}

		studentId := i
		sectionId := i
		grade := string([]rune{'A', 'B', 'C', 'D', 'F'}[rand.Intn(5)])

		_, err = conn.Exec(context.Background(), "INSERT INTO grade (student_id, section_id, grade) VALUES ($1, $2, $3)",
			studentId, sectionId, grade)
		if err != nil {
			log.Fatalf("Unable to insert grade: %v\n", err)
		}

		tutorFirstName := fmt.Sprintf("TutorFirstName%d", i)
		tutorLastName := fmt.Sprintf("TutorLastName%d", i)
		tutorDepartment := fmt.Sprintf("Department%d", rand.Intn(5)+1)

		_, err = conn.Exec(context.Background(), "INSERT INTO tutor (first_name, last_name, department) VALUES ($1, $2, $3)",
			tutorFirstName, tutorLastName, tutorDepartment)
		if err != nil {
			log.Fatalf("Unable to insert tutor: %v\n", err)
		}
	}

	fmt.Println("Sample data inserted successfully!")
}
