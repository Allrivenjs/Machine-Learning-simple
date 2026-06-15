# 03 — Determinante e Inversa

**EJ:** EJ-01-003 | **Prerequisito:** [02-matrices-operaciones](../02-matrices-operaciones/trabajo.md) | **Estado:** ☐ Pendiente

---

## Objetivo
Entender qué significa el determinante geométricamente y cuándo existe la inversa. Necesarios para resolver sistemas lineales (ecuación normal de regresión lineal).

## Revisar antes de empezar
- 3Blue1Brown — The determinant (capítulo 6)
- 3Blue1Brown — Inverse matrices, column space and null space (capítulo 7)
- Math4ML — Sección 2.3.3, 2.3.4

---

## Tarea

### Parte 1: Determinante 2×2
Para la fórmula: **det([[a,b],[c,d]]) = ad − bc**

Calcular det(A) donde **A = [[2, 1], [5, 3]]**

```
Tu solución:

```

Calcular det(B) donde **B = [[1, 2], [2, 4]]**

```
Tu solución:

```

**¿Qué significa que det(B) = 0?** ¿Qué pasó con las filas de B?

```
Tu respuesta:

```

---

### Parte 2: Interpretación geométrica
Dado **A = [[3, 0], [0, 2]]** (matriz de escala):

- ¿Cuánto escala el área un cuadrado unitario?
- Calcular det(A) y relacionarlo con esa escala

```
Tu solución y explicación:

```

Dado **A = [[0, -1], [1, 0]]** (rotación 90°):

- ¿Cambia el área al rotar?
- ¿Cuánto debería ser det(A)?
- Calcular y verificar

```
Tu solución:

```

---

### Parte 3: Inversa 2×2
Fórmula: **A⁻¹ = (1/det(A)) · [[d, -b], [-c, a]]**

Calcular **A⁻¹** donde **A = [[2, 1], [5, 3]]**

```
Tu solución paso a paso:

det(A) = 

A⁻¹ = 

```

Verificar: **A · A⁻¹ = I** (identidad)

```
Verificación:

```

---

### Parte 4: ¿Cuándo NO existe la inversa?
Intentar calcular la inversa de **B = [[1, 2], [2, 4]]**

```
¿Qué pasa? ¿Por qué?

```

Listar 2 condiciones bajo las cuales una matriz NO tiene inversa:

```
1. 
2. 
```

---

### Parte 5: Resolver sistema lineal
Resolver **Ax = b** usando A⁻¹:

```
A = [[2, 1],    b = [5]
     [5, 3]]        [13]
```

Fórmula: **x = A⁻¹b**

```
Tu solución:

```

Verificar sustituyendo x en Ax:

```
Verificación:

```

---

## Criterio de éxito
- det([[2,1],[5,3]]) = 1
- A⁻¹ · A = [[1,0],[0,1]] dentro de tolerancia
- x en Parte 5 = [2, 1]
