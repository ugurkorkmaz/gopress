package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty().MinLen(3).MaxLen(128),
		field.String("email").NotEmpty().MinLen(3).MaxLen(128),
		field.String("password").NotEmpty().MinLen(8),
		field.Enum("role").Values("user", "admin"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
