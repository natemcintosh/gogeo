package gogeo

import (
	"math"
	"testing"
)

func TestPointAngle(t *testing.T) {
	testCases := []struct {
		desc string
		in   Point
		out  float64
	}{
		{
			desc: "Point on the x-axis",
			in:   Point{0, 0},
			out:  0,
		},
		{
			desc: "Point on the y-axis",
			in:   Point{0, 1},
			out:  math.Pi / 2,
		},
		{
			desc: "Point negative x-axis",
			in:   Point{-1, 0},
			out:  math.Pi,
		},
		{
			desc: "Point negative y-axis",
			in:   Point{0, -1},
			out:  -math.Pi / 2,
		},
		{
			desc: "45 deg angle",
			in:   Point{1, 1},
			out:  math.Pi / 4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := tC.in.Angle(); got != tC.out {
				t.Errorf("Angle() = %v, want %v", got, tC.out)
			}
		})
	}
}

func TestPointRotate(t *testing.T) {
	testCases := []struct {
		desc  string
		in    Point
		angle float64
		out   Point
	}{
		{
			desc:  "rotate from x-axis by 45 deg",
			in:    Point{1, 0},
			angle: math.Pi / 4,
			out:   Point{math.Cos(math.Pi / 4), math.Sin(math.Pi / 4)},
		},
		{
			desc:  "rotate from x-axis 90 deg",
			in:    Point{1, 0},
			angle: math.Pi / 2,
			out:   Point{0, 1},
		},
		{
			desc:  "rotate from x-axis 180 deg",
			in:    Point{1, 0},
			angle: math.Pi,
			out:   Point{-1, 0},
		},
		{
			desc:  "rotate from x-axis 270 deg",
			in:    Point{1, 0},
			angle: math.Pi * 3 / 2,
			out:   Point{0, -1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := tC.in.Rotate(tC.angle); !got.AlmostEquals(tC.out) {
				t.Errorf("Rotate() = %v, want %v", got, tC.out)
			}
		})
	}
}

func TestLineSegmentAdd(t *testing.T) {
	testCases := []struct {
		desc string
		l    LineSegment
		p    Point
		out  LineSegment
	}{
		{
			desc: "x-axis line plus nothing",
			l:    LineSegment{Point{0, 0}, Point{1, 0}},
			p:    Point{0, 0},
			out:  LineSegment{Point{0, 0}, Point{1, 0}},
		},
		{
			desc: "x-axis line plus Point{1, 1}",
			l:    LineSegment{Point{0, 0}, Point{1, 0}},
			p:    Point{1, 1},
			out:  LineSegment{Point{1, 1}, Point{2, 1}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := tC.l.Plus(tC.p); !got.Equals(tC.out) {
				t.Errorf("Add() = %v, want %v", got, tC.out)
			}
		})
	}
}

func TestLineSegmentSubtract(t *testing.T) {
	testCases := []struct {
		desc string
		l    LineSegment
		p    Point
		out  LineSegment
	}{
		{
			desc: "x-axis line plus nothing",
			l:    LineSegment{Point{0, 0}, Point{1, 0}},
			p:    Point{0, 0},
			out:  LineSegment{Point{0, 0}, Point{1, 0}},
		},
		{
			desc: "x-axis line plus Point{1, 1}",
			l:    LineSegment{Point{0, 0}, Point{1, 0}},
			p:    Point{1, 1},
			out:  LineSegment{Point{-1, -1}, Point{0, -1}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := tC.l.Minus(tC.p); !got.Equals(tC.out) {
				t.Errorf("Add() = %v, want %v", got, tC.out)
			}
		})
	}
}

func TestLineSegmentAngle(t *testing.T) {
	testCases := []struct {
		desc string
		in   LineSegment
		out  float64
	}{
		{
			desc: "Horizontal Line",
			in:   LineSegment{Point{0, 0}, Point{1, 0}},
			out:  0,
		},
		{
			desc: "Vertical Line",
			in:   LineSegment{Point{0, 0}, Point{0, 1}},
			out:  math.Pi / 2,
		},
		{
			desc: "45 deg angle",
			in:   LineSegment{Point{1, 1}, Point{2, 2}},
			out:  math.Pi / 4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := tC.in.Angle(); got != tC.out {
				t.Errorf("Angle() = %v, want %v", got, tC.out)
			}
		})
	}
}

func TestOpenIntervalIsEmpty(t *testing.T) {
	testCases := []struct {
		desc string
		in   OpenInterval
		out  bool
	}{
		{
			desc: "both NaN",
			in:   OpenInterval{math.NaN(), math.NaN()},
			out:  true,
		},
		{
			desc: "first NaN",
			in:   OpenInterval{math.NaN(), 1},
			out:  true,
		},
		{
			desc: "second NaN",
			in:   OpenInterval{1, math.NaN()},
			out:  true,
		},
		{
			desc: "both are regular numbers",
			in:   OpenInterval{1, 2},
			out:  false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := tC.in.IsEmpty(); got != tC.out {
				t.Errorf("IsEmpty() = %v, want %v", got, tC.out)
			}
		})
	}
}

