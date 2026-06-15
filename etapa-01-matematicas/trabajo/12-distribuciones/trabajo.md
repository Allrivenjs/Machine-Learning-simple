# 12 — Distribuciones de Probabilidad

**EJ:** EJ-01-012 | **Prerequisito:** ninguno (paralelo a cálculo) | **Estado:** ☐ Pendiente

---

## Objetivo
Las distribuciones modelan la incertidumbre. En ML las usamos para modelar el ruido en los datos, las predicciones como probabilidades y las funciones de pérdida como likelihoods.

## Revisar antes de empezar
- ISLR — Sección 2.1 (statistical learning setup)
- Math4ML — Sección 6.1, 6.2

---

## Tarea

### Parte 1: Distribución Normal (Gaussiana)
Fórmula: **p(x) = (1/√(2πσ²)) · exp(−(x−μ)²/(2σ²))**

Parámetros:
- μ = media (centro)
- σ² = varianza (dispersión)

**a)** Para N(0,1) (estándar), ¿cuánto vale p(0)?

```
p(0) = 

```

**b)** Describir cómo cambia la curva cuando σ crece:

```
Tu respuesta:

```

**c)** Nombrar 3 fenómenos reales que se aproximan bien con una Normal:

```
1. 
2. 
3. 
```

**d)** ¿Por qué MSE como función de pérdida asume ruido Gaussiano? (pista: MLE bajo Normal da MSE)

```
Tu intuición:

```

---

### Parte 2: Distribución Bernoulli
Modela experimentos binarios (éxito/fracaso).

**p(X=1) = p**, **p(X=0) = 1−p**

Compacto: **p(x) = pˣ(1−p)¹⁻ˣ** para x ∈ {0, 1}

**a)** Si p=0.7, ¿cuánto vale p(X=1) y p(X=0)?

```
p(X=1) = 
p(X=0) = 

```

**b)** ¿Qué representa p en un clasificador logístico?

```
Tu respuesta:

```

**c)** ¿Por qué Binary Cross-Entropy es la pérdida natural para clasificación binaria? (está relacionada con Bernoulli likelihood)

```
Tu intuición:

```

---

### Parte 3: Distribución Poisson
Modela conteos de eventos en un intervalo: **p(X=k) = (λᵏ e⁻λ) / k!**

λ = tasa esperada de eventos.

**a)** Si λ=3 llamadas/hora, ¿cuál es la probabilidad de exactamente 0 llamadas en una hora?

```
p(X=0) = 

```

**b)** ¿cuál es la probabilidad de exactamente 3 llamadas?

```
p(X=3) = 

```

**c)** Nombrar 2 problemas de ML donde Poisson es adecuado:

```
1. 
2. 
```

---

### Parte 4: Distribución Binomial
n experimentos Bernoulli, X = número de éxitos.

**p(X=k) = C(n,k) · pᵏ · (1−p)ⁿ⁻ᵏ**

Con n=10, p=0.5 (moneda justa):

**a)** p(X=5) = ¿exactamente 5 caras en 10 lanzamientos?

```
C(10,5) = 
p(X=5) = 

```

**b)** ¿Qué pasa con la distribución Binomial cuando n→∞ y np=λ constante?

```
Tu respuesta:

```

---

### Parte 5: Elegir la distribución correcta
Emparejar cada situación con su distribución:

```
Situación                                    | Distribución
---------------------------------------------|-------------
Altura de personas en una población          | 
Si un email es spam o no (0/1)               | 
Número de clicks por hora en un sitio web    | 
Número de éxitos en 20 lanzamientos de dado  | 
Error de predicción de un modelo             | 
```

---

## Criterio de éxito
- p(0) para N(0,1) ≈ 0.399
- Poisson(3): p(X=0) ≈ 0.0498, p(X=3) ≈ 0.224
- Binomial(10, 0.5): p(X=5) ≈ 0.246
- Entiendes por qué MSE ↔ Normal y Cross-Entropy ↔ Bernoulli
