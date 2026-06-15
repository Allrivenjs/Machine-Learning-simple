# Ejercicios: Etapa 5 — Python Científico

---

## NumPy

### EJ-05-001: Reproducir Etapa 2 en NumPy
Implementar todas las operaciones de tu `matrix.go` usando NumPy:
```python
import numpy as np

A = np.array([[1, 2], [3, 4]])
B = np.array([[5, 6], [7, 8]])

# Multiplicación
C = A @ B  # o np.matmul(A, B)

# Transpuesta
At = A.T

# Determinante
det_A = np.linalg.det(A)

# Inversa
A_inv = np.linalg.inv(A)

# Eigenvalores
eigenvals, eigenvecs = np.linalg.eig(A)

# SVD
U, S, Vt = np.linalg.svd(A)
```
Comparar resultados con tu implementación Go.

### EJ-05-002: Broadcasting
Sin usar loops, calcular la distancia euclidiana entre cada par de puntos en dos conjuntos:
```python
X1 = np.random.randn(100, 3)  # 100 puntos de 3 dimensiones
X2 = np.random.randn(50, 3)   # 50 puntos de 3 dimensiones
# Resultado debe ser matriz 100x50 de distancias
```

---

## Pandas

### EJ-05-003: EDA básico
Descargar el dataset Titanic (CSV). En ≤15 líneas de Pandas:
- ¿Cuántos pasajeros sobrevivieron por clase?
- ¿Cuál es la edad media por sexo?
- ¿Hay correlación entre edad y precio del ticket?
- Mostrar la distribución de edades con describe()

### EJ-05-004: Limpieza de datos
En el mismo dataset Titanic:
- Imputar valores faltantes en `Age` con la mediana por clase
- Eliminar columnas con >50% de NaN
- Codificar `Sex` como 0/1

### EJ-05-005: GroupBy avanzado
Usar el dataset de ventas (o cualquier CSV disponible):
- Agrupar por categoría y mes
- Calcular: suma, media, y percentil 90 de ventas
- Identificar el producto con mayor crecimiento mes a mes

---

## Matplotlib / Seaborn

### EJ-05-006: Visualizar modelos
Plotear en la misma figura:
- Datos de entrenamiento (scatter)
- Línea de regresión de tu modelo Go (importar los pesos)
- Línea de regresión de Scikit-Learn
- Comparar visualmente si coinciden

### EJ-05-007: Matriz de correlación
Para el dataset Iris:
- Calcular correlación entre todas las features
- Visualizar con `sns.heatmap`
- Identificar features más correlacionadas

---

## Scikit-Learn — Comparación

### EJ-05-008: Reproducir golearn en Python
Para cada modelo implementado en Etapa 4:
```python
from sklearn.linear_model import LinearRegression
from sklearn.metrics import mean_squared_error

# Python
skl_model = LinearRegression()
skl_model.fit(X_train, y_train)
skl_preds = skl_model.predict(X_test)
skl_mse = mean_squared_error(y_test, skl_preds)

# Tu Go (importar resultados desde archivo)
go_mse = ...  # cargar desde CSV generado por tu Go

print(f"Sklearn MSE: {skl_mse:.4f}")
print(f"Go MSE:      {go_mse:.4f}")
print(f"Diferencia:  {abs(skl_mse - go_mse):.6f}")
```

### EJ-05-009: Pipeline completo
```python
from sklearn.pipeline import Pipeline
from sklearn.preprocessing import StandardScaler
from sklearn.model_selection import GridSearchCV
from sklearn.ensemble import RandomForestClassifier

pipe = Pipeline([
    ('scaler', StandardScaler()),
    ('clf', RandomForestClassifier())
])

param_grid = {'clf__n_estimators': [10, 50, 100], 'clf__max_depth': [3, 5, None]}
grid = GridSearchCV(pipe, param_grid, cv=5, scoring='accuracy')
grid.fit(X_train, y_train)

print(f"Best params: {grid.best_params_}")
print(f"Best score:  {grid.best_score_:.4f}")
```

---

## Proyecto integrador

Hacer EDA completo y modelado del dataset Housing Prices (Kaggle):
1. EDA: distribuciones, correlaciones, outliers
2. Feature engineering: encoding, scaling
3. Entrenar: LinearRegression, RandomForest, GradientBoosting
4. Evaluar con cross-validation
5. Visualizar feature importance
6. Comparar el mejor modelo con tu implementación Go
