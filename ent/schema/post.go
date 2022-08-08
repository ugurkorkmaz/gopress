package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New),
		field.String("title").NotEmpty().MinLen(3).MaxLen(255),
		field.String("slug").NotEmpty().MinLen(3).MaxLen(255),
		field.Text("content").NotEmpty().MinLen(3).MaxLen(65535),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return nil
}
