package schema

import (
	"time"

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
		field.Enum("role").Values("user", "admin").Default("user"),
		field.JSON("more_info", map[string]interface{}{}).Default(map[string]interface{}{}).Optional(),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
