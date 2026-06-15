# 08 — Derivadas Parciales

**EJ:** EJ-01-008 | **Prerequisito:** [07-derivadas](../07-derivadas/trabajo.md) | **Estado:** ☐ Pendiente

---

## Objetivo
Las derivadas parciales extienden las derivadas a funciones de múltiples variables. En ML, los parámetros de un modelo son vectores de miles de variables — necesitamos derivar respecto a cada una.

## Revisar antes de empezar
- Khan Academy — Partial derivatives (gratuito)
- Math4ML — Sección 5.2

---

## Tarea

### Parte 1: Primera derivada parcial
Para calcular ∂f/∂x, tratar y como constante y derivar respecto a x.

Para **f(x, y) = x²y + 3xy²**:

```
∂f/∂x = (tratar y como constante)

∂f/∂y = (tratar x como constante)

```

Para **f(x, y) = eˣʸ**:

```
∂f/∂x = 

∂f/∂y = 

```

---

### Parte 2: El gradiente
El **gradiente** ∇f agrupa todas las derivadas parciales en un vector:

∇f(x, y) = [∂f/∂x, ∂f/∂y]

Para **f(x, y) = x² + y²** (paraboloide):

Calcular ∇f:

```
∇f(x,y) = 

```

Evaluar ∇f en el punto (1, 2):

```
∇f(1,2) = 

```

¿Hacia dónde apunta el gradiente en (1,2)? ¿En qué dirección crece f más rápido?

```
Tu respuesta:

```

¿Dónde es ∇f = [0,0]? ¿Qué significa ese punto?

```
Tu respuesta:

```

---

### Parte 3: Caso 3 variables
Para **f(x, y, z) = x²y + yz + z²x**:

```
∂f/∂x = 

∂f/∂y = 

∂f/∂z = 

∇f = 

```

---

### Parte 4: Gradiente de una función de pérdida simple
Para **L(w₀, w₁) = (w₀ + w₁ · 2 − 5)² + (w₀ + w₁ · 4 − 7)²**

(pérdida para dos puntos: (2,5) y (4,7))

Calcular ∇L respecto a w₀ y w₁:

```
∂L/∂w₀ = 

∂L/∂w₁ = 

∇L = 

```

Si w₀=0, w₁=0, ¿cuánto vale ∇L? ¿En qué dirección debemos mover los parámetros?

```
∇L(0,0) = 

Dirección de update (−∇L) = 

```

---

### Parte 5: Conexión entre ∂f/∂x y regla de la cadena multivariable
Si **z = f(x(t), y(t))**, entonces: **dz/dt = ∂f/∂x · dx/dt + ∂f/∂y · dy/dt**

Para **z = x² + y²**, **x(t) = cos(t)**, **y(t) = sin(t)**:

Calcular dz/dt de dos formas:
1. Usando la regla de la cadena multivariable
2. Sustituyendo x,y primero y luego derivando

```
Método 1 (regla de cadena):

Método 2 (sustituir primero):
z = cos²(t) + sin²(t) = 
dz/dt = 

¿Coinciden?:

```

---

## Criterio de éxito
- ∇(x²+y²) = [2x, 2y]
- El gradiente apunta en la dirección de mayor crecimiento
- Puedes calcular ∇L para una función de pérdida de 2 parámetros
