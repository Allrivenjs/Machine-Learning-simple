# 13 — Esperanza, Varianza y Correlación

**EJ:** EJ-01-013 | **Prerequisito:** [12-distribuciones](../12-distribuciones/trabajo.md) | **Estado:** ☐ Pendiente

---

## Objetivo
Estas medidas resumen una distribución con pocos números. En ML: la media informa el bias, la varianza informa el overfitting, la correlación informa qué features son útiles.

## Revisar antes de empezar
- Math4ML — Sección 6.4
- ISLR — Sección 2.1.1

---

## Tarea

### Parte 1: Esperanza (media)
Para una variable discreta: **E[X] = Σ xᵢ · p(xᵢ)**

Dado X con distribución:

```
x   | 1 | 2 | 3 | 4 | 5
----|---|---|---|---|---
p   |0.1|0.2|0.4|0.2|0.1
```

Calcular E[X]:

```
E[X] = 

```

Propiedades — verificar usando E[X] calculado:

**E[2X + 3]** = 2·E[X] + 3 =

```
Tu cálculo:

```

**E[X²]** (no es (E[X])²):

```
E[X²] = Σ x² · p(x) = 

```

---

### Parte 2: Varianza
**Var(X) = E[(X − E[X])²] = E[X²] − (E[X])²**

Usando los valores de Parte 1:

```
Var(X) = E[X²] - (E[X])² = 

```

¿Qué mide la varianza? ¿Cuándo es 0?

```
Tu respuesta:

```

Para el dataset **{2, 4, 4, 4, 5, 5, 7, 9}**:

```
Media (μ̂) = 

Varianza muestral (usar n-1 en denominador, Bessel's correction):
σ̂² = (1/(n-1)) Σ(xᵢ - μ̂)² = 

Desviación estándar σ̂ = 

```

---

### Parte 3: Covarianza
**Cov(X,Y) = E[(X − E[X])(Y − E[Y])]**

Dado:
```
X = [1, 2, 3, 4, 5]
Y = [2, 4, 5, 4, 5]
```

Calcular Cov(X,Y):

```
E[X] = 
E[Y] = 

Cov(X,Y) = (1/n) Σ(xᵢ - E[X])(yᵢ - E[Y])

= (1/5) · [(1-3)(2-4) + (2-3)(4-4) + (3-3)(5-4) + (4-3)(4-4) + (5-3)(5-4)]
= (1/5) · [            +            +            +            +           ]
= 

```

**¿Qué signo tiene Cov(X,Y)?** ¿Qué significa?

```
Tu respuesta:

```

---

### Parte 4: Correlación de Pearson
**r(X,Y) = Cov(X,Y) / (σ_X · σ_Y)**

Normaliza la covarianza al rango [−1, 1].

Calcular r para X e Y de Parte 3:

```
σ_X = 
σ_Y = 

r = Cov(X,Y) / (σ_X · σ_Y) = 

```

Interpretar:

```
r cercano a 1:  significa
r cercano a -1: significa
r cercano a 0:  significa

Tu r:
```

---

### Parte 5: Bias-Variance Tradeoff
El error esperado de un modelo se descompone:

**E[(y − ŷ)²] = Bias² + Varianza + Ruido**

- **Bias²**: error por simplificación del modelo (underfitting)
- **Varianza**: sensibilidad a fluctuaciones en los datos (overfitting)
- **Ruido**: error irreducible de los datos

Para cada caso, identificar qué domina (Bias o Varianza):

```
a) Modelo de regresión lineal en datos con relación cúbica:
   Domina: 

b) Árbol de decisión con profundidad 100 en dataset de 50 muestras:
   Domina: 

c) Red neuronal pequeña entrenada por mucho tiempo sin regularización:
   Domina: 

d) KNN con K=1:
   Domina: 

```

---

## Criterio de éxito
- E[X] para la distribución discreta = 3.0
- Cov(X,Y) > 0 (correlación positiva)
- r(X,Y) ≈ 0.87
- Puedes identificar underfitting vs overfitting por su bias/varianza
