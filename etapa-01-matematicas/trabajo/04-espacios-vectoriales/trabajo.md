# 04 — Espacios Vectoriales

**EJ:** conceptual | **Prerequisito:** [03-determinante-inversa](../03-determinante-inversa/trabajo.md) | **Estado:** ☐ Pendiente

---

## Objetivo
Entender qué es una base, dimensión, span e independencia lineal. Son los conceptos que explican por qué algunas matrices son invertibles y otras no, y por qué PCA funciona.

## Revisar antes de empezar
- 3Blue1Brown — Span, basis vectors (capítulo 2)
- 3Blue1Brown — Linear independence (capítulo 2)
- Math4ML — Sección 2.4, 2.6

---

## Tarea

### Parte 1: Span
Explicar con tus palabras qué es el **span** de un conjunto de vectores.

```
Tu definición:

```

¿Cuál es el span de **{[1,0], [0,1]}** en ℝ²?

```
Tu respuesta:

```

¿Cuál es el span de **{[1,0], [2,0]}** en ℝ²? ¿Por qué es diferente?

```
Tu respuesta:

```

---

### Parte 2: Independencia lineal
Determinar si los siguientes conjuntos son linealmente independientes:

**a) {[1,0,0], [0,1,0], [0,0,1]}**

```
¿Independientes? Justificación:

```

**b) {[1,2], [2,4]}**

```
¿Independientes? Justificación:
(Pista: ¿puede uno escribirse como combinación lineal del otro?)

```

**c) {[1,1,0], [0,1,1], [1,0,1]}**

```
¿Independientes? 
Intentar resolver: α[1,1,0] + β[0,1,1] + γ[1,0,1] = [0,0,0]
¿La única solución es α=β=γ=0?

```

---

### Parte 3: Base y dimensión
¿Qué es una **base** de un espacio vectorial?

```
Tu definición (2 propiedades):
1. 
2. 
```

¿Cuál es la dimensión de ℝ³? ¿Por qué?

```
Tu respuesta:

```

Dado el espacio generado por **{[1,1,0], [0,1,1], [1,2,1]}**, ¿cuál es su dimensión?
(Pista: verificar si son linealmente independientes primero)

```
Tu solución:

```

---

### Parte 4: Conexión con el determinante
Usar el det para verificar si las filas de una matriz forman una base:

```
A = [[1, 1, 0],
     [0, 1, 1],
     [1, 2, 1]]
```

Calcular det(A). Si det ≠ 0, las filas son linealmente independientes.

```
Tu cálculo (expansión por cofactores o regla de Sarrus):

det(A) = 

Conclusión:

```

---

### Parte 5: Aplicación — Column space y null space
Para la matriz **A = [[1, 2], [2, 4]]**:

**Column space** (span de las columnas de A): ¿qué vectores b tienen solución en Ax = b?

```
Tu respuesta:

```

**Null space** (vectores x tal que Ax = 0): encontrar un vector no-cero en el null space.

```
Tu solución:

```

---

## Criterio de éxito
- Puedes distinguir entre span e independencia lineal
- Entiendes por qué det=0 implica columnas dependientes
- Puedes encontrar un vector en el null space de A = [[1,2],[2,4]]
