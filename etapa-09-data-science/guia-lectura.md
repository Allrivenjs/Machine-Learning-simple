# Etapa 9: Data Science

Duración: 1 mes | Prerequisito: Etapa 5

---

## Objetivo

Dominar el flujo completo de un proyecto de Data Science: desde datos crudos hasta modelo en producción. Con énfasis en Gradient Boosting, el algoritmo más usado en competencias y empresas.

---

## Recursos

### SQL
- Modo Guru en SQLZoo (gratuito): https://sqlzoo.net
- **Learning SQL** (Alan Beaulieu) — Capítulos 8-10

### Feature Engineering
- **Feature Engineering for Machine Learning** (Zheng & Casari)
- Kaggle competitions — analizar notebooks ganadores

### Gradient Boosting
- **StatQuest — Gradient Boost** (YouTube, serie de 4 videos)
- XGBoost paper: "XGBoost: A Scalable Tree Boosting System"
- Documentación de LightGBM

---

## Temas a dominar

### SQL Avanzado
- [ ] Window Functions: `ROW_NUMBER`, `RANK`, `LAG`, `LEAD`
- [ ] CTEs (Common Table Expressions): `WITH`
- [ ] Subqueries correlacionadas
- [ ] `HAVING` vs `WHERE`
- [ ] OLAP: CUBE, ROLLUP
- [ ] Optimización de queries: índices, EXPLAIN

### Feature Engineering
- [ ] Encoding de categóricas: Label, One-Hot, Target Encoding
- [ ] Scaling: StandardScaler, MinMaxScaler, RobustScaler
- [ ] Normalización de distribuciones: log, sqrt, Box-Cox
- [ ] Binning / discretización
- [ ] Feature interactions: productos, ratios
- [ ] Fechas: extraer día, mes, día de semana, días hasta evento
- [ ] Selección de features: mutual information, permutation importance

### EDA (Exploratory Data Analysis)
- [ ] Distribuciones univariadas: histogramas, boxplots
- [ ] Relaciones bivariadas: scatter, violin
- [ ] Outliers: IQR, z-score
- [ ] Correlaciones: Pearson, Spearman
- [ ] Missing values: patrones de missing data

### Gradient Boosting
- [ ] Intuición: ensemble de árboles que corrigen errores residuales
- [ ] XGBoost: regularización, manejo de missing, GPU
- [ ] LightGBM: Leaf-wise growth, binning de features
- [ ] CatBoost: manejo nativo de categóricas
- [ ] Hiperparámetros clave: n_estimators, max_depth, learning_rate, subsample
- [ ] Optuna: optimización bayesiana de hiperparámetros

---

## Señales de que dominas la etapa

- Puedes hacer EDA completo y encontrar insights accionables en <30 minutos
- Tu pipeline de Feature Engineering mejora el baseline en >5% accuracy
- Entiendes por qué LightGBM es más rápido que XGBoost
- Puedes llegar al top 20% en una competencia Kaggle de tabular data
