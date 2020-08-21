package extstring

import (
	"strings"

	pluralize "github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

// ExtString extends string so that it has
// Camelcase()
// Downcase()
// Pluralize()
// SnakeCase()
// Upcase()
type ExtString string

// Camelcase formats string to Camelcase
// (i.e. "happy_meal_for_me" becomes "HappyMealForMe")
func (e ExtString) Camelcase() ExtString {
	return ExtString(strcase.ToCamel(string(e)))
}

// Downcase formats string to all lower
// (i.e. "Happy Meal For Me" becomes "happy meal for me")
func (e ExtString) Downcase() ExtString {
	return ExtString(strings.ToLower(string(e)))
}

// Pluralize returns the plural form of a word
// (i.e. "dog" becomes "dogs")
func (e ExtString) Pluralize() ExtString {
	return ExtString(
		pluralizeClient().Plural(string(e)),
	)
}

// Snakecase returns the snake case form
// (i.e. "HappyMealForMe" becomes "happy_meal_for_me")
func (e ExtString) Snakecase() ExtString {
	return ExtString(strcase.ToSnake(string(e)))
}

// Titlecase returns titleized format of the string
// (i.e. "happy meal for me" becomes "Happy Meal For Me")
func (e ExtString) Titlecase() ExtString {
	return ExtString(strings.Title(string(e)))
}

// Upcase returns uppercase of the string
// (i.e. "happy meal for me" becomes "HAPPY MEAL FOR ME")
func (e ExtString) Upcase() ExtString {
	return ExtString(strings.ToUpper(string(e)))
}

func pluralizeClient() *pluralize.Client {
	return pluralize.NewClient()
}
