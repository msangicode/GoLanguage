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