func TestOpenIntervalIntersection(t *testing.T) {
	testCases := []struct {
		desc string
		o1   OpenInterval
		o2   OpenInterval
		out  OpenInterval
	}{
		{
			desc: "No overlap",
			o1:   OpenInterval{1, 2},
			o2:   OpenInterval{3, 4},
			out:  OpenInterval{math.NaN(), math.NaN()},
		},
		{
			desc: "Some overlap",
			o1:   OpenInterval{1, 2},
			o2:   OpenInterval{1.5, 2.5},
			out:  OpenInterval{1.5, 2},
		},
		{
			desc: "Complete overlap",
			o1:   OpenInterval{1, 2},
			o2:   OpenInterval{1, 2},
			out:  OpenInterval{1, 2},
		},
		{
			desc: "Single number overlap",
			o1:   OpenInterval{1, 2},
			o2:   OpenInterval{2, 3},
			out:  OpenInterval{2, 2},
		},
		{
			desc: "Some more overlap",
			o1:   OpenInterval{-10, -5},
			o2:   OpenInterval{-7.3, -2},
			out:  OpenInterval{-7.3, -5},
		},
		{
			desc: "Both NaN",
			o1:   OpenInterval{math.NaN(), math.NaN()},
			o2:   OpenInterval{math.NaN(), math.NaN()},
			out:  OpenInterval{math.NaN(), math.NaN()},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if got := tC.o1.Intersection(tC.o2); !got.Equals(tC.out) {
				t.Errorf("Intersection() = %v, want %v", got, tC.out)
			}
		})
	}
}

func BenchmarkPointAngle(b *testing.B) {
	benchmarks := []struct {
		desc string
		in   Point
	}{
		{"positive x-axis", Point{1, 0}},
		{"positive y-axis", Point{0, 1}},
		{"negative y-axis", Point{-1, 0}},
		{"random point 1", Point{3.4, -2.3}},
		{"random point 2", Point{100.2, 7.6}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.in.Angle()
			}

		})
	}
}

func BenchmarkPointRotate(b *testing.B) {
	benchmarks := []struct {
		desc  string
		in    Point
		angle float64
	}{
		{"Rotate x-axis with no angle", Point{1, 0}, 0},
		{"Rotate x-axis with 45 deg angle", Point{1, 0}, math.Pi / 4},
		{"Rotate random point by 90 deg angle", Point{3.4, -2.3}, math.Pi / 2},
	}

	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.in.Rotate(bm.angle)
			}

		})
	}
}

func BenchmarkLineSegmentAngle(b *testing.B) {
	benchmarks := []struct {
		desc string
		in   LineSegment
	}{
		{"positive x-axis", LineSegment{Point{1, 0}, Point{2, 0}}},
		{"positive y-axis", LineSegment{Point{0, 1}, Point{0, 2}}},
		{"negative y-axis", LineSegment{Point{-1, 0}, Point{-2, 0}}},
		{"random point 1", LineSegment{Point{3.4, -2.3}, Point{100.2, 7.6}}},
		{"random point 2", LineSegment{Point{23.554, 3990.2}, Point{0, 5.45345}}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.in.Angle()
			}

		})
	}
}

func BenchmarkOpenIntervalIntersection(b *testing.B) {
	benchmarks := []struct {
		desc string
		o1   OpenInterval
		o2   OpenInterval
	}{
		{"No overlap", OpenInterval{1, 2}, OpenInterval{3, 4}},
		{"Some overlap", OpenInterval{1, 2}, OpenInterval{1.5, 2.5}},
		{"Complete overlap", OpenInterval{1, 2}, OpenInterval{1, 2}},
		{"Single number overlap", OpenInterval{1, 2}, OpenInterval{2, 3}},
		{"Some more overlap", OpenInterval{-10, -5}, OpenInterval{-7.3, -2}},
		{"Both NaN", OpenInterval{math.NaN(), math.NaN()}, OpenInterval{math.NaN(), math.NaN()}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.o1.Intersection(bm.o2)
			}

		})
	}
}

func BenchmarkOpenIntervalIsEmpty(b *testing.B) {
	benchmarks := []struct {
		desc string
		in   OpenInterval
	}{
		{"empty", OpenInterval{math.NaN(), math.NaN()}},
		{"first NaN", OpenInterval{math.NaN(), 1}},
		{"second NaN", OpenInterval{1, math.NaN()}},
		{"both are regular numbers", OpenInterval{1, 2}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.in.IsEmpty()
			}

		})
	}
}
