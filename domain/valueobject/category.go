package valueobject

type Category struct {
	Name string
	Type string
}

func NewCategory(name, categoryType string) (Category, error) {
	return Category{
		Name: name,
		Type: categoryType,
	}, nil
}
