package util

import (
	"fmt"
	_ "gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
	"math"
	"time"
)

type wordTfIdf struct {
	nworld string
	value  float64
}

type wordTfIdfs []wordTfIdf

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func (us wordTfIdfs) Len() int {
	return len(us)
}
func (us wordTfIdfs) Less(i, j int) bool {
	return us[i].value > us[j].value
}
func (us wordTfIdfs) Swap(i, j int) {
	us[i], us[j] = us[j], us[i]
}

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}

var (
	v2i   = make(map[string]int, 20)
	i2v   = make(map[int]string, 20)
	idf   *mat.Dense
	tf    *mat.Dense
	tfIdf *mat.Dense
)

func FeatureSelect(listWords [][]string) *mat.Dense {
	docFrequency := make(map[string]float64, 0)
	for _, wordList := range listWords {
		for _, v := range wordList {
			docFrequency[v] += 1
		}
	}
	vocab := getKeys(docFrequency)
	for i, v := range vocab {
		v2i[v] = i
		i2v[i] = v
	}
	idf = mat.NewDense(len(i2v), 1, nil)
	for i := 0; i < len(i2v); i++ {
		dCount := 0
		for _, d := range listWords {
			if IsContain(d, i2v[i]) {
				dCount++
			}
		}
		idf.Set(i, 0, float64(dCount))
	}
	r, _ := idf.Caps()
	for i := 0; i < r; i++ {
		idf.Set(i, 0, 1+math.Log(float64(len(listWords))/(1+idf.At(i, 0))))
	}

	tf = mat.NewDense(len(vocab), len(listWords), nil)
	for i, wordList := range listWords {
		dFrequency := make(map[string]float64, 0)
		for _, v := range wordList {
			dFrequency[v] += 1
		}
		dKeys := getKeys(dFrequency)
		dMax := 0.0
		for _, v := range dKeys {
			if dFrequency[v] > dMax {
				dMax = dFrequency[v]
			}
		}
		for _, v := range dKeys {
			tf.Set(v2i[v], i, dFrequency[v]/dMax)
		}
	}

	r, c := tf.Caps()
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			tf.Set(i, j, math.Log(1+tf.At(i, j)))
		}
	}

	tfIdf = mat.NewDense(len(vocab), len(listWords), nil)
	r, c = tfIdf.Caps()
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			tfIdf.Set(i, j, tf.At(i, j)*idf.At(i, 0))
		}
	}
	return tfIdf
}

func Load() []string {
	//数据
	slice := []string{
		"I am nigger",
		"fuck you nigger",
		"nigger you are",
		"small nigger",
		"fuck me nigger",
	}
	return slice
}


func getKeys(m map[string]float64) []string {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

func getKeysInt(m map[string]int) []string {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func matSquare(m *mat.Dense) *mat.Dense {
	r, c := m.Caps()
	result := mat.NewDense(r, c, nil)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			num := m.At(i, j)
			result.Set(i, j, num*num)
		}
	}
	return result
}

func matSum(m *mat.Dense) *mat.Dense {
	r, c := m.Caps()
	result := mat.NewDense(1, c, nil)
	for j := 0; j < c; j++ {
		sum := 0.0
		for i := 0; i < r; i++ {
			sum += m.At(i, j)
		}
		result.Set(0, j, sum)
	}
	return result
}

func matSqrt(m *mat.Dense) *mat.Dense {
	r, c := m.Caps()
	result := mat.NewDense(r, c, nil)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			result.Set(i, j, math.Sqrt(m.At(i, j)))
		}
	}
	return result
}

func transpose(m *mat.Dense) *mat.Dense {
	r, c := m.Caps()
	result := mat.NewDense(c, r, nil)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			result.Set(j, i, m.At(i, j))
		}
	}
	return result
}

// columns of m and n should be the same
func divElem(m *mat.Dense, n *mat.Dense) *mat.Dense {
	r1, c1 := m.Caps()
	r2, c2 := n.Caps()
	if c1 != c2 || r2 != 1 {
		return nil
	}
	result := mat.NewDense(r1, c1, nil)
	for i := 0; i < r1; i++ {
		for j := 0; j < c1; j++ {
			result.Set(i, j, m.At(i, j)/n.At(0, j))
		}
	}
	return result
}

