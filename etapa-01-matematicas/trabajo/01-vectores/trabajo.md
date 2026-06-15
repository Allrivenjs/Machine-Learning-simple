# 01 — Vectores

**EJ:** EJ-01-001 | **Prerequisito:** ninguno | **Estado:** ☐ Pendiente

---

## Objetivo
Dominar operaciones fundamentales sobre vectores. Todo aquí es base para matrices, gradientes y distancias en ML.

## Revisar antes de empezar
- 3Blue1Brown — Vectors, what even are they? (Essence of Linear Algebra, capítulo 1)
- Mathematics for Machine Learning — Sección 2.4

---

## Tarea

### Parte 1: Operaciones básicas
Dado **v₁ = [2, 3, 1]** y **v₂ = [1, 0, 4]**, calcular a mano:

**a) v₁ + v₂**

```
Tu solución:
v₁ + v₂ = [2 + 1, 3 + 0, 1 + 4] = [3, 3, 5]
```

**b) v₁ − v₂**

```
Tu solución:
v₁ − v₂ = [2 - 1, 3 - 0, 1 - 4] = [1, 3, -3]
```

**c) 3 · v₁ (multiplicación por escalar)**

```
Tu solución:
3 · v₁ = [3*2, 3*3, 3*1] = [6, 9, 3]
```

---

### Parte 2: Producto punto
Calcular **v₁ · v₂**

Recordar: v · w = Σ vᵢwᵢ

```
Tu solución:
v₁ · v₂ = (2*1) + (3*0) + (1*4) = 2 + 0 + 4 = 6
```

**¿Qué significa geométricamente cuando el producto punto es 0?**

```
Tu respuesta:
Cuando el producto punto es 0, significa que los vectores son ortogonales, es decir, 
forman un ángulo de 90 grados entre sí.
Por lo cual es util para determinar si dos vectores son perpendiculares en el espacio, 
lo que tiene aplicaciones en ML como en la reducción de dimensionalidad o 
en la evaluación de similitud entre vectores.
Esto lo podemos persivir en la vida real cuando por ejemplo dos personas caminan 
en direcciones completamente diferentes, sus trayectorias serían ortogonales.
```

---

### Parte 3: Normas
Calcular **‖v₁‖₂** (norma euclidiana)

Recordar: ‖v‖₂ = √(Σ vᵢ²)

```
Tu solución:
‖v₁‖₂ = √(2² + 3² + 1²) = √(4 + 9 + 1) = √14 ≈ 3.74
```

Calcular **‖v₁‖₁** (norma Manhattan)

Recordar: ‖v‖₁ = Σ |vᵢ|

```
Tu solución:
‖v₁‖₁ = |2| + |3| + |1| = 2 + 3 + 1 = 6
```

---

### Parte 4: Ortogonalidad
¿Son v₁ y v₂ ortogonales? Justificar usando el resultado de Parte 2.

```
Tu respuesta:
No, v₁ y v₂ no son ortogonales porque su producto punto es 6, que no es igual a 0.
```

Construir un vector **v₃** que sea ortogonal a **v₁ = [1, 0]** en 2D.

```
Tu solución:
Un vector ortogonal a v₁ = [1, 0] en 2D es v₃ = [0, 1], ya que su producto punto es:
v₁ · v₃ = (1*0) + (0*1) = 0, lo que indica que son ortogonales.
```

---

### Parte 5: Vector unitario
Normalizar **v₁ = [3, 4]** para obtener un vector unitario (‖v‖₂ = 1).

```
Tu solución:
Primero calculamos la norma euclidiana de v₁:
‖v₁‖₂ = √(3² + 4²) = √(9 + 16) = √25 = 5
Luego, para normalizar v₁, dividimos cada componente por la norma:
v₁_unitario = [3/5, 4/5] = [0.6, 0.8]
Quedando el vector unitario de v₁ como [0.6, 0.8].
```

**¿Por qué normalizar vectores importa en ML?** (pista: cosine similarity, KNN)

```
Tu respuesta:
Normalizar vectores es importante en ML porque permite comparar vectores de diferentes magnitudes de manera justa.
Por ejemplo, en la similitud coseno (cosine similarity), se mide el ángulo
entre dos vectores, y si no están normalizados, la magnitud de los vectores puede influir en el resultado,
haciendo que dos vectores con la misma dirección pero diferente longitud no sean considerados similares.
En KNN (K-Nearest Neighbors), la distancia entre puntos es crucial, y si
los vectores no están normalizados, los puntos con magnitudes mayores podrían dominar la distancia,
lo que podría llevar a resultados sesgados en la clasificación o regresión.
Ejemplo en la vida real:
Si estás comparando la similitud entre dos documentos representados como vectores de palabras,
si uno de los documentos es mucho más largo que el otro, su vector tendrá una magnitud
mayor, lo que podría hacer que parezcan menos similares de lo que realmente son si no se normalizan.


```

---

## Criterio de éxito
- ‖[3,4]‖₂ = 5.0 exactamente
- El vector unitario de [3,4] es [0.6, 0.8]
- Puedes explicar la diferencia entre L1 y L2 en tus palabras
