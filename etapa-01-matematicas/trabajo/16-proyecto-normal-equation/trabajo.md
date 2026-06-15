# 16 — PROYECTO FINAL: Derivar la Ecuación Normal

**EJ:** EJ-01-014 | **Prerequisito:** TODOS los anteriores (01–15) | **Estado:** ☐ Pendiente

---

## Objetivo
Demostrar que dominas la Etapa 1 completa. Este proyecto integra álgebra lineal, cálculo matricial y probabilidad en una sola derivación: la solución exacta de regresión lineal.

**Si puedes completar esto de memoria y sin errores, estás listo para la Etapa 2.**

---

## La demostración

### Enunciado
Dado el modelo de regresión lineal: **y = Xw + ε**

Donde:
- X ∈ ℝⁿˣᵈ: matriz de datos (n muestras, d features)
- w ∈ ℝᵈ: vector de pesos (parámetros a encontrar)
- y ∈ ℝⁿ: vector de targets
- ε ~ N(0, σ²I): ruido Gaussiano

Demostrar que el estimador de mínimos cuadrados es:
**w* = (XᵀX)⁻¹Xᵀy**

---

### Paso 1: Definir la función de pérdida
Escribir MSE en forma matricial:

```
L(w) = (1/n) ‖y − Xw‖²
     = 

```

Expandir el cuadrado de la norma:

```
L(w) = (1/n)(y − Xw)ᵀ(y − Xw)
     = (1/n)[                                    ]
     = (1/n)[yᵀy − 2wᵀXᵀy + wᵀXᵀXw]

```

---

### Paso 2: Calcular el gradiente
Aplicar cálculo matricial a cada término:

```
∂/∂w [yᵀy]      = (¿qué es esto?)

∂/∂w [−2wᵀXᵀy]  = (regla: ∂(aᵀw)/∂w = a)

∂/∂w [wᵀXᵀXw]   = (regla: ∂(wᵀAw)/∂w = 2Aw, A=XᵀX es simétrica)

∂L/∂w = (1/n)[                           ]
       = (2/n)[XᵀXw − Xᵀy]

```

---

### Paso 3: Encontrar el mínimo
Igualar el gradiente a cero:

```
(2/n)[XᵀXw − Xᵀy] = 0

XᵀXw = Xᵀy    ← ecuaciones normales

w* =           ← despejar w (¿cuándo existe la inversa?)

```

¿Cuándo NO existe (XᵀX)⁻¹? ¿Qué implica para los datos?

```
Tu respuesta:

```

---

### Paso 4: Verificar que es mínimo (no máximo)
Calcular el Hessiano de L(w):

```
H = ∂²L/∂w² = (2/n) XᵀX

```

¿Es H positivo semidefinido? Demostrar:

```
Para cualquier vector v ≠ 0:
vᵀHv = (2/n) vᵀXᵀXv = (2/n) ‖Xv‖² ≥ 0

¿Cuándo es estrictamente positivo definido?:

```

Conclusión: w* es **mínimo global**.

---

### Paso 5: Perspectiva probabilística
Demostrar que w_MLE = w_OLS (misma solución desde otra perspectiva):

```
Bajo y | X ~ N(Xw, σ²I), el log-likelihood es:

ℓ(w) = −(1/2σ²) ‖y − Xw‖² + constante

Maximizar ℓ(w) ⟺ minimizar ‖y − Xw‖²

∴ w_MLE = w_OLS = (XᵀX)⁻¹Xᵀy   ∎

```

---

### Paso 6: Implementación numérica
Verificar la derivación con un ejemplo concreto.

Dado el dataset (predecir precio de casa por m²):

```
X_raw = [50, 80, 100, 120, 150]   (m²)
y     = [100, 150, 200, 220, 300]  (precio en miles)
```

Agregar columna de unos para el bias (intercept):

```
X = [[1, 50],
     [1, 80],
     [1, 100],
     [1, 120],
     [1, 150]]
```

Calcular manualmente (o con calculadora) w* = (XᵀX)⁻¹Xᵀy:

```
XᵀX = 

Xᵀy = 

(XᵀX)⁻¹ = 

w* = [w₀, w₁] = 

```

Interpretar los coeficientes:

```
w₀ (intercept) = 
w₁ (slope) = 

Precio predicho para una casa de 90m²:
ŷ = w₀ + w₁ · 90 = 

```

---

### Evaluación final

Responder sin mirar notas:

```
1. ¿Cuál es la complejidad computacional de calcular (XᵀX)⁻¹?

2. ¿Por qué en deep learning usamos gradient descent en lugar de la ecuación normal?

3. ¿Qué es la pseudoinversa y cuándo se usa en lugar de (XᵀX)⁻¹Xᵀ?

4. ¿Qué pasa si dos features de X están perfectamente correlacionadas?

5. ¿Cómo cambia w* si agrego regularización L2?
```

---

## Criterio de éxito — Listo para Etapa 2 si:
- Derivaste los 5 pasos sin copiar
- Puedes explicar cada paso a alguien
- Respondiste las 5 preguntas finales
- Los coeficientes de Paso 6 predicen razonablemente bien (R² > 0.95)
- Tiempo total < 2 horas (al principio puede tomar más, con práctica mejora)
