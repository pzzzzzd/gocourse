package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/internal/models"
	"restapi/internal/repository/sqlconnect"
	"strconv"
	"strings"
	"sync"
)

var (
	teachers = make(map[int]models.Teacher)
	mutex    = &sync.Mutex{}
	nextID   = 1
)

// func init() {
// 	teachers[nextID] = models.Teacher{
// 		ID:        nextID,
// 		FirstName: "John",
// 		LastName:  "Abduroz",
// 		Class:     "11A",
// 		Subject:   "Math",
// 	}
// 	nextID++
// 	teachers[nextID] = models.Teacher{
// 		ID:        nextID,
// 		FirstName: "Simon",
// 		LastName:  "Dedov",
// 		Class:     "8B",
// 		Subject:   "Language",
// 	}
// 	nextID++
// 	teachers[nextID] = models.Teacher{
// 		ID:        nextID,
// 		FirstName: "Ne",
// 		LastName:  "Dedov",
// 		Class:     "9V",
// 		Subject:   "Bio",
// 	}
// 	nextID++
// }

func TeacherHendler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Method)
	switch r.Method {
	case http.MethodGet:
		getTeachersHandler(w, r)
	case http.MethodPost:
		addTeacherHandler(w, r)
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on Teachers Route"))
		fmt.Println("Hello PUT method on Teachers Route")
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on Teachers Route"))
		fmt.Println("Hello PATCH method on Teachers Route")
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on Teachers Route"))
		fmt.Println("Hello PATCH method on Teachers Route")
	}

}

func getTeachersHandler(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := strings.TrimSuffix(path, "/")
	fmt.Println(idStr)

	if idStr == "" {
		firstName := r.URL.Query().Get("first_name")
		lastName := r.URL.Query().Get("last_name")

		teacherList := make([]models.Teacher, 0, len(teachers))

		for _, teacher := range teachers {
			if (firstName == "" || teacher.FirstName == firstName) && (lastName == "" || teacher.LastName == lastName) {
				teacherList = append(teacherList, teacher)
			}
		}

		response := struct {
			Status string           `json:"status"`
			Count  int              `json:"count"`
			Data   []models.Teacher `json:"data"`
		}{
			Status: "success",
			Count:  len(teacherList),
			Data:   teacherList,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	teacher, exists := teachers[id]
	if !exists {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(teacher)
}

func addTeacherHandler(w http.ResponseWriter, r *http.Request) {

	// mutex.Lock()
	// defer mutex.Unlock()

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	var newTeachers []models.Teacher
	err = json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	// addedTeachers := make([]models.Teacher, len(newTeachers))
	// for i, newTeachers := range newTeachers {
	// 	newTeachers.ID = nextID
	// 	teachers[nextID] = newTeachers
	// 	addedTeachers[i] = newTeachers
	// 	nextID++
	// }

	stmt, err := db.Prepare("INSERT INTO teachers (first_name, last_name, email, class, subject) VALUES (?,?,?,?,?)")
	if err != nil {
		// log.Printf("Prepare error: %v", err)
		http.Error(w, fmt.Sprintf("Error in preparing SQL query: %v", err), http.StatusInternalServerError)
		return
	}

	addedTeachers := make([]models.Teacher, len(newTeachers))
	for i, newTeacher := range newTeachers {
		res, err := stmt.Exec(newTeacher.FirstName, newTeacher.LastName, newTeacher.Email, newTeacher.Class, newTeacher.Subject)
		if err != nil {
			http.Error(w, "Error inserting data into database", http.StatusInternalServerError)
			return
		}
		lastID, err := res.LastInsertId()
		if err != nil {
			http.Error(w, "Error getting last insert ID", http.StatusInternalServerError)
			return
		}
		newTeacher.ID = int(lastID)
		addedTeachers[i] = newTeacher
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Teacher `json:"data"`
	}{
		Status: "succes",
		Count:  len(addedTeachers),
		Data:   addedTeachers,
	}
	json.NewEncoder(w).Encode(response)

}
