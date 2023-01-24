# fiber-mongo-api

Build a REST API with Golang and MongoDB - Fiber Version

This repository shows the source code for building a user management application with Golang using the Fiber framework and MongoDB.

Article Link

[Article Link](https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-fiber-version-4la0)
https://blog.loginradius.com/engineering/mongodb-as-datasource-in-golang/




	newPage1 := make([]models.Pages, 0)

	//newPage := models.Page[];
	for _, b := range page {
		// Add a new Person object to the array
		p.Id = primitive.NewObjectID()
		p.Name = b.Name
		p.Pid = b.Pid
		p.Ptitle = b.Ptitle
		p.Pdesc = b.Pdesc
		p.Pkeyword = b.Pkeyword
		p.Pimg = b.Pimg
		p.PimgAlt = b.PimgAlt
		p.Pauthor = b.Pauthor
		newPage1 = append(newPage1, p)
		//var a = append(cleaned, p)
	}
	fmt.Print(newPage1)