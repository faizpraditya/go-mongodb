package main

import "go-mongodb/delivery"

func main() {
	var server delivery.Routes
	server.StartGin()

	// mdb, err := db.InitResource()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// ctx, cancel := utils.InitContext()
	// defer cancel()
	// defer func() {
	// 	if err = mdb.Db.Client().Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// if err := mdb.Db.Client().Ping(ctx, readpref.Primary()); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Successfully to connect and pinged")

	// student := repository.NewStudentRepository(mdb)

	// getAll, err := student.GetAll()
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	for _, doc := range getAll {
	// 		fmt.Println(doc)
	// 	}
	// }
	// fmt.Println()

	// getOne, err := student.GetOneByUsername("Faiz")
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	fmt.Println(*getOne)
	// }
	// fmt.Println()

	// newStudent := model.Student{
	// 	Id:       primitive.NewObjectID(),
	// 	Name:     "Jude Bellingham",
	// 	Gender:   "M",
	// 	Age:      18,
	// 	JoinDate: time.Now(),
	// 	IdCard:   "2218",
	// 	Senior:   false,
	// }
	// createOne, err := student.CreateOne(newStudent)
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	fmt.Println(*createOne)
	// }
	// fmt.Println()

	// pagination, err := student.GetWithPage(2, 2)
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	for _, doc := range pagination {
	// 		fmt.Println(doc)
	// 	}
	// }
	// fmt.Println()

}
