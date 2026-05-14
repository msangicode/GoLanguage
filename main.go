package main

import "fmt"

func main() {
	center := NewTrainingCenter("MB820 Training Center")
	center.AddCourse("MB820-01", "Finance and Operations Foundations", 20)
	center.AddCourse("MB820-02", "Business Process Mapping", 15)

	center.EnrollStudent("MB820-01")
	center.EnrollStudent("MB820-01")
	center.EnrollStudent("MB820-02")

	fmt.Printf(
		"%s (%s): %d courses, %d enrolled students\n",
		center.Name,
		center.Program,
		len(center.Courses),
		center.TotalEnrollments(),
	)
}
