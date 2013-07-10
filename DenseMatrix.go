package matrix

type DenseMatrix struct {
  rows int
  cols int
  vals []float64
}

func NewDenseMatrix(rows int, cols int) (*DenseMatrix) {
  return &DenseMatrix{rows, cols, make([]float64, rows * cols)}
}

func (d *DenseMatrix) Size() (rows, cols int) {
  return d.rows, d.cols
}

func (d *DenseMatrix) Get(row, col int) (val float64) {
  return d.vals[(d.cols * row) + col];
}

func (d *DenseMatrix) GetRow(row int) (vals []float64) {
  return d.vals[(d.cols * row):((d.cols * row) + d.cols)]
}

func (d *DenseMatrix) GetCol(col int) (vals []float64) {
  o := make([]float64, d.rows)
  for i := 0; i < d.rows; i++ {
    o[i] = d.vals[(d.cols * i) + col];
  }
  return o
}

func (d *DenseMatrix) Set(row int, col int, val float64) {
  d.vals[(d.cols * row) + col] = val;
}

func (d *DenseMatrix) SetRow(row int, val []float64) {
  for i:= 0; i < d.cols; i++ {
    d.vals[(d.cols * row) + i] = val[i]
  }
}

func (d *DenseMatrix) SetCol(col int, val[]float64) {
  for i := 0; i < d.rows; i++ {
    d.vals[(d.cols * i) + col] = val[i]
  }
}

func (d *DenseMatrix) Scale(scale float64) {
  for i:= 0; i < d.rows * d.cols; i++ {
    d.vals[i] *= scale
  }
}

func (d *DenseMatrix) Transpose() {
  valsNew := make([]float64, d.cols * d.rows)
  for i:= 0; i < d.rows; i++ {
    for j:= 0; j < d.cols; j++ {
      valsNew[(d.rows * j) + i] = d.vals[(d.cols * i) + j]
    }
  }
  d.vals = valsNew
  d.rows, d.cols = d.cols, d.rows
}

//Row operations
func (d *DenseMatrix) RowSwap(row1, row2 int) {
  for i := 0; i < d.cols; i++ {
    d.vals[(d.cols * row1) + i], d.vals[(d.cols * row2) + i] = d.vals[(d.cols * row2) + i], d.vals[(d.cols * row1) + i]
  }
}

func (d *DenseMatrix) RowScale(row int, scale float64) {
  for i := 0; i < d.cols; i++ {
    d.vals[(d.cols * row) + i] *= scale
  }
}

func (d *DenseMatrix) RowAdd(row int, vals []float64) {
  for i := 0; i < d.cols; i++ {
    d.vals[(d.cols * row) + i] += vals[i]
  }
}

//Arithmetic
