// Пакет workspace_lib представляет утилиту Add
package workspace_lib

import "golang.org/x/exp/constraints"

// AddNums принимает два экземпляра типа int или float и возвращает один исходного типа - сумму
// Более подробные сведения можно найти на сайте [MathIsFun].
//
// [MathIsFun]: https://www.mathsisfun.com/numbers/addition.html
func AddNums[T Number](a, b T) T {
	return a + b
}

// Number интерфейс реализует два импортируемых типа Integer and Float
type Number interface {
	constraints.Integer | constraints.Float
}
