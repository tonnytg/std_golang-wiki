package main

func main() {
	fmt.Println("Lear Generics")

	cc := cache.New[entities.Category]{}
	categories := blog.GetCategories()
	for _, value := range categories {
		cc.Set(value.Slug, value)
	}

	cp := cache.Cache[entities.Post]{}
	posts := blog.GetPosts()
	for _, value := range posts {
		cp.Set(value.Slug, value)
	}


	fmt.Println(cc.Get("generics"))
	fmt.Println("---")
	fmt.Println(categories)
	fmt.Println(posts)

}