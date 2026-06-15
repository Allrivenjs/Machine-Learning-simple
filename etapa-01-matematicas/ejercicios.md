# Ejercicios: Etapa 1 — Matemáticas Fundamentales

Todos con lápiz y papel. Sin código todavía.

---

## Álgebra Lineal

### EJ-01-001: Operaciones vectoriales
Dado v₁ = [2, 3, 1] y v₂ = [1, 0, 4]:
- Calcular v₁ + v₂
- Calcular v₁ · v₂ (producto punto)
- Calcular ‖v₁‖₂ (norma euclidiana)
- ¿Son ortogonales? Justificar.

### EJ-01-002: Multiplicación matricial
```
A = [[1, 2], [3, 4]]
B = [[5, 6], [7, 8]]
```
- Calcular AB y BA
- ¿AB = BA? ¿Por qué importa esto en ML?

### EJ-01-003: Determinante e inversa
Para A = [[2, 1], [5, 3]]:
- Calcular det(A)
- Calcular A⁻¹
- Verificar: A · A⁻¹ = I

### EJ-01-004: Eigenvalores
Para A = [[4, 1], [2, 3]]:
- Encontrar los eigenvalores resolviendo det(A - λI) = 0
- Para cada eigenvalor, encontrar el eigenvector correspondiente

### EJ-01-005: SVD conceptual
Explicar con tus palabras:
- ¿Qué descomposición hace SVD? (A = UΣVᵀ)
- ¿Qué representan U, Σ, V?
- ¿Cómo se usa SVD para comprimir una imagen?

---

## Cálculo

### EJ-01-006: Gradiente de MSE
La función de pérdida MSE es: L(w) = (1/n) Σ(yᵢ - wᵀxᵢ)²

Derivar ∂L/∂w paso a paso.

### EJ-01-007: Regla de la cadena
Si f(x) = (3x² + 2)⁵, calcular f'(x) usando regla de la cadena.

### EJ-01-008: Derivadas parciales
Para f(x, y) = x²y + 3xy²:
- Calcular ∂f/∂x
- Calcular ∂f/∂y
- Calcular el gradiente ∇f

### EJ-01-009: Punto de mínimo
Para f(x) = x² - 4x + 5:
- Encontrar el mínimo analíticamente (igualando f'(x) = 0)
- Verificar que es mínimo y no máximo usando f''(x)

---

## Probabilidad y Estadística

### EJ-01-010: Bayes desde cero
Un test médico tiene:
- P(positivo | enfermo) = 0.95
- P(positivo | sano) = 0.05
- P(enfermo) = 0.01

Si una persona da positivo, ¿cuál es P(enfermo | positivo)?

### EJ-01-011: MLE para distribución Normal
Dado un dataset {2, 4, 4, 4, 5, 5, 7, 9}:
- Calcular la media muestral μ̂
- Calcular la varianza muestral σ̂²
- Demostrar que estos son los estimadores MLE para una Normal

### EJ-01-012: Distribuciones
Describir situaciones reales donde usarías:
- Distribución Normal
- Distribución Bernoulli
- Distribución Poisson

### EJ-01-013: Correlación
Dado X = [1, 2, 3, 4, 5] e Y = [2, 4, 5, 4, 5]:
- Calcular la correlación de Pearson
- Interpretar el resultado

---

## Proyecto integrador
Derivar la solución closed-form de regresión lineal:

Dado el problema de minimizar ‖y - Xw‖², demostrar que la solución óptima es:
```
w* = (XᵀX)⁻¹Xᵀy
```
Esto requiere: producto matricial, transpuesta, inversa, gradiente, e igualar a cero.
