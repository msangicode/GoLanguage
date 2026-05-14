package main

import "sync"

type Course struct {
	Code          string
	Title         string
	MaxStudents   int
	EnrolledCount int
}

type TrainingCenter struct {
	Name    string
	Program string
	Courses map[string]*Course
	mu      sync.Mutex
}

func NewTrainingCenter(name string) *TrainingCenter {
	return &TrainingCenter{
		Name:    name,
		Program: "MB820",
		Courses: map[string]*Course{},
	}
}

func (tc *TrainingCenter) AddCourse(code string, title string, maxStudents int) {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	if _, exists := tc.Courses[code]; exists {
		return
	}

	tc.Courses[code] = &Course{
		Code:        code,
		Title:       title,
		MaxStudents: maxStudents,
	}
}

func (tc *TrainingCenter) EnrollStudent(code string) bool {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	course, exists := tc.Courses[code]
	if !exists || course.EnrolledCount >= course.MaxStudents {
		return false
	}

	course.EnrolledCount++
	return true
}

func (tc *TrainingCenter) TotalEnrollments() int {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	total := 0
	for _, course := range tc.Courses {
		total += course.EnrolledCount
	}
	return total
}
