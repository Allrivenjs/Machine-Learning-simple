# Ejercicios: Etapa 6 — Redes Neuronales desde Cero

Sin librerías de deep learning. Solo tu linalg de Etapa 2.

---

## Una neurona

### EJ-06-001: Activaciones
```go
func Sigmoid(z float64) float64
func SigmoidDerivative(z float64) float64

func ReLU(z float64) float64
func ReLUDerivative(z float64) float64

func Tanh(z float64) float64
func TanhDerivative(z float64) float64

func Softmax(z Vector) Vector
```
Test: Sigmoid(0)=0.5, ReLU(-1)=0, ReLU(3)=3

### EJ-06-002: Una neurona completa
Implementar una neurona que:
1. Recibe input x (vector)
2. Calcula z = w·x + b
3. Calcula a = σ(z)
4. Calcula pérdida MSE
5. Backpropaga: calcula ∂L/∂w y ∂L/∂b
6. Actualiza w y b con SGD

---

## Red Neuronal

### EJ-06-003: Capa densa
```go
type Dense struct {
    Weights Matrix  // shape: [out_size, in_size]
    Biases  Vector  // shape: [out_size]
    Activation string
}

func (d *Dense) Forward(input Matrix) Matrix  // batch forward
func (d *Dense) Backward(gradOutput Matrix) Matrix  // retorna gradInput
```

### EJ-06-004: Red completa
```go
type Network struct {
    Layers    []*Dense
    LossFunc  string
}

func (n *Network) Forward(X Matrix) Matrix
func (n *Network) Backward(yTrue Matrix) // calcula todos los gradientes
func (n *Network) UpdateWeights(lr float64)
func (n *Network) Train(X, y Matrix, epochs int, lr float64, batchSize int)
```

### EJ-06-005: XOR
Construir red 2→4→1 y resolver XOR:
```
Input  | Expected
[0, 0] | 0
[0, 1] | 1
[1, 0] | 1
[1, 1] | 0
```
Debe converger en <1000 epochs con lr=0.1

---

## Optimizadores

### EJ-06-006: SGD con Momentum
```go
type SGDMomentum struct {
    LR       float64
    Momentum float64
    velocity map[string]interface{} // velocidades para W y b
}
```

### EJ-06-007: Adam
```go
type Adam struct {
    LR      float64
    Beta1   float64  // 0.9
    Beta2   float64  // 0.999
    Epsilon float64  // 1e-8
    t       int      // timestep
    m, v    map[string]interface{}
}
```
Comparar curvas de entrenamiento: SGD vs SGD+Momentum vs Adam en XOR.

---

## MNIST

### EJ-06-008: Cargar MNIST
Descargar MNIST en formato CSV o IDX. Implementar loader:
```go
func LoadMNIST(imgPath, labelPath string) (Matrix, Vector, error)
```

### EJ-06-009: MLP para MNIST
Arquitectura: 784 → 128 → 64 → 10
- Activación ocultas: ReLU
- Activación salida: Softmax
- Pérdida: Cross-Entropy
- Optimizador: Adam

Meta: >95% accuracy en test set.

### EJ-06-010: Visualizar entrenamiento
Loggear por epoch:
```
Epoch 10/100 | loss: 0.2341 | train_acc: 92.3% | val_acc: 91.8%
```
Guardar en CSV para graficar con Python.

---

## Proyecto integrador

Implementar Dropout:
```go
type Dropout struct {
    Rate float64  // probabilidad de apagar neurona
    Mask Matrix   // generada en forward
}

func (d *Dropout) Forward(input Matrix, training bool) Matrix
func (d *Dropout) Backward(gradOutput Matrix) Matrix
```
Comparar MNIST con y sin Dropout (Rate=0.5 en capas ocultas). 
¿Mejora la generalización en test?
