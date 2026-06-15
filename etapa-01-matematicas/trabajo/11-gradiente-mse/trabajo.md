# 11 — Gradiente del MSE

**EJ:** EJ-01-006 | **Prerequisito:** [09-gradiente-regla-cadena](../09-gradiente-regla-cadena/trabajo.md) + [02-matrices-operaciones](../02-matrices-operaciones/trabajo.md) | **Estado:** ☐ Pendiente

---

## Objetivo
Derivar el gradiente de la función de pérdida MSE respecto a los pesos. Este es el cálculo que hace gradient descent en regresión lineal — si lo entiendes aquí, entiendes el corazón del aprendizaje.

## Revisar antes de empezar
- Math4ML — Sección 5.2.1 (matrix calculus)
- [09-gradiente-regla-cadena](../09-gradiente-regla-cadena/trabajo.md) completado

---

## Tarea

### Parte 1: Cálculo matricial — reglas clave
Memorizar estas dos reglas (las usaremos todo el tiempo):

```
Si f(w) = wᵀAw  →  ∂f/∂w = (A + Aᵀ)w = 2Aw (si A es simétrica)
Si f(w) = aᵀw   →  ∂f/∂w = a
```

Verificar la segunda regla para **f(w) = aᵀw = a₁w₁ + a₂w₂**:

```
∂f/∂w₁ = 
∂f/∂w₂ = 
∂f/∂w = (como vector) = 

¿Coincide con la regla?:

```

---

### Parte 2: MSE escalar (1 muestra)
Una muestra: x, y, predicción ŷ = wᵀx + b

Pérdida: **L = (y − ŷ)²**

Calcular **∂L/∂w** y **∂L/∂b** usando la regla de la cadena:

```
Sea e = y - ŷ = y - wᵀx - b  (el error)

∂L/∂ŷ = 

∂ŷ/∂w = 

∂ŷ/∂b = 

∂L/∂w = ∂L/∂ŷ · ∂ŷ/∂w = 

∂L/∂b = ∂L/∂ŷ · ∂ŷ/∂b = 

```

---

### Parte 3: MSE sobre N muestras (forma vectorial)
**L(w) = (1/n) ‖y − Xw‖²**

Donde X es la matriz de datos (n×d), y es el vector de targets (n×1), w son los pesos (d×1).

Expandir: **L = (1/n)(y − Xw)ᵀ(y − Xw)**

```
Expandir el producto:
L = (1/n)[yᵀy − yᵀXw − wᵀXᵀy + wᵀXᵀXw]

Simplificar (nota: yᵀXw es escalar, igual a wᵀXᵀy):
L = (1/n)[yᵀy − 2wᵀXᵀy + wᵀXᵀXw]

```

Aplicar las reglas de cálculo matricial para calcular **∂L/∂w**:

```
∂/∂w [yᵀy]    = (constante respecto a w)

∂/∂w [−2wᵀXᵀy] = (usar regla ∂(aᵀw)/∂w = a)

∂/∂w [wᵀXᵀXw]  = (usar regla ∂(wᵀAw)/∂w = 2Aw con A=XᵀX)

∂L/∂w = (1/n) · [                                     ]
       = (2/n) · [                                     ]

```

---

### Parte 4: Ecuación normal (mínimo analítico)
Igualando ∂L/∂w = 0:

```
(2/n)[XᵀXw − Xᵀy] = 0
→ XᵀXw = Xᵀy
→ w* = (XᵀX)⁻¹Xᵀy

```

Verificar con datos pequeños. Dado:

```
X = [[1, 1],    y = [2]
     [1, 2],        [3]
     [1, 3]]        [4]
```

Calcular w* = (XᵀX)⁻¹Xᵀy a mano:

```
XᵀX = 

Xᵀy = 

(XᵀX)⁻¹ = 

w* = 

```

Verificar: ¿las predicciones Xw* se acercan a y?

```
Xw* = 

Error = y - Xw* = 

```

---

### Parte 5: Gradient descent vs ecuación normal
**Gradient descent:** iterativo, w ← w − α · ∂L/∂w
**Ecuación normal:** directo, w* = (XᵀX)⁻¹Xᵀy

Completar la tabla comparativa:

```
                    | Gradient Descent | Ecuación Normal
--------------------|------------------|----------------
Complejidad         | O(n·d·epochs)    | O(n·d² + d³)
Funciona para d>10k |                  |
Necesita lr tuning  |                  |
Generaliza a DL     |                  |
Exacto              |                  |

```

---

## Criterio de éxito
- ∂L/∂w = (2/n)(XᵀXw − Xᵀy)
- w* de Parte 4 = [1, 1] (intercept=1, slope=1)
- Entiendes por qué usar GD en DL aunque la ecuación normal sea más directa
