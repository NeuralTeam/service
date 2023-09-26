package main

func main() {
	if err := Service.Run(); err != nil {
		_ = Service.Error(err)
	}
}
