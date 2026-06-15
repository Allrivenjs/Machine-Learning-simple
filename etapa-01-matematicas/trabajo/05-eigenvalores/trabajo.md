# 05 — Eigenvalores y Eigenvectores

**EJ:** EJ-01-004 | **Prerequisito:** [03-determinante-inversa](../03-determinante-inversa/trabajo.md) + [04-espacios-vectoriales](../04-espacios-vectoriales/trabajo.md) | **Estado:** ☐ Pendiente

---

## Objetivo
Los eigenvectores son las direcciones que una transformación lineal no rota — solo escala. Son la base de PCA, PageRank y la comprensión de redes neuronales profundas.

## Revisar antes de empezar
- 3Blue1Brown — Eigenvectors and eigenvalues (capítulo 14) — **muy importante**
- Math4ML — Sección 2.7

---

## Tarea

### Parte 1: Definición
Un eigenvector **v** de A satisface: **Av = λv**

Explicar con tus palabras qué significa esto geométricamente:

```
Tu explicación:

```

---

### Parte 2: Encontrar eigenvalores
Para encontrar eigenvalores, resolver: **det(A − λI) = 0**

Dado **A = [[4, 1], [2, 3]]**, encontrar los eigenvalores:

```
Paso 1: A - λI = [[4-λ, 1], [2, 3-λ]]

Paso 2: det(A - λI) = (4-λ)(3-λ) - (1)(2)
                    = 
                    = (expandir)
                    = (simplificar — debe quedar polinomio en λ)

Paso 3: igualar a 0 y resolver:

λ₁ = 
λ₂ = 

```

---

### Parte 3: Encontrar eigenvectores
Para cada eigenvalor λᵢ, resolver **(A − λᵢI)v = 0**

**Para λ₁:**

```
(A - λ₁I)v = 0

Sistema de ecuaciones:

Solución (eigenvector v₁):

```

**Para λ₂:**

```
(A - λ₂I)v = 0

Sistema de ecuaciones:

Solución (eigenvector v₂):

```

---

### Parte 4: Verificación
Verificar que **Av₁ = λ₁v₁** (sustituir y multiplicar):

```
Av₁ = 

λ₁v₁ = 

¿Son iguales?:

```

---

### Parte 5: Matriz diagonal
Si A tiene eigenvectores linealmente independientes, se puede escribir: **A = PDP⁻¹**
- P = matriz con eigenvectores como columnas
- D = matriz diagonal con eigenvalores

Para A = [[4,1],[2,3]]:

```
P = [v₁ | v₂] = 

D = [[λ₁, 0],
     [0, λ₂]] = 

```

¿Para qué sirve la diagonalización? (pista: potencias de matrices, Aⁿ = PDⁿP⁻¹)

```
Tu respuesta:

```

---

### Parte 6: Conexión con ML
Marcar con ✓ los usos de eigenvalores en ML que puedes intuir:

- [ ] PCA: las componentes principales son eigenvectores de la matriz de covarianza
- [ ] Estabilidad de gradient descent: eigenvalores del Hessiano determinan el learning rate máximo
- [ ] Google PageRank: ranking de páginas = eigenvector dominante de la matriz de links
- [ ] Compresión de imágenes con SVD

Para cada uno marcado, escribir 1 oración explicando la conexión:

```
Tu respuesta:

```

---

## Criterio de éxito
- Eigenvalores de [[4,1],[2,3]] son λ₁=5 y λ₂=2
- Cada eigenvector verifica Av = λv al sustituir
- Entiendes por qué eigenvectors "no rotan"
