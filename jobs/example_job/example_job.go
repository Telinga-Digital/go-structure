package jobs

import "fmt"

var (
	Channel = "example_job"
)

func Handle(payload interface{}) {
	fmt.Println(payload)
}
