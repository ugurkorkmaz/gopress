package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
		field.JSON("more_info", map[string]interface{}{}).Default(map[string]interface{}{}).Optional(),
		field.Enum("status").Values("draft", "published", "archived").Default("draft"),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).Ref("posts").Unique(),
	}
}
