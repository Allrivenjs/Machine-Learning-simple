# 10 — Optimización: Puntos Mínimos

**EJ:** EJ-01-009 | **Prerequisito:** [09-gradiente-regla-cadena](../09-gradiente-regla-cadena/trabajo.md) | **Estado:** ☐ Pendiente

---

## Objetivo
Encontrar mínimos analíticamente y entender cuándo un punto crítico es mínimo, máximo o punto silla. En ML, el mínimo de la función de pérdida es el modelo óptimo.

## Revisar antes de empezar
- Math4ML — Sección 5.1.1, 7.1
- Khan Academy — Multivariable calculus: finding local extrema

---

## Tarea

### Parte 1: Condición necesaria en 1D
Un mínimo o máximo local requiere **f'(x) = 0**.
El tipo se determina con f''(x): positivo = mínimo, negativo = máximo.

Para **f(x) = x³ − 3x + 2**:

```
f'(x) = 

Igualar a 0:
Puntos críticos: x₁ = , x₂ = 

f''(x) = 

f''(x₁) = → ¿mínimo o máximo?
f''(x₂) = → ¿mínimo o máximo?

f(x₁) = 
f(x₂) = 

```

---

### Parte 2: Condición en 2D — Hessiano
Para f(x,y), el Hessiano es:

```
H = [[∂²f/∂x², ∂²f/∂x∂y],
     [∂²f/∂x∂y, ∂²f/∂y²]]
```

Un punto crítico es mínimo si H es **positivo definido** (todos eigenvalores > 0).

Para **f(x,y) = x² + 4y² − 2xy**:

Calcular el gradiente y encontrar el punto crítico:

```
∂f/∂x = 
∂f/∂y = 

Igualar a [0,0]:
Sistema de ecuaciones:

Punto crítico: (x*, y*) = 

```

Calcular el Hessiano en ese punto:

```
∂²f/∂x² = 
∂²f/∂y² = 
∂²f/∂x∂y = 

H = [[  ,  ],
     [  ,  ]]

```

¿Es H positivo definido? (verificar: det(H) > 0 y H₁₁ > 0)

```
det(H) = 
H₁₁ = 
¿Mínimo?: 

```

---

### Parte 3: Punto silla
Para **f(x,y) = x² − y²**:

```
Punto crítico:

Hessiano:

Eigenvalores del Hessiano:

¿Qué tipo de punto es? ¿Por qué?

```

Los puntos silla son comunes en redes neuronales profundas. ¿Por qué son menos problemáticos de lo que parecen?

```
Tu respuesta (pista: gradient descent puede escapar de ellos con momentum):

```

---

### Parte 4: Convexidad
Una función **convexa** tiene un único mínimo global.

Para funciones convexas, gradient descent siempre converge al óptimo.

¿Cuáles de estas funciones son convexas? Justificar.

```
a) f(x) = x²            → ¿convexa? 
b) f(x) = |x|           → ¿convexa? 
c) f(x) = x³            → ¿convexa? 
d) f(x) = MSE de regresión lineal → ¿convexa? (pista: es cuadrática en los pesos)
e) f(x) = loss de red neuronal    → ¿convexa? 

```

---

### Parte 5: Conexión con ML
La función de pérdida MSE es: **L(w) = (1/n) Σ(yᵢ − wᵀxᵢ)²**

Demostrar que L(w) es **convexa** respecto a w:

```
Pista: L es cuadrática en w. Una función cuadrática f(w) = wᵀAw + bᵀw + c
es convexa si y solo si A es semidefinida positiva.

Identificar A en L(w):

¿Es A semidefinida positiva? ¿Por qué?

Conclusión: ¿tiene L un único mínimo global?

```

---

## Criterio de éxito
- Punto crítico de x³−3x+2: mínimo en x=1, máximo en x=−1
- MSE de regresión lineal es convexa → garantiza convergencia de GD
- Entiendes la diferencia entre punto silla, mínimo local y mínimo global
