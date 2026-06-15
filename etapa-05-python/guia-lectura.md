# Etapa 5: Python Científico

Duración: 1 mes | Prerequisito: Etapas 1-4

---

## Objetivo

Aprender las herramientas estándar del ecosistema Python para ML. Comparar cada resultado con tu implementación Go. El objetivo es saber cuándo usar las librerías y entender qué hacen internamente.

---

## NumPy

### Recursos
- Documentación oficial: https://numpy.org/doc/
- **Python Data Science Handbook** (VanderPlas) — Capítulo 2 (gratuito en GitHub)

### Temas
- [ ] ndarray vs listas Python
- [ ] dtype, shape, reshape
- [ ] Broadcasting rules
- [ ] Operaciones vectorizadas (sin loops)
- [ ] `np.dot`, `np.matmul`, `@` operator
- [ ] Slicing y fancy indexing
- [ ] `np.linalg`: inv, det, eig, svd, solve
- [ ] Generación de datos: `np.random`, `np.linspace`

---

## Pandas

### Recursos
- Documentación oficial
- **Python Data Science Handbook** — Capítulo 3

### Temas
- [ ] Series y DataFrame
- [ ] Lectura de CSV: `pd.read_csv`
- [ ] Selección: `loc`, `iloc`, `[]`
- [ ] Filtros booleanos
- [ ] `groupby` + `agg`
- [ ] `merge` y `join`
- [ ] Pivot tables
- [ ] Manejo de NaN: `dropna`, `fillna`
- [ ] `apply` para transformaciones custom
- [ ] `value_counts`, `describe`

---

## Matplotlib y Seaborn

### Temas
- [ ] `plt.plot`, `plt.scatter`, `plt.hist`
- [ ] Subplots: `plt.subplots`
- [ ] Guardar figuras: `plt.savefig`
- [ ] `sns.heatmap` para matrices de correlación
- [ ] `sns.pairplot` para EDA rápido
- [ ] `sns.boxplot`, `sns.violinplot`

---

## Scikit-Learn

### Recursos
- Documentación oficial: https://scikit-learn.org
- **Hands-On ML** (Géron) Capítulos 2-6

### Temas
- [ ] API: `fit`, `predict`, `score`
- [ ] `train_test_split`, `cross_val_score`
- [ ] Pipelines: `Pipeline`, `ColumnTransformer`
- [ ] Preprocessing: `StandardScaler`, `LabelEncoder`, `OneHotEncoder`
- [ ] `GridSearchCV`, `RandomizedSearchCV`
- [ ] Métricas: `classification_report`, `confusion_matrix`

---

## Comparación Go vs Python

Para cada algoritmo implementado en Etapa 3:
1. Entrenar con tu implementación Go
2. Entrenar con Scikit-Learn en Python
3. Comparar pesos, parámetros y predictions
4. Documentar diferencias encontradas

---

## Señales de que dominas la etapa

- Puedes hacer EDA completo en menos de 10 líneas de Pandas
- Entiendes por qué NumPy broadcasting es más rápido que loops
- Puedes reproducir cualquier resultado de Scikit-Learn en tu Go
- Sabes en qué casos preferirías Go sobre Python (producción, latencia)
