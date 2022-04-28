package sql

import (
	"github.com/ranefattesingh/task-management-app/server/data"
	"gorm.io/gorm"
)

func (s *sql) AddTask(t *data.Task) (*data.Task, error) {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		tx.Create(t)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return t, err
}

func (s *sql) GetTasks() (data.Tasks, error) {

	tl := data.Tasks{}
	err := s.db.Transaction(func(tx *gorm.DB) error {
		tx.Find(&tl)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return tl, nil
}

func (s *sql) UpdateTask(id int, t *data.Task) (*data.Task, error) {
	d := &data.Task{}
	err := s.db.Transaction(func(tx *gorm.DB) error {
		tx.First(d, "ID = ?", id)

		if d == nil {
			return data.ErrTaskNotFound
		}

		d.Text = t.Text
		d.Reminder = t.Reminder
		d.Day = t.Day

		tx.Updates(d)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return d, nil

}

func (s *sql) DeleteTask(id int) error {
	d := &data.Task{}
	err := s.db.Transaction(func(tx *gorm.DB) error {
		tx.First(d, "ID = ?", id)

		if d == nil {
			return data.ErrTaskNotFound
		}
		tx.Delete(d)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *sql) GetTask(id int) (*data.Task, error) {
	d := &data.Task{}
	err := s.db.Transaction(func(tx *gorm.DB) error {
		tx.First(d, "ID = ?", id)

		if d == nil {
			return data.ErrTaskNotFound
		}
		tx.Delete(d)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return d, nil
}
