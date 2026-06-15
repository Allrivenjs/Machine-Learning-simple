# Etapa 8: Transformers

Duración: 1 mes | Prerequisito: Etapa 7

---

## Objetivo

Entender la arquitectura que domina el ML moderno. Implementar un GPT pequeño desde cero en PyTorch.

---

## Recursos (en orden)

1. **Paper original**: "Attention Is All You Need" (Vaswani et al., 2017) — leer primero
2. **Andrej Karpathy — Let's build GPT** (YouTube, 2h) — imprescindible
3. **The Annotated Transformer** — harvardnlp.github.io/annotated-transformer
4. **Dive Into Deep Learning** — Capítulo 11
5. **BERT paper**: "BERT: Pre-training of Deep Bidirectional Transformers" (Devlin et al.)

---

## Temas a dominar

### Attention
- [ ] Intuición: qué pregunta resuelve attention
- [ ] Scaled Dot-Product Attention:
  ```
  Attention(Q,K,V) = softmax(QKᵀ / √d_k) V
  ```
- [ ] Por qué dividir por √d_k
- [ ] Self-attention vs cross-attention

### Multi-Head Attention
- [ ] Por qué múltiples heads
- [ ] Proyecciones W_Q, W_K, W_V
- [ ] Concatenación y proyección final W_O
- [ ] Implementación eficiente en batch

### Positional Encoding
- [ ] Por qué los transformers necesitan posición (vs RNN)
- [ ] Encoding sinusoidal: PE(pos, 2i) = sin(pos/10000^(2i/d))
- [ ] Positional embeddings aprendidos (GPT)

### Arquitectura completa
- [ ] Encoder (BERT): bidireccional
- [ ] Decoder (GPT): causal (solo atiende al pasado)
- [ ] Layer Normalization (Pre-LN vs Post-LN)
- [ ] Feed-Forward sublayer: 2 capas densas con ReLU
- [ ] Residual connections

### Pre-training y Fine-tuning
- [ ] Masked Language Modeling (BERT)
- [ ] Causal Language Modeling (GPT)
- [ ] Fine-tuning para clasificación
- [ ] LoRA (Low-Rank Adaptation) — fine-tuning eficiente

---

## Proyectos

### Proyecto 1: Implementar Attention desde cero
En PyTorch, sin `nn.MultiheadAttention`:
```python
def attention(Q, K, V, mask=None):
    d_k = Q.shape[-1]
    scores = Q @ K.transpose(-2, -1) / (d_k ** 0.5)
    if mask is not None:
        scores = scores.masked_fill(mask == 0, -1e9)
    return F.softmax(scores, dim=-1) @ V
```

### Proyecto 2: GPT pequeño (nanoGPT)
Basado en Karpathy's nanoGPT:
- Dataset: texto de Shakespeare o similar
- Modelo: 6 layers, 6 heads, d_model=384
- Generar texto coherente después de entrenamiento

---

## Señales de que dominas la etapa

- Puedes implementar multi-head attention sin copiar código
- Entiendes por qué causal masking es necesario en GPT
- Tu nanoGPT genera texto con estructura de frase reconocible
- Puedes explicar la diferencia entre BERT y GPT a alguien sin contexto técnico
