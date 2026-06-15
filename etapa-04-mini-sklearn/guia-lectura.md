# Etapa 4: Mini Scikit-Learn en Go ("golearn")

Duración: 1 mes | Prerequisito: Etapas 1, 2, 3

---

## Objetivo

Unificar todos los algoritmos de Etapa 3 bajo una interfaz común. El resultado es una librería Go usable: `golearn`.

---

## Recursos
- Código fuente de Scikit-Learn (GitHub) — para ver cómo diseñan la interfaz
- **Clean Code** (Robert Martin) — para diseño de APIs
- Documentación de Go interfaces

---

## Diseño de la interfaz

```go
// Model es la interfaz central
type Model interface {
    Fit(X Matrix, y Vector) error
    Predict(X Matrix) (Vector, error)
}

// Scorer agrega evaluación
type Scorer interface {
    Model
    Score(X Matrix, y Vector) float64
}

// Pipeline encadena transformaciones
type Pipeline struct {
    Steps []Step
}
```

---

## Temas a implementar

### Core
- [ ] Interfaz `Model` con `Fit` y `Predict`
- [ ] Interfaz `Transformer` con `FitTransform`
- [ ] Interfaz `Scorer`

### Modelos (refactorizar desde Etapa 3)
- [ ] `LinearRegression` — normal equation + gradient descent
- [ ] `LogisticRegression` — binaria y multiclase (softmax)
- [ ] `KNN` — clasificación y regresión
- [ ] `GaussianNB`
- [ ] `DecisionTree` — clasificación y regresión
- [ ] `RandomForest`

### Transformers
- [ ] `StandardScaler` — media 0, std 1
- [ ] `MinMaxScaler` — escala [0, 1]
- [ ] `LabelEncoder` — strings a enteros
- [ ] `OneHotEncoder`

### Model Selection
- [ ] `TrainTestSplit`
- [ ] `KFold`
- [ ] `CrossValidate`
- [ ] `GridSearch` — búsqueda de hiperparámetros

### Métricas
- [ ] `Accuracy`, `Precision`, `Recall`, `F1`
- [ ] `MSE`, `RMSE`, `MAE`, `R2`
- [ ] `ConfusionMatrix`

### Data Loading
- [ ] `LoadCSV(path string) (Matrix, Vector, error)`
- [ ] Datasets builtin: Iris, Boston, Breast Cancer

---

## Estructura de paquetes

```
golearn/
├── linalg/       (de Etapa 2)
├── models/
│   ├── linear.go
│   ├── logistic.go
│   ├── knn.go
│   ├── naive_bayes.go
│   ├── tree.go
│   └── forest.go
├── preprocessing/
│   ├── scalers.go
│   └── encoders.go
├── model_selection/
│   ├── split.go
│   └── cross_val.go
├── metrics/
│   ├── classification.go
│   └── regression.go
├── datasets/
│   └── iris.go
└── pipeline.go
```

---

## Señales de que dominas la etapa

- Puedes entrenar cualquier modelo con el mismo código cliente:
  ```go
  model.Fit(X_train, y_train)
  preds := model.Predict(X_test)
  score := model.Score(X_test, y_test)
  ```
- Pipeline funciona encadenando scaler + model
- GridSearch encuentra el mejor K para KNN automáticamente
