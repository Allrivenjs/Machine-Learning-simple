# 02 — Matrices: Operaciones

**EJ:** EJ-01-002 | **Prerequisito:** [01-vectores](../01-vectores/trabajo.md) | **Estado:** ☐ Pendiente

---

## Objetivo
Entender la multiplicación matricial como composición de transformaciones. Es la operación central de forward propagation en redes neuronales.

## Revisar antes de empezar
- 3Blue1Brown — Linear transformations and matrices (capítulo 3)
- 3Blue1Brown — Matrix multiplication as composition (capítulo 4)
- Math4ML — Sección 2.2, 2.3

---

## Tarea

### Parte 1: Suma y resta
```
A = [[1, 2],    B = [[5, 6],
     [3, 4]]         [7, 8]]
```

**a) A + B**

```
Tu solución:
A + B = [[1+5, 2+6],
         [3+7, 4+8]] = [[6, 8],
                        [10, 12]]
```

**b) A − B**

```
Tu solución:
A - B = [[1-5, 2-6],
         [3-7, 4-8]] = [[-4, -4],
                        [-4, -4]]
```

---

### Parte 2: Multiplicación por escalar
**2 · A**

```
Tu solución:
A · 2 = [[1*2, 2*2],
           [3*2, 4*2]] = [[2, 4],
                          [6, 8]]
```

---

### Parte 3: Multiplicación matricial
Calcular **AB** y **BA** para A y B de arriba.

Recordar: (AB)ᵢⱼ = Σₖ Aᵢₖ · Bₖⱼ

**AB:**

```
Tu solución (paso a paso, calcula cada elemento):

C[0,0] = A[0,0]*B[0,0] + A[0,1]*B[1,0] = [1*5 + 2*7] = 19
C[0,1] = A[0,0]*B[0,1] + A[0,1]*B[1,1] = [1*6 + 2*8] = 22
C[1,0] = A[1,0]*B[0,0] + A[1,1]*B[1,0] = [3*5 + 4*7] = 43
C[1,1] = A[1,0]*B[0,1] + A[1,1]*B[1,1] = [3*6 + 4*8] = 50

Resultado:
AB = [[19, 22],
      [43, 50]]
```

**BA:**

```
Tu solución:
BA = [[5*1 + 6*3, 5*2 + 6*4],
       [7*1 + 8*3, 7*2 + 8*4]] = [[5+18, 10+24],
                                   [7+24, 14+32]] = [[23, 34],
                                                     [31, 46]]
```

**¿AB = BA?** ¿Por qué importa esto en ML?

```
Tu respuesta:
No, AB ≠ BA. Esto es importante porque el orden de las operaciones en 
redes neuronales afecta el resultado final. Por ejemplo, aplicar una 
transformación A seguida de B no es lo mismo que aplicar B seguida de A.
Ejemplo: si A es una rotación y B es una escala, el resultado será diferente dependiendo del orden.

```

---

### Parte 4: Transpuesta
Calcular **Aᵀ** y **Bᵀ**

```
Aᵀ: [[1, 3],
     [2, 4]]

Bᵀ: [[5, 7],
     [6, 8]]

```

Verificar la propiedad: **(AB)ᵀ = BᵀAᵀ**

```
Verificación:
(AB)ᵀ = [[19, 43],
         [22, 50]]ᵀ = [[19, 22],
                        [43, 50]]

BᵀAᵀ = [[5, 7],
        [6, 8]] · [[1, 3],
                   [2, 4]] = [[5*1 + 7*2, 5*3 + 7*4],
                              [6*1 + 8*2, 6*3 + 8*4]] = [[19, 22],
                                                          [43, 50]]

¿Son iguales?:
Sí, (AB)ᵀ = BᵀAᵀ, lo que confirma la propiedad de
la transpuesta en la multiplicación matricial. 
Esto es fundamental en ML para entender cómo se propagan los 
gradientes y cómo se calculan las derivadas en redes neuronales.
```

---

### Parte 5: Producto matriz-vector
Calcular **Av** donde **v = [1, 2]ᵀ**

```
Tu solución: 
Av = [[1, 2],
      [3, 4]] · [1,
                 2] = [[1*1 + 2*2],
                        [3*1 + 4*2]] = [[5],
                                        [11]]
```

**Interpretar geométricamente:** ¿qué hace A al vector v? (escalar, rotar, proyectar...)

```
Tu respuesta: 
A transforma el vector v aplicando una combinación lineal de sus columnas. 
En este caso, A está escalando y rotando el vector v. 
El resultado [5, 11] es una nueva posición en el espacio que 
representa la transformación de v por A.
```

---

### Parte 6: Caso práctico — Una capa de red neuronal
En una red neuronal, una capa hace: **z = Wx + b**

Dados:
- W (pesos) = [[0.5, 0.2], [0.1, 0.8], [0.3, 0.4]]  (3 neuronas, 2 inputs)
- x (input) = [1.0, 0.5]
- b (bias) = [0.1, 0.2, 0.3]

Calcular z:
z = Wx + b
z = [[0.5, 0.2], 
     [0.1, 0.8], 
     [0.3, 0.4]] · [1.0, 0.5] + [0.1, 0.2, 0.3]
Calculamos Wx:
W[0]·x = 0.5*1.0 + 0.2*0.5 = 0.5 + 0.1 = 0.6
W[1]·x = 0.1*1.0 + 0.8*0.5 = 0.1 + 0.4 = 0.5
W[2]·x = 0.3*1.0 + 0.4*0.5 = 0.3 + 0.2 = 0.5
Entonces, Wx = [0.6, 0.5, 0.5]
Finalmente, sumamos el bias:
z = [0.6 + 0.1, 0.5 + 0.2, 0.5 + 0.3] = [0.7, 0.7, 0.8]
z = [0.7, 0.7, 0.8]


```
Tu solución:

¿Qué shape tiene z?:

```

---

## Criterio de éxito
- AB ≠ BA (verifica que salen distintos)
- (AB)ᵀ = BᵀAᵀ (propiedad verificada)
- z en Parte 6 = [0.7, 0.55, 0.8]
