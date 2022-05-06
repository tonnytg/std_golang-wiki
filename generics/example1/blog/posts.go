package blog

func GetPosts() entities.Post {
	return []entities.Post {
		{
			ID: 1,
			Category: &entities.Category{ID: 1, Name: "Generics"},
			Title: "Aprendendo Generics",
		},
	}
}