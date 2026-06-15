# Ejercicios: Etapa 4 — Mini Scikit-Learn en Go

---

## Interfaz y diseño

### EJ-04-001: Interfaz Model
Definir y testear la interfaz base:
```go
type Model interface {
    Fit(X Matrix, y Vector) error
    Predict(X Matrix) (Vector, error)
}
```
Verificar que LinearRegression, LogisticRegression, KNN implementan la interfaz sin errores de compilación.

### EJ-04-002: Transformer
```go
type Transformer interface {
    Fit(X Matrix) error
    Transform(X Matrix) (Matrix, error)
    FitTransform(X Matrix) (Matrix, error)
}
```
Implementar `StandardScaler` que cumpla esta interfaz.

---

## Preprocessing

### EJ-04-003: StandardScaler
```go
type StandardScaler struct {
    Means  Vector
    Stds   Vector
}
```
Test: después de transformar, cada feature debe tener media ≈ 0 y std ≈ 1.

### EJ-04-004: MinMaxScaler
```go
type MinMaxScaler struct {
    Mins Vector
    Maxs Vector
}
```
Test: todos los valores en [0, 1].

### EJ-04-005: OneHotEncoder
Convertir vector de categorías [0, 2, 1, 0] en matrix de indicadores.

---

## Pipeline

### EJ-04-006: Pipeline básico
```go
type Pipeline struct {
    Scaler Transformer
    Model  Model
}

func (p *Pipeline) Fit(X Matrix, y Vector) error
func (p *Pipeline) Predict(X Matrix) (Vector, error)
```
Usar: StandardScaler → LogisticRegression

---

## Model Selection

### EJ-04-007: KFold
```go
func KFoldSplit(n, k int) [][]int // retorna índices por fold
```

### EJ-04-008: CrossValidate
```go
func CrossValidate(model Model, X Matrix, y Vector, k int, metric func(Vector, Vector) float64) float64
```

### EJ-04-009: GridSearch
```go
type Param struct {
    Name  string
    Value interface{}
}

func GridSearch(
    modelFn func([]Param) Model,
    paramGrid [][]Param,
    X Matrix, y Vector,
    k int,
) (bestParams []Param, bestScore float64)
```
Usar para encontrar mejor K en KNN (rango 1-20).

---

## Datasets

### EJ-04-010: CSV Loader
```go
func LoadCSV(path string, targetCol int) (Matrix, Vector, error)
```

### EJ-04-011: Dataset Iris builtin
Hardcodear (o embed) el dataset Iris en Go. Retornar Matrix 150×4 y Vector 150.

---

## Proyecto final de etapa

Reproducir este pipeline de Scikit-Learn en Go:
```python
from sklearn.pipeline import Pipeline
from sklearn.preprocessing import StandardScaler
from sklearn.model_selection import cross_val_score
from sklearn.linear_model import LogisticRegression

pipe = Pipeline([('scaler', StandardScaler()), ('clf', LogisticRegression())])
scores = cross_val_score(pipe, X, y, cv=5)
print(scores.mean())  # debe ser > 0.95 en Iris
```

Tu versión Go debe producir accuracy comparable.
