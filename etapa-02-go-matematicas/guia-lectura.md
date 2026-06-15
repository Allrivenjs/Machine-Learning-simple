# Etapa 2: Programación Matemática en Go

Duración: 1 mes | Prerequisito: Etapa 1

---

## Objetivo

Implementar todo lo aprendido en Etapa 1 como código Go. Sin librerías externas — solo stdlib. El objetivo es que las operaciones matemáticas dejen de ser abstractas.

---

## Recursos

1. **A Tour of Go** — https://go.dev/tour (si no conoces Go bien)
2. **Numerical Recipes** (referencia para algoritmos numéricos)
3. El código de referencia es la matemática de Etapa 1

---

## Temas a implementar (en orden)

### Bloque 1: Tipos base
- [ ] `type Vector []float64`
- [ ] `type Matrix struct { Rows, Cols int; Data [][]float64 }`
- [ ] Constructor `NewMatrix(rows, cols int) Matrix`
- [ ] Constructor `NewVector(data ...float64) Vector`

### Bloque 2: Operaciones vectoriales
- [ ] Suma de vectores
- [ ] Resta de vectores
- [ ] Multiplicación por escalar
- [ ] Producto punto
- [ ] Norma L1
- [ ] Norma L2 (euclidiana)
- [ ] Normalización

### Bloque 3: Operaciones matriciales
- [ ] Suma de matrices
- [ ] Multiplicación matricial
- [ ] Transpuesta
- [ ] Traza
- [ ] Determinante (para 2×2 y 3×3 inicialmente, luego general)
- [ ] Inversa (Gauss-Jordan)

### Bloque 4: Descomposiciones
- [ ] Eliminación Gaussiana
- [ ] LU Decomposition
- [ ] QR Decomposition (Gram-Schmidt)
- [ ] Eigenvalores (método de potencia para el dominante)
- [ ] SVD (implementación básica)

### Bloque 5: Estadística
- [ ] Media
- [ ] Varianza
- [ ] Desviación estándar
- [ ] Covarianza
- [ ] Correlación de Pearson

---

## Estructura de archivos recomendada

```
etapa-02-go-matematicas/
├── vector.go
├── vector_test.go
├── matrix.go
├── matrix_test.go
├── decompositions.go
├── decompositions_test.go
├── stats.go
└── stats_test.go
```

---

## Señales de que dominas la etapa

- Puedes multiplicar matrices de cualquier dimensión
- Tu implementación de inversa produce I cuando multiplicas A × A⁻¹
- Puedes resolver un sistema lineal Ax = b usando tu LU
- Tus tests pasan comparando resultados contra cálculos manuales de Etapa 1
