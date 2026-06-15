# Ejercicios: Etapa 9 — Data Science

---

## SQL Avanzado

### EJ-09-001: Window Functions
Dado una tabla de ventas (id, producto, fecha, monto):
```sql
-- 1. Ranking de ventas por producto dentro de cada mes
SELECT producto, fecha, monto,
       RANK() OVER (PARTITION BY DATE_TRUNC('month', fecha) ORDER BY monto DESC) as rank_mes
FROM ventas;

-- 2. Ventas del día anterior (LAG)
SELECT fecha, monto,
       LAG(monto, 1) OVER (ORDER BY fecha) as monto_ayer,
       monto - LAG(monto, 1) OVER (ORDER BY fecha) as delta
FROM ventas;

-- 3. Media móvil 7 días
SELECT fecha, monto,
       AVG(monto) OVER (ORDER BY fecha ROWS BETWEEN 6 PRECEDING AND CURRENT ROW) as ma7
FROM ventas;
```

### EJ-09-002: CTEs complejos
Encontrar clientes que compraron en el mes 1 pero no en el mes 2:
```sql
WITH mes1 AS (
    SELECT DISTINCT cliente_id FROM compras
    WHERE DATE_TRUNC('month', fecha) = '2024-01-01'
),
mes2 AS (
    SELECT DISTINCT cliente_id FROM compras
    WHERE DATE_TRUNC('month', fecha) = '2024-02-01'
)
SELECT m1.cliente_id
FROM mes1 m1
LEFT JOIN mes2 m2 ON m1.cliente_id = m2.cliente_id
WHERE m2.cliente_id IS NULL;
```
Extender: calcular cuánto gastaron estos clientes en el mes 1.

---

## Feature Engineering

### EJ-09-003: Encoding pipeline
Para el dataset Titanic:
```python
from sklearn.preprocessing import LabelEncoder, OneHotEncoder
from sklearn.impute import SimpleImputer
from sklearn.pipeline import Pipeline
from sklearn.compose import ColumnTransformer

numeric_features = ['Age', 'Fare', 'SibSp']
categorical_features = ['Sex', 'Embarked', 'Pclass']

preprocessor = ColumnTransformer(transformers=[
    ('num', Pipeline([
        ('imputer', SimpleImputer(strategy='median')),
        ('scaler', StandardScaler())
    ]), numeric_features),
    ('cat', Pipeline([
        ('imputer', SimpleImputer(strategy='most_frequent')),
        ('encoder', OneHotEncoder(handle_unknown='ignore'))
    ]), categorical_features)
])
```

### EJ-09-004: Feature interactions
Para un dataset de precios de casas:
- Crear feature: `precio_por_m2 = precio / superficie`
- Crear feature: `habitaciones_por_baño = habitaciones / baños`
- Crear features de fecha: `año`, `mes`, `dia_semana`, `es_fin_de_semana`
- Medir impacto en el modelo con permutation importance

### EJ-09-005: Target Encoding
Implementar target encoding con k-fold para evitar data leakage:
```python
def target_encode_kfold(df, col, target, n_splits=5):
    """Target encoding con cross-fitting para evitar leakage"""
    ...
```
Comparar contra One-Hot encoding en una variable con alta cardinalidad.

---

## EDA

### EJ-09-006: EDA completo
Usar el dataset de precios de casas (House Prices Kaggle):
1. Distribución del target (precio) — ¿necesita log transform?
2. Heatmap de correlaciones — top 10 features más correlacionadas con precio
3. Boxplots de precio por barrio/zona
4. Missing values: ¿qué patrón tienen? ¿son MCAR, MAR, MNAR?
5. Outliers: identificar y decidir si eliminar o winsorizar

---

## Gradient Boosting

### EJ-09-007: XGBoost baseline
```python
import xgboost as xgb
from sklearn.model_selection import cross_val_score

model = xgb.XGBClassifier(
    n_estimators=100,
    max_depth=4,
    learning_rate=0.1,
    subsample=0.8,
    colsample_bytree=0.8,
    random_state=42
)
scores = cross_val_score(model, X, y, cv=5, scoring='accuracy')
print(f"XGB: {scores.mean():.4f} ± {scores.std():.4f}")
```
Comparar con RandomForest y LogisticRegression del mismo dataset.

### EJ-09-008: Optuna hyperparameter search
```python
import optuna

def objective(trial):
    params = {
        'n_estimators': trial.suggest_int('n_estimators', 50, 500),
        'max_depth': trial.suggest_int('max_depth', 2, 10),
        'learning_rate': trial.suggest_float('learning_rate', 1e-4, 0.3, log=True),
        'subsample': trial.suggest_float('subsample', 0.5, 1.0),
        'colsample_bytree': trial.suggest_float('colsample_bytree', 0.5, 1.0),
    }
    model = xgb.XGBClassifier(**params)
    return cross_val_score(model, X_train, y_train, cv=3).mean()

study = optuna.create_study(direction='maximize')
study.optimize(objective, n_trials=100)
print(study.best_params)
```

### EJ-09-009: Feature Importance
```python
# 3 tipos de importancia
model.fit(X_train, y_train)

# 1. Gain importance (default XGBoost)
xgb.plot_importance(model, importance_type='gain')

# 2. Permutation importance
from sklearn.inspection import permutation_importance
result = permutation_importance(model, X_test, y_test, n_repeats=10)

# 3. SHAP values
import shap
explainer = shap.TreeExplainer(model)
shap_values = explainer.shap_values(X_test)
shap.summary_plot(shap_values, X_test)
```

---

## Proyecto integrador — Kaggle competition

Participar en "House Prices: Advanced Regression Techniques":
1. EDA completo (EJ-09-006)
2. Feature engineering (EJ-09-004)
3. Modelo base: LinearRegression
4. Modelo avanzado: XGBoost con Optuna
5. Ensemble: promedio de XGBoost + LightGBM + CatBoost
6. Submit a Kaggle — meta: top 30%
