package models

import (
	"log"
	"reakgo/utility"
)

// To do list data base
type TodoList struct {
	Id   int32  `json:"id"`
	Task string `json:"task"`
	Date string `json:"date"`
}

func (myTask TodoList) Add(task, date string) error {
	log.Println("Attempting to insert data:", task, date)
	result, err := utility.Db.Exec("INSERT INTO todo (task, date) VALUES (?, ?)", task, date)
	if err != nil {
		log.Println("Error inserting data:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error fetching rows affected:", err)
		return err
	}

	log.Println("Data inserted successfully, rows affected:", rowsAffected)
	return nil
}

func (myTask TodoList) View() ([]TodoList, error) {
	var resultSet []TodoList

	rows, err := utility.Db.Query("SELECT * FROM todo")
	if err != nil {
		log.Println(err)
	} else {
		defer rows.Close()
		var singleRow TodoList

		for rows.Next() {
			err = rows.Scan(&singleRow.Id, &singleRow.Task, &singleRow.Date)
			if err != nil {
				log.Println(err)
			} else {
				resultSet = append(resultSet, singleRow)
			}
		}
	}
	return resultSet, err
}
func DeleteTask(id int) error {
	_, err := utility.Db.Exec("DELETE FROM todo WHERE id = ?", id)
	return err
}

func UpdateTask(task TodoList) error {
	_, err := utility.Db.Exec("UPDATE todo SET task = ?, date = ? WHERE id = ?", task.Task, task.Date, task.Id)
	return err
}

// package models

// import (
// 	"log"
// 	"reakgo/utility"
// )

// type TodoList struct {
// 	ID   int    `json:"id"`
// 	Task string `json:"task"`
// 	Date string `json:"date"`
// }

// func (myTask TodoList) Add(task string, date string) {
// 	log.Println("insert data")
// 	utility.Db.MustExec("INSERT INTO todo (task, date) VALUES (?, ?)", task, date)
// }
// func (myTask TodoList) View() ([]TodoList, error) {
// 	var resultSet []TodoList

// 	rows, err := utility.Db.Query("SELECT * FROM todo")
// 	if err != nil {
// 		log.Println(err)
// 	} else {
// 		defer rows.Close()

// 		for rows.Next() {
// 			var singleRow TodoList
// 			err = rows.Scan(&singleRow.ID, &singleRow.Task, &singleRow.Date)
// 			if err != nil {
// 				log.Println(err)
// 			} else {
// 				resultSet = append(resultSet, singleRow)
// 			}
// 		}
// 	}
// 	return resultSet, err
// }

// func GetAllTasks() ([]TodoList, error) {
// 	rows, err := utility.Db.Query("SELECT id, title, date FROM tasks")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var tasks []TodoList
// 	for rows.Next() {
// 		var task TodoList
// 		err := rows.Scan(&task.ID, &task.Task, &task.Date)
// 		if err != nil {
// 			return nil, err
// 		}
// 		tasks = append(tasks, task)
// 	}

// 	return tasks, nil
// }

//	func AddTask(task TodoList) error {
//		log.Println("insert data")
//		_, err := utility.Db.Exec("INSERT INTO tasks (title, date) VALUES (?, ?)", task.Task, task.Date)
//		return err
//	}
// func AddTask(task TodoList) error {
// 	log.Println("Attempting to insert data:", task.Task, task.Date)
// 	result, err := utility.Db.Exec("INSERT INTO tasks (title, date) VALUES (?, ?)", task.Task, task.Date)
// 	if err != nil {
// 		log.Println("Error inserting data:", err)
// 		return err
// 	}

// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		log.Println("Error fetching rows affected:", err)
// 		return err
// 	}

// 	log.Println("Data inserted successfully, rows affected:", rowsAffected)
// 	return nil
// }
