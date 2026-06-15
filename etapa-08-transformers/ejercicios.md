# Ejercicios: Etapa 8 — Transformers

---

## Attention

### EJ-08-001: Scaled Dot-Product Attention
```python
import torch
import torch.nn.functional as F

def scaled_dot_product_attention(Q, K, V, mask=None):
    """
    Q: (batch, heads, seq_len, d_k)
    K: (batch, heads, seq_len, d_k)
    V: (batch, heads, seq_len, d_v)
    """
    d_k = Q.shape[-1]
    scores = Q @ K.transpose(-2, -1) / (d_k ** 0.5)
    if mask is not None:
        scores = scores.masked_fill(mask == 0, float('-inf'))
    weights = F.softmax(scores, dim=-1)
    return weights @ V, weights
```
Test con datos dummy: verificar shapes de salida.

### EJ-08-002: Visualizar attention weights
Para una oración simple, visualizar la matriz de attention:
```python
import seaborn as sns
sentence = "el perro come huesos"
tokens = sentence.split()
# Correr tu attention
# Plotear weights como heatmap con tokens como ejes
```

---

## Multi-Head Attention

### EJ-08-003: MultiHeadAttention desde cero
```python
class MultiHeadAttention(nn.Module):
    def __init__(self, d_model, n_heads):
        super().__init__()
        self.d_k = d_model // n_heads
        self.n_heads = n_heads
        self.W_q = nn.Linear(d_model, d_model)
        self.W_k = nn.Linear(d_model, d_model)
        self.W_v = nn.Linear(d_model, d_model)
        self.W_o = nn.Linear(d_model, d_model)
    
    def forward(self, x, mask=None):
        B, T, C = x.shape
        # Proyectar, reshape a (B, n_heads, T, d_k), attention, concat, proyectar
        ...
```

---

## Positional Encoding

### EJ-08-004: Sinusoidal PE
```python
def positional_encoding(seq_len, d_model):
    PE = torch.zeros(seq_len, d_model)
    position = torch.arange(seq_len).unsqueeze(1)
    div_term = torch.exp(torch.arange(0, d_model, 2) * (-math.log(10000.0) / d_model))
    PE[:, 0::2] = torch.sin(position * div_term)
    PE[:, 1::2] = torch.cos(position * div_term)
    return PE
```
Visualizar el encoding como heatmap. ¿Puedes ver el patrón sinusoidal?

---

## Transformer Block

### EJ-08-005: Block completo
```python
class TransformerBlock(nn.Module):
    def __init__(self, d_model, n_heads, d_ff, dropout=0.1):
        super().__init__()
        self.attention = MultiHeadAttention(d_model, n_heads)
        self.norm1 = nn.LayerNorm(d_model)
        self.ff = nn.Sequential(
            nn.Linear(d_model, d_ff),
            nn.ReLU(),
            nn.Linear(d_ff, d_model)
        )
        self.norm2 = nn.LayerNorm(d_model)
        self.dropout = nn.Dropout(dropout)
    
    def forward(self, x, mask=None):
        # Pre-LN: norm antes de sublayer
        x = x + self.dropout(self.attention(self.norm1(x), mask))
        x = x + self.dropout(self.ff(self.norm2(x)))
        return x
```

---

## GPT pequeño

### EJ-08-006: nanoGPT
```python
class GPT(nn.Module):
    def __init__(self, vocab_size, d_model, n_heads, n_layers, max_len):
        super().__init__()
        self.token_emb = nn.Embedding(vocab_size, d_model)
        self.pos_emb = nn.Embedding(max_len, d_model)
        self.blocks = nn.Sequential(*[TransformerBlock(d_model, n_heads, 4*d_model) 
                                       for _ in range(n_layers)])
        self.norm = nn.LayerNorm(d_model)
        self.head = nn.Linear(d_model, vocab_size, bias=False)
    
    def forward(self, idx, targets=None):
        B, T = idx.shape
        tok = self.token_emb(idx)
        pos = self.pos_emb(torch.arange(T, device=idx.device))
        x = self.blocks(tok + pos)
        logits = self.head(self.norm(x))
        loss = None
        if targets is not None:
            loss = F.cross_entropy(logits.view(-1, logits.size(-1)), targets.view(-1))
        return logits, loss
```

### EJ-08-007: Entrenar nanoGPT
Dataset: tiny_shakespeare.txt (~1MB de texto)
Config: d_model=256, n_heads=4, n_layers=4, max_len=128
Meta: loss < 1.5 en validación, texto generado coherente

### EJ-08-008: Generar texto
```python
def generate(model, idx, max_new_tokens, temperature=1.0):
    for _ in range(max_new_tokens):
        idx_cond = idx[:, -model.max_len:]
        logits, _ = model(idx_cond)
        logits = logits[:, -1, :] / temperature
        probs = F.softmax(logits, dim=-1)
        idx_next = torch.multinomial(probs, num_samples=1)
        idx = torch.cat([idx, idx_next], dim=1)
    return idx
```
Experimentar con temperature=0.5, 1.0, 1.5. ¿Cómo afecta al texto?
