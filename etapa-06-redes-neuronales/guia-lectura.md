# Etapa 6: Redes Neuronales desde Cero

Duración: 2 meses | Prerequisito: Etapas 1-5

---

## Objetivo

Construir una red neuronal completa en Go sin librerías. Entender cada operación matemática que ocurre durante forward pass y backpropagation.

---

## Recursos (en orden)

1. **3Blue1Brown — Neural Networks** (YouTube, 4 videos) — ver antes de leer
2. **Deep Learning** (Goodfellow) — Capítulos 6, 7, 8
3. **Dive Into Deep Learning** — Capítulos 4, 5 (gratuito: d2l.ai)
4. **CS231n Stanford** — Notas del curso (gratuitas online)
5. **Andrej Karpathy — micrograd** (GitHub) — implementación minimalista en Python para referencia

---

## Temas a dominar

### Una neurona
- [ ] Modelo: a = σ(Wx + b)
- [ ] Funciones de activación: Sigmoid, ReLU, Tanh, Softmax
- [ ] Derivadas de cada activación (necesario para backprop)

### Forward Propagation
- [ ] Capas densas (fully connected)
- [ ] Matrix de pesos por capa
- [ ] Propagación de activaciones

### Funciones de pérdida
- [ ] MSE (regresión)
- [ ] Binary Cross-Entropy (clasificación binaria)
- [ ] Categorical Cross-Entropy (multiclase)

### Backpropagation
- [ ] Regla de la cadena aplicada a redes
- [ ] Gradiente de la pérdida respecto a cada peso
- [ ] Acumulación de gradientes

### Optimizadores
- [ ] SGD (Stochastic Gradient Descent)
- [ ] Mini-batch gradient descent
- [ ] Momentum
- [ ] RMSProp
- [ ] Adam (crítico — el más usado)

### Regularización
- [ ] Dropout
- [ ] L2 weight decay
- [ ] Batch Normalization (conceptual)

### Técnicas de entrenamiento
- [ ] Learning rate scheduling
- [ ] Early stopping
- [ ] Weight initialization: Xavier, He

---

## Arquitecturas a construir

### MLP (Multi-Layer Perceptron)
- [ ] Resolver XOR (problema no linealmente separable)
- [ ] Clasificar MNIST (70k imágenes, 10 clases)

---

## Señales de que dominas la etapa

- Puedes derivar el algoritmo de backprop desde la regla de la cadena sin mirar notas
- Tu MLP alcanza >97% accuracy en MNIST
- Entiendes por qué ReLU > Sigmoid para redes profundas (vanishing gradient)
- Puedes explicar qué hace Adam diferente a SGD