func denseDot(m *mat.Dense, n *mat.Dense) *mat.Dense {
	r1, c1 := m.Caps()
	r2, c2 := n.Caps()
	if c1 != r2 {
		return nil
	}
	result := mat.NewDense(r1, c2, nil)
	for i := 0; i < r1; i++ {
		for j := 0; j < c2; j++ {
			temp := 0.0
			for k := 0; k < c1; k++ {
				temp += m.At(i, k) * n.At(k, j)
			}
			result.Set(i, j, temp)
		}
	}
	return result
}

func cosineSimilarity(q *mat.Dense, tfIdf *mat.Dense) *mat.Dense {
	unitsQ := divElem(q, matSqrt(matSum(matSquare(q))))
	temp := divElem(tfIdf, matSqrt(matSum(matSquare(tfIdf))))
	unitsDs := transpose(temp)
	result := denseDot(unitsDs, unitsQ)

	r, c := result.Caps()
	// Flatten
	similarity := mat.NewDense(r*c, 1, nil)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			similarity.Set(i*c+j, 0, result.At(i, j))
		}
	}
	return similarity
}

// columns of m and n should be the same
func appendHorizontal(m *mat.Dense, n *mat.Dense) *mat.Dense {
	r1, c1 := m.Caps()
	r2, c2 := n.Caps()
	if c1 != c2 {
		return nil
	}
	result := mat.NewDense(r1+r2, c1, nil)
	i := 0
	for ; i < r1; i++ {
		for j := 0; j < c1; j++ {
			result.Set(i, j, m.At(i, j))
		}
	}
	for ; i < r1+r2; i++ {
		for j := 0; j < c2; j++ {
			result.Set(i, j, n.At(i-r1, j))
		}
	}
	return result
}

// The shape of two dense should be the same
func mulElem(m *mat.Dense, n *mat.Dense) *mat.Dense {
	r1, c1 := m.Caps()
	r2, c2 := n.Caps()
	if r1 != r2 || c1 != c2 {
		return nil
	}
	result := mat.NewDense(r1, c1, nil)
	for i := 0; i < r1; i++ {
		for j := 0; j < c1; j++ {
			result.Set(i, j, m.At(i, j)*n.At(i, j))
		}
	}
	return result
}

func DocsScore(qWords []string) *mat.Dense {

	unknownV := 0
	for _, v := range qWords {
		v2iKeys := getKeysInt(v2i)
		if !IsContain(v2iKeys, v) {
			v2i[v] = len(v2i)
			i2v[len(v2i)-1] = v
			unknownV += 1
		}
	}
	var _idf, _tf_idf *mat.Dense
	if unknownV > 0 {
		temp := mat.NewDense(unknownV, 1, nil)
		_idf = appendHorizontal(idf, temp)
		_, c := tfIdf.Caps()
		temp2 := mat.NewDense(unknownV, c, nil)
		_tf_idf = appendHorizontal(tfIdf, temp2)
	} else {
		_idf, _tf_idf = idf, tfIdf
	}
	//fmt.Println(_idf)
	r, _ := _idf.Caps()
	qTf := mat.NewDense(r, 1, nil)
	qFrequency := make(map[string]int, 0)
	for _, w := range qWords {
		qFrequency[w] += 1
	}
	wordsKey := getKeysInt(qFrequency)
	for _, v := range wordsKey {
		qTf.Set(v2i[v], 0, float64(qFrequency[v]))
	}
	var qVec, qScore *mat.Dense
	qVec = mulElem(qTf, _idf)
	qScore = cosineSimilarity(qVec, _tf_idf)
	return qScore
}

func Dense2slice(m *mat.Dense) []float64 {
	r, c := m.Caps()
	slice := make([]float64, r*c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			slice[i*c+j] = m.At(i, j)
		}
	}
	return slice
}

func getKeywords(n int) {
	r, _ := tfIdf.Caps()
	for i := 0; i < 3; i++ {
		col := make([]float64, r)
		for j := 0; j < r; j++ {
			col[j] = tfIdf.At(j, i)
		}
		inds := make([]int, r)
		floats.Argsort(col, inds)
		idx := inds[len(inds)-n:]
		vec := make([]string, n)
		for j := 0; j < n; j++ {
			vec[j] = i2v[idx[j]]
		}
		fmt.Printf("doc%v, top%v keywords %v\n", i, n, vec)
	}
}

func Reverse(arr *[]int) {
	length := len(*arr)
	var temp int
	for i := 0; i < length/2; i++ {
		temp = (*arr)[i]
		(*arr)[i] = (*arr)[length-1-i]
		(*arr)[length-1-i] = temp
	}
}

