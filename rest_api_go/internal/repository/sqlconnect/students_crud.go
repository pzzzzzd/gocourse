package sqlconnect

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"restapi/internal/models"
	"restapi/pkg/utils"
	"strconv"
)

func GetStudentsDbHandler(students []models.Student, r *http.Request) ([]models.Student, error) {
	db, err := ConnectDb()
	if err != nil {
		return nil, utils.ErrorHandler(err, "error retrieving data")
	}

	defer db.Close()

	query := "SELECT id, first_name, last_name, email, class FROM students WHERE 1=1"
	var args []interface{}

	query, args = utils.AddFilters(r, query, args)

	query = utils.AddSorting(r, query)

	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
		return nil, utils.ErrorHandler(err, "error retrieving data")
	}
	defer rows.Close()

	for rows.Next() {
		var student models.Student
		err = rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Class)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func GetStudentByID(id int) (models.Student, error) {
	db, err := ConnectDb()
	if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "error retrieving data")
	}

	defer db.Close()

	var student models.Student
	err = db.QueryRow("SELECT id, first_name, last_name, email, class FROM students WHERE id = ?", id).Scan(&student.ID, &student.FirstName, &student.LastName,
		&student.Email, &student.Class)
	if err == sql.ErrNoRows {
		return models.Student{}, utils.ErrorHandler(err, "error retrieving data")
	} else if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "error retrieving data")
	}
	return student, nil
}

func AddStudentsDbHandler(newStudents []models.Student) ([]models.Student, error) {
	db, err := ConnectDb()
	if err != nil {
		return nil, utils.ErrorHandler(err, "error adding data 1")
	}

	defer db.Close()

	// stmt, err := db.Prepare("INSERT INTO students (first_name, last_name, email, class) VALUES (?,?,?,?,?)")
	stmt, err := db.Prepare(utils.GenerateInsertQuery("students", models.Student{}))
	if err != nil {
		return nil, utils.ErrorHandler(err, "error adding data 2")
	}

	addedStudents := make([]models.Student, len(newStudents))
	for i, newStudent := range newStudents {
		// res, err := stmt.Exec(newStudent.FirstName, newStudent.LastName, newStudent.Email, newStudent.Class, newStudent.Subject)
		values := utils.GetStructValues(newStudent)
		res, err := stmt.Exec(values...)
		if err != nil {
			return nil, utils.ErrorHandler(err, "error adding data 3")
		}
		lastID, err := res.LastInsertId()
		if err != nil {
			return nil, utils.ErrorHandler(err, "error adding data 4")
		}
		newStudent.ID = int(lastID)
		addedStudents[i] = newStudent
	}
	return addedStudents, nil
}

func UpdateStudent(id int, updatedStudents models.Student) (models.Student, error) {
	db, err := ConnectDb()
	if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "error updating data")
	}
	defer db.Close()

	var existingStudent models.Student
	err = db.QueryRow("SELECT id, first_name, last_name, email, class FROM students WHERE id = ?", id).Scan(&existingStudent.ID, &existingStudent.FirstName, &existingStudent.LastName,
		&existingStudent.Email, &existingStudent.Class)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Student{}, utils.ErrorHandler(err, "error updating data")
		}
		return models.Student{}, utils.ErrorHandler(err, "error updating data")
	}

	updatedStudents.ID = existingStudent.ID
	_, err = db.Exec("UPDATE students SET first_name = ?, last_name = ?, email = ?, class = ? WHERE id = ?", updatedStudents.FirstName, updatedStudents.LastName, updatedStudents.Email,
		updatedStudents.Class, updatedStudents.ID)
	if err == sql.ErrNoRows {
		return models.Student{}, utils.ErrorHandler(err, "error updating data")
	}
	return updatedStudents, nil
}

