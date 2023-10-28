package diy

import "gorm.io/gen"

type Querier interface {
	// SELECT * FROM @@table WHERE id=@id
	GetByID(id int) (gen.T, error) // GetByID query data by id and return it as *struct*

	// SELECT * FROM @@table WHERE id IN @ids
	MGet(ids ...string) ([]*gen.T, error)

	// QueryWith
	//SELECT * FROM @@table
	//  {{if p != nil}}
	//      {{if p.ID > 0}}
	//          WHERE id=@p.ID
	//      {{else if p.Name != ""}}
	//          WHERE name=@p.Name
	//      {{end}}
	//  {{end}}
	QueryWith(p *gen.T) (gen.T, error)
}
