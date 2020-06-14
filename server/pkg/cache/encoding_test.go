package cache

import "testing"

func BenchmarkMarshal(b *testing.B) {
	a := make([]int, 0, 400)
	for i := 0; i < 400; i++ {
		a = append(a, i)
	}
	jsonEncoding := JSONEncoding{}

	for n := 0; n < b.N; n++ {
		_, err := jsonEncoding.Marshal(a)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	a := make([]int, 0, 400)
	for i := 0; i < 400; i++ {
		a = append(a, i)
	}
	jsonEncoding := JSONEncoding{}
	data, err := jsonEncoding.Marshal(a)
	if err != nil {
		b.Error(err)
	}

	var result []int
	for n := 0; n < b.N; n++ {
		err = jsonEncoding.UnMarshal(data, &result)
		if err != nil {
			b.Error(err)
		}
	}
}