func PatchStudents(updates []map[string]interface{}) error {
	db, err := ConnectDb()
	if err != nil {
		return utils.ErrorHandler(err, "Invalid ID")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return utils.ErrorHandler(err, "Invalid ID")
	}

	for _, update := range updates {
		idStr, ok := update["id"].(string)
		if !ok {
			tx.Rollback()
			return utils.ErrorHandler(err, "Invalid ID")
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			tx.Rollback()
			return utils.ErrorHandler(err, "Invalid ID")
		}

		var studentFromDb models.Student
		err = db.QueryRow("SELECT id, first_name, last_name, email, class FROM students WHERE id = ?", id).Scan(&studentFromDb.ID, &studentFromDb.FirstName,
			&studentFromDb.LastName, &studentFromDb.Email, &studentFromDb.Class)

		if err != nil {
			log.Println("ID:", id)
			log.Printf("Type: %T", id)
			log.Println(err)
			tx.Rollback()
			if err == sql.ErrNoRows {
				return utils.ErrorHandler(err, "student not found")
			}
			return utils.ErrorHandler(err, "Invalid ID")
		}

		studentVal := reflect.ValueOf(&studentFromDb).Elem()
		studentType := studentVal.Type()
		for k, v := range update {
			if k == "id" {
				continue
			}
			for i := 0; i < studentVal.NumField(); i++ {

				field := studentType.Field(i)
				if field.Tag.Get("json") == k+",omitempty" {
					fieldVal := studentVal.Field(i)
					if fieldVal.CanSet() {
						val := reflect.ValueOf(v)
						if val.Type().ConvertibleTo(fieldVal.Type()) {
							fieldVal.Set(val.Convert(fieldVal.Type()))
						} else {
							tx.Rollback()
							log.Printf("Cant convert %v to %v\n", val.Type(), fieldVal.Type())
							return utils.ErrorHandler(err, "Invalid ID")
						}

					}
					break
				}
			}
		}
		_, err = tx.Exec("UPDATE students SET first_name = ?, last_name = ?, email = ?, class = ? WHERE id = ?", studentFromDb.FirstName, studentFromDb.LastName, studentFromDb.Email,
			studentFromDb.Class, studentFromDb.ID)
		if err != nil {
			tx.Rollback()
			return utils.ErrorHandler(err, "error updating data")
		}
	}

	err = tx.Commit()
	if err != nil {
		return utils.ErrorHandler(err, "error updating data")
	}
	return nil
}

func PatchOneStudent(id int, updates map[string]interface{}) (models.Student, error) {
	db, err := ConnectDb()
	if err != nil {
		log.Println(err)
		return models.Student{}, utils.ErrorHandler(err, "error updating data")
	}
	defer db.Close()

	var existingStudent models.Student
	err = db.QueryRow("SELECT id, first_name, last_name, email, class FROM students WHERE id = ?", id).Scan(&existingStudent.ID, &existingStudent.FirstName, &existingStudent.LastName,
		&existingStudent.Email, &existingStudent.Class)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return models.Student{}, utils.ErrorHandler(err, "student not found")
		}
		return models.Student{}, utils.ErrorHandler(err, "error updating data")
	}

	studentVal := reflect.ValueOf(&existingStudent).Elem()
	studentType := studentVal.Type()

	for k, v := range updates {
		for i := 0; i < studentVal.NumField(); i++ {
			field := studentType.Field(i)
			if field.Tag.Get("json") == k+",omitempty" {
				if studentVal.Field(i).CanSet() {
					fieldVal := studentVal.Field(i)
					fmt.Println("fieldVal", fieldVal)
					fmt.Println("studentVal.Field(i).Type()", studentVal.Field(i).Type())
					fmt.Println("reflect.ValueOf(v)", reflect.ValueOf(v))
					fieldVal.Set(reflect.ValueOf(v).Convert(studentVal.Field(i).Type()))
				}
			}
		}
	}

	_, err = db.Exec("UPDATE students SET first_name = ?, last_name = ?, email = ?, class = ? WHERE id = ?", existingStudent.FirstName, existingStudent.LastName, existingStudent.Email,
		existingStudent.Class, existingStudent.ID)
	if err == sql.ErrNoRows {
		log.Println(err)

		return models.Student{}, utils.ErrorHandler(err, "error updating data")
	}
	return existingStudent, nil
}

func DeleteOneStudent(id int) error {
	db, err := ConnectDb()
	if err != nil {
		log.Println(err)
		return utils.ErrorHandler(err, "error deleting student")
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM students WHERE id = ?", id)
	if err != nil {
		return utils.ErrorHandler(err, "error deleting student")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return utils.ErrorHandler(err, "error deleting student")
	}

	if rowsAffected == 0 {
		return utils.ErrorHandler(err, "student not found")
	}
	return nil
}

func DeleteStudents(ids []int) ([]int, error) {
	db, err := ConnectDb()
	if err != nil {
		log.Println(err)
		return nil, utils.ErrorHandler(err, "error deleting student")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return nil, utils.ErrorHandler(err, "error deleting student")
	}

	stmt, err := tx.Prepare("DELETE FROM students WHERE id = ?")
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, utils.ErrorHandler(err, "error deleting student")
	}
	defer stmt.Close()

	deletedIds := []int{}
	for _, id := range ids {
		result, err := stmt.Exec(id)
		if err != nil {
			tx.Rollback()
			log.Println(err)

			return nil, utils.ErrorHandler(err, "error deleting student")
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			tx.Rollback()

			return nil, utils.ErrorHandler(err, "error deleting student")
		}
		if rowsAffected > 0 {
			deletedIds = append(deletedIds, id)
		}
		if rowsAffected < 1 {
			tx.Rollback()

			return nil, utils.ErrorHandler(err, fmt.Sprintf("ID %d not found", id))
		}

	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)

		return nil, utils.ErrorHandler(err, "error deleting student")
	}

	if len(deletedIds) < 1 {

		return nil, utils.ErrorHandler(err, "IDs do not found")
	}
	return deletedIds, nil
}
