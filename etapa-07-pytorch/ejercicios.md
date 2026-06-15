# Ejercicios: Etapa 7 — PyTorch

---

## Tensores y Autograd

### EJ-07-001: Autograd manual vs PyTorch
Implementar regresión lineal de dos formas y verificar que los gradientes son idénticos:
```python
import torch

# Forma 1: tu implementación manual (numpy)
# Forma 2: PyTorch autograd
x = torch.tensor([1.0, 2.0, 3.0], requires_grad=False)
y = torch.tensor([2.0, 4.0, 6.0])
w = torch.tensor(0.0, requires_grad=True)
b = torch.tensor(0.0, requires_grad=True)

pred = w * x + b
loss = ((pred - y) ** 2).mean()
loss.backward()

print(w.grad)  # debe coincidir con tu cálculo manual
```

### EJ-07-002: Custom autograd function
Implementar Sigmoid como `torch.autograd.Function`:
```python
class MySigmoid(torch.autograd.Function):
    @staticmethod
    def forward(ctx, x):
        ...
    @staticmethod
    def backward(ctx, grad_output):
        ...
```

---

## MLP en PyTorch

### EJ-07-003: MNIST con nn.Module
```python
class MLP(nn.Module):
    def __init__(self):
        super().__init__()
        self.layers = nn.Sequential(
            nn.Linear(784, 128),
            nn.ReLU(),
            nn.Linear(128, 64),
            nn.ReLU(),
            nn.Linear(64, 10)
        )
    
    def forward(self, x):
        return self.layers(x)
```
Meta: >98% en MNIST. Comparar velocidad vs tu implementación Go de Etapa 6.

### EJ-07-004: Training loop completo
```python
def train(model, loader, optimizer, criterion, epochs):
    for epoch in range(epochs):
        for X_batch, y_batch in loader:
            optimizer.zero_grad()
            preds = model(X_batch)
            loss = criterion(preds, y_batch)
            loss.backward()
            optimizer.step()
```
Agregar: logging de loss/accuracy, early stopping, guardado de checkpoints.

---

## CNN

### EJ-07-005: CNN para CIFAR-10
```python
class CNN(nn.Module):
    def __init__(self):
        super().__init__()
        self.conv = nn.Sequential(
            nn.Conv2d(3, 32, kernel_size=3, padding=1),
            nn.ReLU(),
            nn.MaxPool2d(2),
            nn.Conv2d(32, 64, kernel_size=3, padding=1),
            nn.ReLU(),
            nn.MaxPool2d(2),
        )
        self.fc = nn.Sequential(
            nn.Linear(64 * 8 * 8, 256),
            nn.ReLU(),
            nn.Linear(256, 10)
        )
    
    def forward(self, x):
        return self.fc(self.conv(x).flatten(1))
```
Meta: >80% accuracy en CIFAR-10 test.

### EJ-07-006: Visualizar feature maps
Para la primera capa conv de tu CNN entrenada:
- Visualizar los 32 filtros aprendidos
- Visualizar los feature maps de una imagen de entrada

---

## RNN / LSTM

### EJ-07-007: Character-level language model
Dataset: texto de un libro (ej. Don Quijote en txt).
```python
class CharRNN(nn.Module):
    def __init__(self, vocab_size, hidden_size):
        super().__init__()
        self.embed = nn.Embedding(vocab_size, hidden_size)
        self.rnn = nn.LSTM(hidden_size, hidden_size, batch_first=True)
        self.fc = nn.Linear(hidden_size, vocab_size)
    
    def forward(self, x, hidden=None):
        emb = self.embed(x)
        out, hidden = self.rnn(emb, hidden)
        return self.fc(out), hidden
```
Después de entrenamiento, generar texto carácter a carácter.

### EJ-07-008: Vanishing gradient
Entrenar RNN y LSTM en secuencias de longitud 10, 50, 100.
- Plotear la norma del gradiente por layer
- Demostrar empiricamente por qué LSTM > RNN en secuencias largas

---

## Proyecto integrador

Transfer Learning con ResNet pretrained:
```python
import torchvision.models as models

resnet = models.resnet18(pretrained=True)
# Congelar todos los pesos
for param in resnet.parameters():
    param.requires_grad = False

# Reemplazar la última capa
resnet.fc = nn.Linear(512, num_classes)
```
Fine-tuning en un dataset pequeño de tu elección.
Comparar: entrenar desde cero vs transfer learning.
