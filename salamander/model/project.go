package model

const (
	projectCreateSQL = `INSERT INTO project(id) VALUES(?)`
	projectLookupSQL = `SELECT * FROM project WHERE id=?`
	projectUpdateSQL = `UPDATE project SET(id=?)`
	projectDeleteSQL = `DELETE FROM project WHERE id=?`
)

// Project model
type Project struct {
	ID string `json:"project_id"`
}

// Scan method
func (p *Project) Scan(sc scanner) error {
	return sc.Scan(&p.ID)
}

// GetCreateSQL returns SQL query for creating a new database record.
func (p *Project) GetCreateSQL() string { return projectCreateSQL }

// GetCreateValues returns values list for placeholders in query returned GetCreateSQL().
func (p *Project) GetCreateValues() []interface{} {
	return []interface{}{p.ID}
}

// GetReadSQL returns SQL query for read from database record.
func (p *Project) GetReadSQL() string { return projectLookupSQL }

// GetReadValues returns values list for placeholders in query returned GetReadSQL().
func (p *Project) GetReadValues() []interface{} { return []interface{}{p.ID} }

// GetUpdateSQL returns SQL query for update a database record.
func (p *Project) GetUpdateSQL() string { return projectUpdateSQL }

// GetUpdateValues returns values list for placeholders in query returned GetUpdateSQL().
func (p *Project) GetUpdateValues() []interface{} {
	return []interface{}{p.ID}
}

// GetDeleteSQL returns SQL query for delete a database record.
func (p *Project) GetDeleteSQL() string { return projectDeleteSQL }

// GetDeleteValues returns values list for placeholders in query returned GetDeleteSQL().
func (p *Project) GetDeleteValues() []interface{} { return []interface{}{p.ID} }
