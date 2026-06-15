# 07 — Derivadas

**EJ:** EJ-01-009 (parte) | **Prerequisito:** ninguno (paralelo a álgebra) | **Estado:** ☐ Pendiente

---

## Objetivo
Las derivadas miden cómo cambia una función. En ML, las usamos para encontrar el mínimo de la función de pérdida. Este es el punto de inicio del cálculo para ML.

## Revisar antes de empezar
- Khan Academy — Basic derivatives (gratuito)
- Math4ML — Sección 5.1, 5.2

---

## Tarea

### Parte 1: Derivadas básicas (reglas de memoria)
Completar la tabla:

```
f(x)         | f'(x)
-------------|--------
xⁿ           | 
eˣ           | 
ln(x)        | 
sin(x)       | 
cos(x)       | 
constante c  | 
```

---

### Parte 2: Calcular derivadas
**a) f(x) = 3x² + 2x − 5**

```
f'(x) = 
```

**b) f(x) = eˣ · x²** (regla del producto: (uv)' = u'v + uv')

```
f'(x) = 
```

**c) f(x) = ln(x² + 1)** (regla de la cadena)

```
f'(x) = 
```

**d) f(x) = (2x + 3)⁴** (regla de la cadena)

```
f'(x) paso a paso:

Sea u = 2x + 3, entonces f = u⁴
df/dx = df/du · du/dx = 

```

---

### Parte 3: Regla de la cadena — foco
La regla de la cadena es **crítica** para backpropagation.

Si **h(x) = f(g(x))**, entonces **h'(x) = f'(g(x)) · g'(x)**

Calcular la derivada de **h(x) = (3x² + 2)⁵**:

```
Identificar:
  g(x) = 
  f(u) = 

  g'(x) = 
  f'(u) = 

  h'(x) = f'(g(x)) · g'(x) = 

```

Calcular la derivada de **h(x) = σ(x) = 1/(1 + e⁻ˣ)** (función sigmoid):

```
Reescribir: σ(x) = (1 + e⁻ˣ)⁻¹

Aplicar regla de la cadena:

σ'(x) = 

Resultado elegante: σ'(x) = σ(x) · (1 − σ(x))

Verificar evaluando en x=0:
σ(0) = 
σ'(0) = 

```

---

### Parte 4: Encontrar mínimos
Para **f(x) = x² − 4x + 5**:

Igualando f'(x) = 0:

```
f'(x) = 
Igualando a 0: 
x* = 
f(x*) = 

```

Verificar que es mínimo (no máximo): **f''(x*) > 0**

```
f''(x) = 
f''(x*) = 
¿Es mínimo?: 

```

---

### Parte 5: Conexión con ML — gradient descent
Gradient descent actualiza parámetros en dirección opuesta al gradiente:
**w ← w − α · f'(w)**

Para **f(w) = (2w − 4)²** (pérdida simple), simular 3 pasos con w₀ = 0 y α = 0.1:

```
f'(w) = 

Paso 1: w₁ = w₀ - 0.1 · f'(w₀) = 
Paso 2: w₂ = w₁ - 0.1 · f'(w₁) = 
Paso 3: w₃ = w₂ - 0.1 · f'(w₂) = 

¿Hacia dónde converge w?: 

```

---

## Criterio de éxito
- Derivada de σ(x) = σ(x)(1−σ(x))
- Mínimo de x²−4x+5 está en x=2
- Gradient descent en Parte 5 converge hacia w=2
