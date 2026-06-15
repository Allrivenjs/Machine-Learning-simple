# Ejercicios: Etapa 2 — Programación Matemática en Go

Implementar todo desde cero. Sin gonum ni librerías externas.

---

## Vectores

### EJ-02-001: Tipo Vector básico
```go
// Implementar en vector.go
type Vector []float64

func (v Vector) Add(other Vector) Vector
func (v Vector) Sub(other Vector) Vector
func (v Vector) Scale(s float64) Vector
func (v Vector) Dot(other Vector) float64
func (v Vector) NormL1() float64
func (v Vector) NormL2() float64
func (v Vector) Normalize() Vector
```
Test: verificar con v₁=[3,4] que NormL2=5.0

### EJ-02-002: Operaciones adicionales
```go
func (v Vector) Mean() float64
func (v Vector) Sum() float64
func (v Vector) Max() float64
func (v Vector) Min() float64
func (v Vector) ArgMax() int
```

---

## Matrices

### EJ-02-003: Tipo Matrix básico
```go
type Matrix struct {
    Rows, Cols int
    Data       [][]float64
}

func NewMatrix(rows, cols int) Matrix
func (m Matrix) At(i, j int) float64
func (m Matrix) Set(i, j int, val float64)
func (m Matrix) Add(other Matrix) Matrix
func (m Matrix) Mul(other Matrix) Matrix
func (m Matrix) T() Matrix  // transpuesta
func (m Matrix) Trace() float64
```

### EJ-02-004: Determinante
Implementar det para 2×2 y extender a NxN usando expansión por cofactores:
```go
func (m Matrix) Det() float64
```
Test: det([[1,2],[3,4]]) = -2

### EJ-02-005: Inversa por Gauss-Jordan
```go
func (m Matrix) Inv() (Matrix, error)
```
Test: A × A.Inv() ≈ Identity (dentro de tolerancia 1e-9)

---

## Descomposiciones

### EJ-02-006: Eliminación Gaussiana
```go
// Resolver Ax = b
func SolveGaussian(A Matrix, b Vector) (Vector, error)
```

### EJ-02-007: LU Decomposition
```go
// Descomponer A = LU donde L es lower triangular y U upper triangular
func LUDecomposition(A Matrix) (L, U Matrix, err error)
```
Usar para resolver sistemas lineales más eficientemente.

### EJ-02-008: QR por Gram-Schmidt
```go
func QRDecomposition(A Matrix) (Q, R Matrix)
```
Verificar: Q es ortogonal (QᵀQ = I), R es upper triangular.

### EJ-02-009: Eigenvalor dominante
```go
// Método de potencia: encuentra el eigenvalor de mayor magnitud
func PowerIteration(A Matrix, maxIter int, tol float64) (eigenval float64, eigenvec Vector)
```

---

## Estadística

### EJ-02-010: Stats básicas
```go
func Mean(v Vector) float64
func Variance(v Vector) float64
func StdDev(v Vector) float64
func Covariance(x, y Vector) float64
func Pearson(x, y Vector) float64
```

### EJ-02-011: PCA manual
Usando las funciones anteriores, implementar PCA para 2 dimensiones:
1. Centrar los datos (restar media)
2. Calcular matriz de covarianza
3. Encontrar eigenvectores
4. Proyectar datos sobre componentes principales

Input: matrix N×2, Output: matrix N×2 transformada

---

## Proyecto integrador

Resolver el sistema lineal Ax = b para:
```
A = [[2, 1, -1],
     [-3, -1, 2],
     [-2, 1, 2]]

b = [8, -11, -3]
```
Solución esperada: x = [2, 3, -1]

Resolver usando:
1. Eliminación Gaussiana
2. LU Decomposition
3. Inversa: x = A⁻¹b

Comparar tiempos de ejecución de los 3 métodos.
