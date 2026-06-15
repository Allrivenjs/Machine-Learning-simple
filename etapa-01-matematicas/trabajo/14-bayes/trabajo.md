# 14 — Teorema de Bayes

**EJ:** EJ-01-010 | **Prerequisito:** [12-distribuciones](../12-distribuciones/trabajo.md) + [13-esperanza-varianza-correlacion](../13-esperanza-varianza-correlacion/trabajo.md) | **Estado:** ☐ Pendiente

---

## Objetivo
Bayes es la base del razonamiento probabilístico: actualizar creencias con evidencia. Es el fundamento de Naive Bayes, estimación MAP y muchos modelos bayesianos.

## Revisar antes de empezar
- Math4ML — Sección 6.3
- 3Blue1Brown — Bayes theorem (video independiente)

---

## Tarea

### Parte 1: Probabilidad condicional
**P(A|B) = P(A ∩ B) / P(B)**

En un dataset de 100 personas:
- 60 son programadores
- 40 son diseñadores
- 20 programadores usan Mac
- 30 diseñadores usan Mac

```
P(Mac) = 
P(Mac | programador) = 
P(programador | Mac) = (calcular con Bayes)

```

---

### Parte 2: Derivar el Teorema de Bayes
Partiendo de las definiciones de probabilidad condicional:

```
P(A|B) = P(A ∩ B) / P(B)    ...(1)
P(B|A) = P(B ∩ A) / P(A)    ...(2)

Como P(A ∩ B) = P(B ∩ A), despejar P(A ∩ B) de (2) y sustituir en (1):

P(A|B) = 

```

---

### Parte 3: Diagnóstico médico
Un test de cáncer tiene:
- P(positivo | cáncer) = 0.99  (sensibilidad)
- P(positivo | sano) = 0.01   (tasa de falsos positivos)
- P(cáncer) = 0.001            (prevalencia)

Calcular P(cáncer | positivo):

```
P(cáncer | positivo) = P(positivo | cáncer) · P(cáncer) / P(positivo)

P(positivo) = P(positivo | cáncer) · P(cáncer) + P(positivo | sano) · P(sano)
            = 0.99 · 0.001 + 0.01 · 0.999
            = 

P(cáncer | positivo) = 

```

¿Sorprendente el resultado? Explicar la intuición:

```
Tu explicación:

```

---

### Parte 4: Terminología bayesiana
Identificar cada término en el contexto de ML:

**P(θ | datos) = P(datos | θ) · P(θ) / P(datos)**

```
P(θ)         = (prior)    →
P(datos | θ) = (likelihood) →
P(θ | datos) = (posterior)  →
P(datos)     = (evidence)   →

```

¿Qué pasa cuando el prior P(θ) es muy informativo? ¿Y cuando es "flat" (uniforme)?

```
Prior informativo:

Prior flat:

```

---

### Parte 5: Naive Bayes — clasificación de texto
Para clasificar un email como spam o no, Naive Bayes calcula:

**P(spam | palabras) ∝ P(spam) · P(palabra₁ | spam) · P(palabra₂ | spam) · ...**

La asunción "naive": las palabras son **condicionalmente independientes** dado la clase.

Dado:
- P(spam) = 0.3
- P("gratis" | spam) = 0.8, P("gratis" | no_spam) = 0.05
- P("reunión" | spam) = 0.1, P("reunión" | no_spam) = 0.4

Email: "gratis reunión gratis"

```
P(spam | email) ∝ P(spam) · P("gratis"|spam)² · P("reunión"|spam)
                = 0.3 · (0.8)² · 0.1
                = 

P(no_spam | email) ∝ P(no_spam) · P("gratis"|no_spam)² · P("reunión"|no_spam)
                   = 0.7 · (0.05)² · 0.4
                   = 

Normalizar:
P(spam | email) = 

¿Es spam?:

```

---

## Criterio de éxito
- P(cáncer | positivo) ≈ 0.09 (solo ~9%, aunque el test sea 99% preciso)
- Email de Parte 5 clasificado como spam
- Entiendes por qué prevalencia baja hace que aún con test preciso, muchos positivos sean falsos positivos
