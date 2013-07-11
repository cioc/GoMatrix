package matrix

type Vector []float64

func dot(v1 Vector, v2 Vector) (float64) {
  o := float64(0)
  for key := range v1 {
    o += v1[key] * v2[key] 
  }
  return o
}
