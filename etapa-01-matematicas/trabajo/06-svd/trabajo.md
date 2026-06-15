# 06 — SVD (Singular Value Decomposition)

**EJ:** EJ-01-005 | **Prerequisito:** [05-eigenvalores](../05-eigenvalores/trabajo.md) | **Estado:** ☐ Pendiente

---

## Objetivo
SVD generaliza la diagonalización a matrices no cuadradas. Es la base de PCA, compresión de imágenes, sistemas de recomendación y resolución de sistemas sobredeterminados.

## Revisar antes de empezar
- 3Blue1Brown — Singular Value Decomposition (video separado de la serie)
- Math4ML — Sección 2.7.2

---

## Tarea

### Parte 1: La descomposición
SVD descompone cualquier matriz A (m×n) como: **A = UΣVᵀ**

Describir qué representa cada componente:

```
U (m×m): 

Σ (m×n): 

Vᵀ (n×n): 

```

¿Qué son los **valores singulares** (diagonal de Σ)?

```
Tu respuesta:

```

---

### Parte 2: Relación con eigenvalores
Los valores singulares de A son las raíces cuadradas de los eigenvalores de **AᵀA**.

Para **A = [[1, 1], [0, 1], [1, 0]]** (3×2):

Calcular **AᵀA**:

```
AᵀA = Aᵀ · A = 

```

Los eigenvalores de AᵀA son los σᵢ² (valores singulares al cuadrado).

Calcular det(AᵀA − λI) = 0:

```
Tu cálculo:

λ₁ = 
λ₂ = 

σ₁ = √λ₁ = 
σ₂ = √λ₂ = 

```

---

### Parte 3: SVD truncada — compresión
Una imagen en escala de grises es una matriz A (m×n).

SVD permite aproximarla usando solo los k valores singulares más grandes:
**Aₖ = U[:, :k] · Σ[:k, :k] · Vᵀ[:k, :]**

Explicar:

**¿Por qué los primeros valores singulares capturan la "información más importante"?**

```
Tu respuesta:

```

**¿Cuántos números necesitas guardar para una imagen 100×100 con k=10 vs guardando la imagen completa?**

```
Con k=10: U(100×10) + Σ(10) + Vᵀ(10×100) = 
Imagen completa: 100×100 = 
Factor de compresión: 

```

---

### Parte 4: SVD y PCA
PCA usando SVD:
1. Centrar los datos: X_c = X − mean(X)
2. Calcular SVD: X_c = UΣVᵀ
3. Las componentes principales son las columnas de V
4. Los scores son U·Σ

¿Por qué preferir SVD sobre calcular la matriz de covarianza + eigenvalores?

```
Tu respuesta (pistas: numericamente más estable, funciona para m < n):

```

---

### Parte 5: Intuición geométrica
SVD dice que cualquier transformación lineal se puede descomponer en:
1. **Rotación** (Vᵀ)
2. **Escalado** (Σ)
3. **Rotación** (U)

Dibujar (o describir) qué hace A = UΣVᵀ a un círculo unitario:

```
Descripción:
Paso 1 (Vᵀ): 
Paso 2 (Σ): 
Paso 3 (U): 
Resultado final:

```

---

## Criterio de éxito
- Sabes qué shape tienen U, Σ, V para una matriz m×n
- Los valores singulares de [[1,1],[0,1],[1,0]] son σ₁ ≈ 1.73, σ₂ ≈ 1.0
- Puedes explicar por qué SVD permite comprimir imágenes
