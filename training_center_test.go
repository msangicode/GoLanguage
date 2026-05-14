package main

import "testing"

func TestNewTrainingCenterDefaultsToMB820(t *testing.T) {
	center := NewTrainingCenter("Contoso Learning")

	if center.Program != "MB820" {
		t.Fatalf("expected program MB820, got %s", center.Program)
	}
}

func TestAddCourseAndEnrollStudent(t *testing.T) {
	center := NewTrainingCenter("Contoso Learning")
	center.AddCourse("MB820-01", "Foundations", 2)

	if !center.EnrollStudent("MB820-01") {
		t.Fatal("expected first enrollment to succeed")
	}
	if !center.EnrollStudent("MB820-01") {
		t.Fatal("expected second enrollment to succeed")
	}
	if center.EnrollStudent("MB820-01") {
		t.Fatal("expected third enrollment to fail due to capacity")
	}

	if got := center.TotalEnrollments(); got != 2 {
		t.Fatalf("expected total enrollments 2, got %d", got)
	}
}

func TestAddCourseIgnoresDuplicates(t *testing.T) {
	center := NewTrainingCenter("Contoso Learning")
	center.AddCourse("MB820-01", "Foundations", 2)
	center.AddCourse("MB820-01", "Different title", 10)

	if got := len(center.Courses); got != 1 {
		t.Fatalf("expected only one course entry, got %d", got)
	}

	course := center.Courses["MB820-01"]
	if course.Title != "Foundations" || course.MaxStudents != 2 {
		t.Fatal("expected duplicate add to preserve original course data")
	}
}
