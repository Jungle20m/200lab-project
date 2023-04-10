package storage

import usermodel "200lab/internal/modules/user/model"

func (s *Storage) GetAllUsers() ([]usermodel.User, error) {
	var records []usermodel.User
	sql := `
			SELECT id, name, email, phone, create_time, update_time
			FROM users
			LIMIT 10
		   `
	err := s.db.Raw(sql).Find(&records).Error
	return records, err
}
