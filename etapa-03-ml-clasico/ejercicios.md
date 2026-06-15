# Ejercicios: Etapa 3 — Machine Learning Clásico

Implementar cada algoritmo en Go. Luego validar contra Scikit-Learn en Python.

---

## Regresión Lineal

### EJ-03-001: Gradient Descent básico
```go
type LinearRegression struct {
    Weights Vector
    Bias    float64
    LR      float64
    Epochs  int
}

func (lr *LinearRegression) Fit(X Matrix, y Vector)
func (lr *LinearRegression) Predict(X Matrix) Vector
func (lr *LinearRegression) MSE(yTrue, yPred Vector) float64
```
Dataset de prueba: generar y = 2x + 1 + ruido gaussiano

### EJ-03-002: Normal Equation
```go
// Solución analítica: w = (XᵀX)⁻¹Xᵀy
func (lr *LinearRegression) FitAnalytic(X Matrix, y Vector)
```
Comparar pesos obtenidos vs Gradient Descent.

### EJ-03-003: Regularización Ridge
Modificar la pérdida: L = MSE + λ‖w‖²
Comparar pesos con y sin regularización en un dataset con multicolinealidad.

---

## Regresión Logística

### EJ-03-004: Clasificador binario
```go
type LogisticRegression struct {
    Weights Vector
    Bias    float64
    LR      float64
    Epochs  int
}

func sigmoid(z float64) float64
func (lr *LogisticRegression) Fit(X Matrix, y Vector)
func (lr *LogisticRegression) Predict(X Matrix) Vector      // probabilidades
func (lr *LogisticRegression) PredictClass(X Matrix) Vector // 0 o 1
```

### EJ-03-005: Métricas de clasificación
```go
func Accuracy(yTrue, yPred Vector) float64
func Precision(yTrue, yPred Vector) float64
func Recall(yTrue, yPred Vector) float64
func F1Score(yTrue, yPred Vector) float64
func ConfusionMatrix(yTrue, yPred Vector) Matrix
```
Aplicar sobre el clasificador de EJ-03-004.

---

## KNN

### EJ-03-006: KNN Clasificador
```go
type KNN struct {
    K    int
    X    Matrix
    Y    Vector
}

func euclideanDistance(a, b Vector) float64
func (k *KNN) Fit(X Matrix, y Vector)
func (k *KNN) Predict(X Matrix) Vector
```
Probar con K=1, 3, 5, 10 y ver cómo cambia accuracy.

### EJ-03-007: Efecto del K
Generar un dataset con 2 clases. Visualizar (o loggear) la frontera de decisión para K=1 vs K=20.

---

## Naive Bayes

### EJ-03-008: Gaussian Naive Bayes
```go
type GaussianNB struct {
    Means     Matrix  // media por clase y feature
    Variances Matrix  // varianza por clase y feature
    Priors    Vector  // P(clase)
}

func (gnb *GaussianNB) Fit(X Matrix, y Vector)
func (gnb *GaussianNB) Predict(X Matrix) Vector
func gaussianPDF(x, mean, variance float64) float64
```

---

## Árbol de Decisión

### EJ-03-009: Entropía e Information Gain
```go
func entropy(y Vector) float64
func informationGain(y Vector, leftY, rightY Vector) float64
func giniImpurity(y Vector) float64
```
Calcular IG para splits manuales en un dataset simple.

### EJ-03-010: Árbol completo
```go
type TreeNode struct {
    FeatureIdx int
    Threshold  float64
    Left       *TreeNode
    Right      *TreeNode
    Value      float64 // solo en hojas
}

type DecisionTree struct {
    Root     *TreeNode
    MaxDepth int
    MinSplit int
}

func (dt *DecisionTree) Fit(X Matrix, y Vector)
func (dt *DecisionTree) Predict(X Matrix) Vector
```

### EJ-03-011: Random Forest
```go
type RandomForest struct {
    Trees   []*DecisionTree
    NTrees  int
    MaxFeat int // features por árbol
}

func (rf *RandomForest) Fit(X Matrix, y Vector)
func (rf *RandomForest) Predict(X Matrix) Vector
```
Comparar accuracy de 1 árbol vs 10 vs 100 árboles.

---

## SVM

### EJ-03-012: SVM lineal (comprensión)
No implementar el optimizador QP completo (muy complejo). En cambio:
1. Implementar la función de decisión: f(x) = wᵀx + b
2. Implementar hinge loss: max(0, 1 - yᵢ(wᵀxᵢ + b))
3. Implementar gradient descent sobre hinge loss (soft-SVM)
4. Visualizar el margen para datos linealmente separables

---

## Conceptos transversales

### EJ-03-013: Cross-Validation
```go
func KFoldCV(model Model, X Matrix, y Vector, k int) []float64
func TrainTestSplit(X Matrix, y Vector, testRatio float64) (Matrix, Matrix, Vector, Vector)
```

### EJ-03-014: Comparación final
Usar el dataset Iris (o similar CSV):
- Entrenar todos los modelos implementados
- Evaluar con 5-fold CV
- Crear tabla comparativa de accuracy
