# Etapa 3: Machine Learning Clásico desde Cero

Duración: 2 meses | Prerequisito: Etapas 1 y 2

---

## Objetivo

Implementar cada algoritmo en Go desde cero antes de ver cómo Scikit-Learn lo hace. La comprensión viene de la implementación, no del uso de APIs.

---

## Recursos por algoritmo

### Regresión Lineal
- **ISLR** Capítulo 3
- **Hands-On ML** (Géron) Capítulo 4
- Video: StatQuest — "Linear Regression" (YouTube)

### Regresión Logística
- **ISLR** Capítulo 4
- **Hands-On ML** Capítulo 4
- Video: StatQuest — "Logistic Regression" (YouTube)

### KNN
- **ISLR** Capítulo 2
- Video: StatQuest — "K-Nearest Neighbors"

### Naive Bayes
- **Pattern Recognition and ML** (Bishop) Capítulo 4
- Video: StatQuest — "Naive Bayes"

### Árboles de Decisión
- **ISLR** Capítulo 8
- **Hands-On ML** Capítulo 6
- Video: StatQuest — "Decision Trees"

### SVM
- **ISLR** Capítulo 9
- **Pattern Recognition and ML** Capítulo 7
- Video: StatQuest — "Support Vector Machines"

---

## Temas a dominar

### Regresión Lineal
- [ ] Modelo: y = wᵀx + b
- [ ] MSE como función de pérdida
- [ ] Gradient descent
- [ ] Solución analítica (normal equation)
- [ ] Regularización: Ridge (L2), Lasso (L1)
- [ ] R² score

### Regresión Logística
- [ ] Función sigmoid: σ(z) = 1/(1+e⁻ᶻ)
- [ ] Binary cross-entropy loss
- [ ] Clasificación binaria
- [ ] Métricas: accuracy, precision, recall, F1, AUC-ROC
- [ ] Extensión multiclase: softmax

### KNN
- [ ] Distancia euclidiana
- [ ] K como hiperparámetro
- [ ] Votación por mayoría (clasificación)
- [ ] Media ponderada (regresión)
- [ ] KD-Trees para búsqueda eficiente

### Naive Bayes
- [ ] Teorema de Bayes aplicado
- [ ] Gaussian Naive Bayes (features continuas)
- [ ] Multinomial Naive Bayes (conteos, texto)
- [ ] Laplace smoothing

### Árboles de Decisión
- [ ] Entropía: H(S) = -Σ pᵢ log₂(pᵢ)
- [ ] Information Gain
- [ ] Índice Gini
- [ ] Criterio de parada (profundidad, min_samples)
- [ ] Poda (pruning)
- [ ] Random Forest: bagging + feature sampling

### SVM
- [ ] Hiperplano de separación
- [ ] Margen máximo
- [ ] Vectores de soporte
- [ ] Kernel trick: RBF, polinomial
- [ ] Soft margin (C como hiperparámetro)

---

## Conceptos transversales (críticos)
- [ ] Train/validation/test split
- [ ] Cross-validation (k-fold)
- [ ] Bias-variance tradeoff
- [ ] Overfitting y underfitting
- [ ] Hyperparameter tuning

---

## Señales de que dominas la etapa

- Puedes explicar por qué gradient descent converge hacia el mínimo
- Entiendes la diferencia entre parámetros e hiperparámetros
- Puedes elegir el algoritmo correcto dado un problema
- Tus implementaciones en Go producen resultados idénticos a Scikit-Learn
