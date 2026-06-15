# Etapa 7: PyTorch

Duración: 1 mes | Prerequisito: Etapa 6

---

## Objetivo

Aprender PyTorch entendiéndolo desde adentro. Ya sabes cómo funciona backprop — ahora aprendes cómo PyTorch lo automatiza con autograd.

---

## Recursos (en orden)

1. **PyTorch documentation** — tutorials oficiales: pytorch.org/tutorials
2. **Dive Into Deep Learning** — Capítulos 5-9 (d2l.ai, usa PyTorch)
3. **Andrej Karpathy — Neural Networks: Zero to Hero** (YouTube) — imprescindible
4. **Hands-On ML** (Géron) — Capítulos 10-12

---

## Temas a dominar

### Tensores
- [ ] Diferencias vs NumPy ndarray
- [ ] `torch.tensor`, `torch.zeros`, `torch.randn`
- [ ] Device: CPU vs CUDA
- [ ] `requires_grad=True`
- [ ] Operaciones: `@`, `sum`, `mean`, `view`, `reshape`

### Autograd
- [ ] Computation graph
- [ ] `tensor.backward()`
- [ ] `tensor.grad`
- [ ] `torch.no_grad()` — para inference
- [ ] Entender por qué autograd elimina backprop manual

### nn.Module
- [ ] `nn.Linear`, `nn.ReLU`, `nn.Sigmoid`
- [ ] `nn.Sequential`
- [ ] `model.parameters()`
- [ ] `model.train()` vs `model.eval()`
- [ ] Custom Module: heredar `nn.Module`, implementar `forward`

### Optimizadores
- [ ] `optim.SGD`
- [ ] `optim.Adam`
- [ ] `optimizer.zero_grad()`, `optimizer.step()`
- [ ] Learning rate schedulers

### Datos
- [ ] `Dataset` — clase abstracta
- [ ] `DataLoader` — batching, shuffle, workers
- [ ] `torchvision.datasets`: MNIST, CIFAR-10
- [ ] `transforms.Compose`

### GPU
- [ ] `tensor.to('cuda')`
- [ ] `model.cuda()`
- [ ] Por qué GPU acelera: paralelismo masivo

---

## Arquitecturas a construir

### MLP
- [ ] Reproducir MNIST de Etapa 6 en PyTorch — debe ser más rápido y corto

### CNN (Convolutional Neural Network)
- [ ] `nn.Conv2d`, `nn.MaxPool2d`
- [ ] Entender receptive field
- [ ] Aplicar en CIFAR-10

### RNN / LSTM
- [ ] `nn.RNN`, `nn.LSTM`, `nn.GRU`
- [ ] Hidden state y cell state
- [ ] Problema: sequence modeling
- [ ] Aplicación: predecir siguiente carácter en texto

---

## Señales de que dominas la etapa

- Puedes implementar cualquier arquitectura del paper sin buscar tutoriales
- Entiendes qué hace `optimizer.zero_grad()` y por qué es necesario
- Sabes depurar: `tensor.shape`, `tensor.dtype`, NaN en gradientes
- Tu CNN supera 80% en CIFAR-10
