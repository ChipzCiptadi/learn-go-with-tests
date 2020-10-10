package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	t.Run("with test cases", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				"struct with one string field",
				struct {
					Name string
				}{"Chipz"},
				[]string{"Chipz"},
			},
			{
				"struct with two string fields",
				struct {
					Name string
					City string
				}{"Chipz", "Jakarta"},
				[]string{"Chipz", "Jakarta"},
			},
			{
				"struct with non string fields",
				struct {
					Name string
					Age  int
				}{"Chipz", 10},
				[]string{"Chipz"},
			},
			{
				"Nested fields",
				Person{
					"Chipz",
					Profile{10, "Jakarta"}},
				[]string{"Chipz", "Jakarta"},
			},
			{
				"Pointers to things",
				&Person{
					"Chipz",
					Profile{10, "Jakarta"},
				},
				[]string{"Chipz", "Jakarta"},
			},
			{
				"Slices",
				[]Profile{
					{10, "Jakarta"},
					{15, "Pontianak"},
				},
				[]string{"Jakarta", "Pontianak"},
			},
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var got []string

				walk(test.Input, func(input string) {
					got = append(got, input)
				})

				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("got %v, want %v", got, test.ExpectedCalls)
				}
			})
		}
	})

	t.Run("Maps", func(t *testing.T) {
		aMap := map[string]string{
			"key":   "Hello",
			"value": "World",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Hello")
		assertContains(t, got, "World")
	})

	t.Run("Channel", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{10, "Jakarta"}
			aChannel <- Profile{15, "Pontianak"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Jakarta", "Pontianak"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{10, "Jakarta"}, Profile{15, "Pontianak"}
		}

		var got []string
		want := []string{"Jakarta", "Pontianak"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()

	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
			break
		}
	}

	if contains == false {
		t.Errorf("expected %v to contains %q but it didn't", haystack, needle)
	}
}
