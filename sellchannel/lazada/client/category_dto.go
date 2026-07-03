package client

// CategoryNode is one node of the Lazada category tree (/category/tree/get).
// Only leaf categories (Leaf == true) can be used as a product's PrimaryCategory.
type CategoryNode struct {
	CategoryID int64          `json:"category_id"`
	Name       string         `json:"name"`
	Leaf       bool           `json:"leaf"`
	Children   []CategoryNode `json:"children"`
}

// CategoryAttribute is one attribute a category accepts (/category/attributes/get).
// IsMandatory == 1 means the attribute must be supplied when creating a product.
type CategoryAttribute struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	IsMandatory int    `json:"is_mandatory"`
	InputType   string `json:"input_type"`
	Options     []struct {
		Name string `json:"name"`
	} `json:"options"`
}
