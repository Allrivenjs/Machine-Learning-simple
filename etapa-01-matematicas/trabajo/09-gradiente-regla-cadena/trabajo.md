# 09 — Gradiente y Regla de la Cadena

**EJ:** EJ-01-006 (parte) + EJ-01-007 | **Prerequisito:** [08-derivadas-parciales](../08-derivadas-parciales/trabajo.md) | **Estado:** ☐ Pendiente

---

## Objetivo
Dominar la regla de la cadena para funciones compuestas de múltiples variables. Es exactamente lo que hace backpropagation: aplicar la regla de la cadena capa por capa.

## Revisar antes de empezar
- Math4ML — Sección 5.2.2 (chain rule)
- 3Blue1Brown — What is backpropagation really doing? (Neural Networks series)

---

## Tarea

### Parte 1: Regla de la cadena — repaso 1D
Si **h(x) = f(g(x))**: **h'(x) = f'(g(x)) · g'(x)**

Calcular la derivada de **h(x) = ln(sin(x²))**:

```
Identificar capas (de afuera hacia adentro):
  f(u) = ln(u)       → f'(u) = 
  g(u) = sin(u)      → g'(u) = 
  p(x) = x²          → p'(x) = 

Aplicar cadena: h'(x) = f'(g(p(x))) · g'(p(x)) · p'(x)
             = 

```

---

### Parte 2: Regla de la cadena multivariable
Si **z = f(x, y)** donde **x = g(t)** e **y = h(t)**:
**dz/dt = (∂f/∂x)(dx/dt) + (∂f/∂y)(dy/dt)**

Para **z = x² + 2xy**, **x = t²**, **y = 3t**:

```
∂z/∂x = 
∂z/∂y = 
dx/dt = 
dy/dt = 

dz/dt = 

Verificar sustituyendo primero:
z = (t²)² + 2(t²)(3t) = 
dz/dt = 

¿Coincide?:

```

---

### Parte 3: Jacobiano
El **Jacobiano** J de **f: ℝⁿ → ℝᵐ** es la matriz de todas las derivadas parciales:

Jᵢⱼ = ∂fᵢ/∂xⱼ

Para **f(x,y) = [x² + y, xy − y²]**:

```
J = [[∂f₁/∂x, ∂f₁/∂y],
     [∂f₂/∂x, ∂f₂/∂y]]

  = [[       ,        ],
     [       ,        ]]

Evaluar J en (1, 2):

J(1,2) = 

```

---

### Parte 4: Red neuronal simple — backprop a mano
Red de 1 capa, 1 neurona:

```
x → [w, b] → z = wx + b → a = σ(z) → L = (a - y)²
```

Con **x = 1**, **y = 0**, **w = 0.5**, **b = 0**:

**Forward pass:**

```
z = wx + b = 
a = σ(z) = 1/(1 + e⁻ᶻ) = 
L = (a - y)² = 

```

**Backward pass** (regla de la cadena hacia atrás):

```
∂L/∂a = 

∂a/∂z = σ(z)(1-σ(z)) = 

∂z/∂w = 

∂z/∂b = 

∂L/∂w = ∂L/∂a · ∂a/∂z · ∂z/∂w = 

∂L/∂b = ∂L/∂a · ∂a/∂z · ∂z/∂b = 

```

**Update** con lr = 0.1:

```
w_nuevo = w - 0.1 · ∂L/∂w = 
b_nuevo = b - 0.1 · ∂L/∂b = 

```

---

### Parte 5: Verificar numéricamente (diferencias finitas)
Aproximar ∂L/∂w numéricamente: **(L(w+ε) − L(w−ε)) / (2ε)** con ε = 0.001

```
L(w + 0.001) = (con w=0.501)
L(w - 0.001) = (con w=0.499)

∂L/∂w ≈ 

¿Coincide con el resultado analítico de Parte 4?:

```

Este truco (gradient check) se usa para verificar implementaciones de backprop.

---

## Criterio de éxito
- Backprop de Parte 4: ∂L/∂w ≈ 0.12 (aprox)
- La verificación numérica coincide con el gradiente analítico (diferencia < 1e-5)
- Puedes hacer backprop de memoria para una red de 1 neurona
