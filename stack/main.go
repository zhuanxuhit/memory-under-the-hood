package main

func main() {
	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// Pass the "value of" the count.
	increment(count)

	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	println("====================escape analysis========================")
	u1 := createUserV1()
	u2 := createUserV2()

	println("u1", &u1, "u2", &u2)
}

//go:noinline
func increment(inc int) {
	// Increment the "value of" inc.
	inc++
	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}


type user struct {
	name  string
	email string
}


//go:noinline
func createUserV1() user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V1", &u)
	return u
}

//go:noinline
func createUserV2() *user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V2", &u)
	return &u
}
