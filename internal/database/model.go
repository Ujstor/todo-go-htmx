package database

type Todo struct {
	Id uint64 `json:"id"`
	Todo string `json:"todo"`
	Done bool `json:"done"`
}

func (s *service) CreateTodo(todo string) error {

	statment := `INSERT INTO todos(todo, done) values($1, $2);`
	_, err := s.db.Query(statment, todo, false)
	return err
}

func (s *service) GetAllTodos() ([]Todo, error) {
	var todos []Todo
	statement := `SELECT id, todo, done FROM todos;`

	rows, err := s.db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t Todo
		err := rows.Scan(&t.Id, &t.Todo, &t.Done)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *service) GetTodo(id uint64) (Todo, error) {
    statement := `SELECT todo, done FROM todos WHERE id=$1;`

    var todo Todo
    todo.Id = id

    row := s.db.QueryRow(statement, id)
    err := row.Scan(&todo.Todo, &todo.Done)
    if err != nil {
        return todo, err
    }

    return todo, nil
}

func (s *service) MarkDone(id uint64) error {
    statement := `UPDATE todos SET done = NOT done WHERE id=$1;`
    _, err := s.db.Exec(statement, id)
    return err
}

func (s *service) Delete(id uint64) error {
    statement := `DELETE FROM todos WHERE id=$1;`
    _, err := s.db.Exec(statement, id)
    return err
}
