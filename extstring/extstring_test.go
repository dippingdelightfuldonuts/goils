package extstring

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExtString_Camelcase(t *testing.T) {
	tests := []struct {
		name string
		e    ExtString
		want ExtString
	}{
		{
			name: "converts camelcase",
			e:    "HappyMealIsForMe",
			want: "HappyMealIsForMe",
		},
		{
			name: "converts snakecase",
			e:    "happy_meal_is_for_me",
			want: "HappyMealIsForMe",
		},
		{
			name: "converts spaces",
			e:    "happy meal is for me",
			want: "HappyMealIsForMe",
		},
		{
			name: "converts titlecase",
			e:    "Happy Meal Is For Me",
			want: "HappyMealIsForMe",
		},
		{
			name: "converts upcase",
			e:    "HAPPY MEAL IS FOR ME",
			want: "HAPPYMEALISFORME",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Convey("camlecase", t, func() {
				So(tt.e.Camelcase(), ShouldResemble, tt.want)
			})
		})
	}
}

func TestExtString_Downcase(t *testing.T) {
	tests := []struct {
		name string
		e    ExtString
		want ExtString
	}{
		{
			name: "converts camelcase",
			e:    "HappyMealIsForMe",
			want: "happymealisforme",
		},
		{
			name: "converts snakecase",
			e:    "happy_meal_is_for_me",
			want: "happy_meal_is_for_me",
		},
		{
			name: "converts spaces",
			e:    "happy meal is for me",
			want: "happy meal is for me",
		},
		{
			name: "converts titlecase",
			e:    "Happy Meal Is For Me",
			want: "happy meal is for me",
		},
		{
			name: "converts upcase",
			e:    "HAPPY MEAL IS FOR ME",
			want: "happy meal is for me",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Convey("downcase", t, func() {
				So(tt.e.Downcase(), ShouldResemble, tt.want)
			})
		})
	}
}

func TestExtString_Pluralize(t *testing.T) {
	tests := []struct {
		name string
		e    ExtString
		want ExtString
	}{
		{
			name: "simple plural",
			e:    "book",
			want: "books",
		},
		{
			name: "end in y plural",
			e:    "candy",
			want: "candies",
		},
		{
			name: "end in se plural",
			e:    "case",
			want: "cases",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Convey("simple plural", t, func() {
				So(tt.e.Pluralize(), ShouldResemble, tt.want)
			})
		})
	}
}

func TestExtString_Snakecase(t *testing.T) {
	tests := []struct {
		name string
		e    ExtString
		want ExtString
	}{
		{
			name: "converts camelcase",
			e:    "HappyMealIsForMe",
			want: "happy_meal_is_for_me",
		},
		{
			name: "converts snakecase",
			e:    "happy_meal_is_for_me",
			want: "happy_meal_is_for_me",
		},
		{
			name: "converts spaces",
			e:    "happy meal is for me",
			want: "happy_meal_is_for_me",
		},
		{
			name: "converts titlecase",
			e:    "Happy Meal Is For Me",
			want: "happy_meal_is_for_me",
		},
		{
			name: "converts upcase",
			e:    "HAPPY MEAL IS FOR ME",
			want: "happy_meal_is_for_me",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Convey("snake case", t, func() {
				So(tt.e.Snakecase(), ShouldResemble, tt.want)
			})
		})
	}
}

func TestExtString_Titlecase(t *testing.T) {
	tests := []struct {
		name string
		e    ExtString
		want ExtString
	}{
		{
			name: "converts camelcase",
			e:    "HappyMealIsForMe",
			want: "HappyMealIsForMe",
		},
		{
			name: "converts snakecase",
			e:    "happy_meal_is_for_me",
			want: "Happy_meal_is_for_me",
		},
		{
			name: "converts spaces",
			e:    "happy meal is for me",
			want: "Happy Meal Is For Me",
		},
		{
			name: "converts titlecase",
			e:    "Happy Meal Is For Me",
			want: "Happy Meal Is For Me",
		},
		{
			name: "converts upcase",
			e:    "HAPPY MEAL IS FOR ME",
			want: "HAPPY MEAL IS FOR ME",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Convey("Titlecase", t, func() {
				So(tt.e.Titlecase(), ShouldResemble, tt.want)
			})
		})
	}
}

func TestExtString_Upcase(t *testing.T) {
	tests := []struct {
		name string
		e    ExtString
		want ExtString
	}{
		{
			name: "converts camelcase",
			e:    "HappyMealIsForMe",
			want: "HAPPYMEALISFORME",
		},
		{
			name: "converts snakecase",
			e:    "happy_meal_is_for_me",
			want: "HAPPY_MEAL_IS_FOR_ME",
		},
		{
			name: "converts spaces",
			e:    "happy meal is for me",
			want: "HAPPY MEAL IS FOR ME",
		},
		{
			name: "converts upcase",
			e:    "HAPPY MEAL IS FOR ME",
			want: "HAPPY MEAL IS FOR ME",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Convey("Upcase", t, func() {
				So(tt.e.Upcase(), ShouldResemble, tt.want)
			})
		})
	}
}
