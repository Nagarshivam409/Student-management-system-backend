package app

import (
	"backend/pkg/models"
	"backend/pkg/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func createStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&student)

	repository.GetDB().Create(&student)

	json.NewEncoder(w).Encode(student)
}
func createTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher models.Teacher
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&teacher)

	repository.GetDB().Create(&teacher)

	json.NewEncoder(w).Encode(teacher)
}
func PunchinStudent(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	year := t.Year()
	day := t.Day()
	month := t.Month()

	var student models.Student
	var student1 models.Student
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&student1)
	fmt.Println(r.Body)
	fmt.Println(student1)
	repository.GetDB().Where("s_id = ?", student1.SID).First(&student)
	if student.Name == "" {
		json.NewEncoder(w).Encode("Student Not found")
	} else {
		var attendance models.AttendanceStudent
		var updatedattendance models.AttendanceStudent
		repository.GetDB().Where("s_id = ? AND year = ? AND month = ? AND day = ?", student.SID, year, month, day).Last(&attendance)
		attendance.SID = student.SID
		attendance.Year = year
		attendance.Month = month
		attendance.Day = day
		attendance.Class = student.Class
		updatedattendance.SID = student.SID
		updatedattendance.Year = year
		updatedattendance.Month = month
		updatedattendance.Day = day
		updatedattendance.Class = student.Class

		if attendance.PunchIn != "" && attendance.PunchOut == "" {
			json.NewEncoder(w).Encode("Already Punched In")
		} else {
			repository.GetDB().Where("s_id = ? AND year = ? AND month = ? AND day = ?", student.SID, year, month, day).Delete(models.AttendanceStudent{})
			updatedattendance.PunchIn = time.Now().Format("15:04:05")

			repository.GetDB().Create(&updatedattendance)
			json.NewEncoder(w).Encode("Successfully Punched In")
		}

	}

}
func PunchoutStudent(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	year := t.Year()
	day := t.Day()
	month := t.Month()

	var student models.Student
	var student1 models.Student
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&student1)
	repository.GetDB().Where("s_id = ?", student1.SID).First(&student)
	if student.Name == "" {
		json.NewEncoder(w).Encode("Student Not found")
	} else {
		var attendance models.AttendanceStudent
		var updatedattendance models.AttendanceStudent
		repository.GetDB().Where("s_id = ? AND year = ? AND month = ? AND day = ?", student.SID, year, month, day).Last(&attendance)
		attendance.SID = student.SID
		attendance.Year = year
		attendance.Month = month
		attendance.Day = day
		attendance.Class = student.Class
		updatedattendance.SID = student.SID
		updatedattendance.Year = year
		updatedattendance.Month = month
		updatedattendance.Day = day
		updatedattendance.Class = student.Class

		if attendance.PunchIn != "" && attendance.PunchOut == "" {
			var attendance2 models.AttendanceStudent
			repository.GetDB().Where("s_id = ? AND year = ? AND month = ? AND day = ?", student.SID, year, month, day).Last(&attendance2)
			repository.GetDB().Where("s_id = ? AND year = ? AND month = ? AND day = ? AND punchin = ?", student.SID, year, month, day, attendance2.PunchIn).Delete(models.AttendanceStudent{})
			updatedattendance.PunchIn = attendance2.PunchIn
			updatedattendance.PunchOut = time.Now().Format("15:04:05")

			repository.GetDB().Create(&updatedattendance)
			json.NewEncoder(w).Encode("Successfully Punched Out")

		} else {
			json.NewEncoder(w).Encode("You need to Punch in First")
		}

	}
}
func PunchinTeacher(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	year := t.Year()
	day := t.Day()
	month := t.Month()

	var teacher models.Teacher
	var teacher1 models.Teacher
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&teacher1)
	repository.GetDB().Where("t_id = ?", teacher1.TID).First(&teacher)
	if teacher.Name == "" {
		json.NewEncoder(w).Encode("Teacher Not found")
	} else {
		var attendance models.AttendanceTeacher
		var updatedattendance models.AttendanceTeacher

		repository.GetDB().Where("t_id = ? AND year = ? AND month = ? AND day = ?", teacher.TID, year, month, day).Last(&attendance)
		attendance.TID = teacher.TID
		attendance.Year = year
		attendance.Month = month
		attendance.Day = day

		updatedattendance.TID = teacher.TID
		updatedattendance.Year = year
		updatedattendance.Month = month
		updatedattendance.Day = day

		if attendance.PunchIn != "" && attendance.PunchOut == "" {
			json.NewEncoder(w).Encode("Already Punched In")
		} else {
			repository.GetDB().Where("s_id = ? AND year = ? AND month = ? AND day = ?", teacher.TID, year, month, day).Delete(models.AttendanceTeacher{})
			updatedattendance.PunchIn = time.Now().Format("15:04:05")

			repository.GetDB().Create(&updatedattendance)
			json.NewEncoder(w).Encode("Successfully Punched In")
		}

	}
}
func PunchoutTeacher(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	year := t.Year()
	day := t.Day()
	month := t.Month()

	var teacher models.Teacher
	var teacher1 models.Teacher
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&teacher1)
	repository.GetDB().Where("t_id = ?", teacher1.TID).First(&teacher)
	if teacher.Name == "" {
		json.NewEncoder(w).Encode("Student Not found")
	} else {
		var attendance models.AttendanceTeacher
		var updatedattendance models.AttendanceTeacher
		repository.GetDB().Where("t_id = ? AND year = ? AND month = ? AND day = ?", teacher.TID, year, month, day).Last(&attendance)
		attendance.TID = teacher.TID
		attendance.Year = year
		attendance.Month = month
		attendance.Day = day

		updatedattendance.TID = teacher.TID
		updatedattendance.Year = year
		updatedattendance.Month = month
		updatedattendance.Day = day

		if attendance.PunchIn != "" && attendance.PunchOut == "" {
			var attendance1 models.AttendanceTeacher
			repository.GetDB().Where("t_id = ? AND year = ? AND month = ? AND day = ?", teacher.TID, year, month, day).Last(&attendance1)
			repository.GetDB().Where("t_id = ? AND year = ? AND month = ? AND day = ? AND punchin = ?", teacher.TID, year, month, day, attendance1.PunchIn).Delete(models.AttendanceStudent{})
			updatedattendance.PunchIn = attendance1.PunchIn
			updatedattendance.PunchOut = time.Now().Format("15:04:05")

			repository.GetDB().Create(&updatedattendance)
			json.NewEncoder(w).Encode("Successfully Punched Out")

		} else {
			json.NewEncoder(w).Encode("You need to Punch in First")
		}

	}
}
func getClassAttendance(w http.ResponseWriter, r *http.Request) {
	var student []models.AttendanceStudent

	params := mux.Vars(r)
	repository.GetDB().Where("class = ?", params["class"]).Find(&student)
	if len(student) == 0 {
		json.NewEncoder(w).Encode("No students in the class")
	} else {
		json.NewEncoder(w).Encode(student)
	}
}
func getStudentAttendance(w http.ResponseWriter, r *http.Request) {
	var student []models.AttendanceStudent

	params := mux.Vars(r)
	repository.GetDB().Where("s_id = ? AND month= ? AND year = ?", params["id"], params["month"], params["year"]).Find(&student)
	if len(student) == 0 {
		json.NewEncoder(w).Encode("No students Attendance Found")
	} else {
		json.NewEncoder(w).Encode(student)
	}
}
func getTeacherAttendance(w http.ResponseWriter, r *http.Request) {
	var teacher []models.AttendanceTeacher

	params := mux.Vars(r)
	repository.GetDB().Where("t_id = ? AND month= ? AND year = ?", params["id"], params["month"], params["year"]).Find(&teacher)
	if len(teacher) == 0 {
		json.NewEncoder(w).Encode("No Teacher Attendance Found")
	} else {
		json.NewEncoder(w).Encode(teacher)
	}
}
