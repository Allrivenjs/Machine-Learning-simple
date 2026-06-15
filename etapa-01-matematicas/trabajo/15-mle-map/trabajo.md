# 15 — MLE y MAP

**EJ:** EJ-01-011 | **Prerequisito:** [14-bayes](../14-bayes/trabajo.md) + [09-gradiente-regla-cadena](../09-gradiente-regla-cadena/trabajo.md) | **Estado:** ☐ Pendiente

---

## Objetivo
MLE y MAP son los principios que justifican por qué usamos ciertas funciones de pérdida. MSE viene de MLE bajo ruido Gaussiano. L2 regularización viene de prior Gaussiano sobre los pesos (MAP).

## Revisar antes de empezar
- Math4ML — Sección 8.3
- ISLR — Sección 6.2 (Ridge como MAP)

---

## Tarea

### Parte 1: Maximum Likelihood Estimation (MLE)
MLE encuentra los parámetros θ que **maximizan la probabilidad de observar los datos**.

**θ_MLE = argmax_θ P(datos | θ) = argmax_θ L(θ)**

Para N observaciones iid: **L(θ) = Π P(xᵢ | θ)**

Para evitar underflow numérico se maximiza el **log-likelihood**:
**ℓ(θ) = Σ log P(xᵢ | θ)**

---

### Parte 2: MLE para la Normal
Dado dataset **{x₁, ..., xₙ}** asumiendo xᵢ ~ N(μ, σ²):

**ℓ(μ, σ²) = Σ log p(xᵢ)**

Escribir el log-likelihood expandido:

```
log p(xᵢ) = log[(1/√(2πσ²)) · exp(−(xᵢ−μ)²/(2σ²))]
           = 

ℓ(μ, σ²) = Σ log p(xᵢ) = 

```

Maximizar respecto a μ (∂ℓ/∂μ = 0):

```
∂ℓ/∂μ = 

Igualando a 0:

μ_MLE = 

```

¿Coincide con la media muestral? ¿Por qué tiene sentido?

```
Tu respuesta:

```

Aplicar al dataset **{2, 4, 4, 4, 5, 5, 7, 9}**:

```
μ_MLE = 
σ²_MLE = (1/n) Σ(xᵢ - μ)² =    ← nota: divide por n, no n-1

```

---

### Parte 3: MLE → MSE (regresión lineal)
Si asumimos **yᵢ = wᵀxᵢ + εᵢ** donde εᵢ ~ N(0, σ²):

Entonces **yᵢ | xᵢ ~ N(wᵀxᵢ, σ²)**

Log-likelihood:

```
ℓ(w) = Σ log p(yᵢ | xᵢ, w)
      = Σ log N(yᵢ; wᵀxᵢ, σ²)
      = (expandir el log de la Gaussiana)
      = Constante − (1/2σ²) Σ(yᵢ − wᵀxᵢ)²

```

Maximizar ℓ(w) es equivalente a minimizar:

```
¿Qué función de pérdida conocida resulta?

```

**Conclusión:** MSE no es arbitrario — es la pérdida correcta si el ruido es Gaussiano.

---

### Parte 4: MAP — incorporar prior
**θ_MAP = argmax_θ P(θ | datos) = argmax_θ [P(datos | θ) · P(θ)]**

Si asumimos un **prior Gaussiano** sobre los pesos: w ~ N(0, τ²I)

El log del posterior es:

```
log P(w | datos) ∝ ℓ(w) + log P(w)
                 ∝ −(1/2σ²) Σ(yᵢ − wᵀxᵢ)² − (1/2τ²) ‖w‖²

```

Maximizar esto es equivalente a minimizar:

```
¿Qué función de pérdida con regularización resulta?

λ = σ²/τ²

```

**Conclusión:** Ridge Regression (L2) = MLE + prior Gaussiano sobre pesos.

---

### Parte 5: Comparación MLE vs MAP
Completar:

```
                    | MLE              | MAP
--------------------|------------------|------------------
Usa prior           |                  |
Puede overfittear   |                  |
Con datos → ∞       |                  | → MLE (prior se diluye)
Regularización      | ninguna          |
Ejemplo en ML       | regresión sin reg| 

```

¿Cuándo usarías MAP sobre MLE?

```
Tu respuesta (pista: pocos datos, quieres regularizar):

```

---

## Criterio de éxito
- MLE de Normal da media muestral (μ̂ = x̄)
- MLE para regresión lineal → MSE
- MAP con prior Gaussiano → Ridge Regression
- μ_MLE para {2,4,4,4,5,5,7,9} = 5.0
