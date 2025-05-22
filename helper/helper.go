package helper

var (
	version     = "1.0.0" // cannot be accessed globally because using lowecase
	Application = "golang"
)

func sayGoodBye(name string) string { // cannot be accessed globally, but still can be used inside this package
	return "Good Bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}
