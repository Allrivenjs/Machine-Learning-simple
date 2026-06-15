# Ejercicios: Etapa 10 — MLOps

---

## MLflow

### EJ-10-001: Primer experimento
```python
import mlflow
import mlflow.sklearn

with mlflow.start_run():
    # Hiperparámetros
    mlflow.log_param("n_estimators", 100)
    mlflow.log_param("max_depth", 4)
    
    # Entrenar
    model = RandomForestClassifier(n_estimators=100, max_depth=4)
    model.fit(X_train, y_train)
    
    # Métricas
    accuracy = model.score(X_test, y_test)
    mlflow.log_metric("accuracy", accuracy)
    
    # Modelo
    mlflow.sklearn.log_model(model, "random_forest")
```
Correr 5 veces con diferentes hiperparámetros. Comparar en la UI.

### EJ-10-002: Model Registry
Promover el mejor modelo a "Production":
```python
client = mlflow.tracking.MlflowClient()
client.create_registered_model("fraud_detector")

# Registrar modelo desde un run
model_uri = f"runs:/{best_run_id}/random_forest"
client.create_model_version("fraud_detector", model_uri, run_id=best_run_id)

# Promover a producción
client.transition_model_version_stage("fraud_detector", version=1, stage="Production")
```

---

## FastAPI

### EJ-10-003: API de predicción
```python
# app.py
from fastapi import FastAPI
from pydantic import BaseModel
import mlflow.sklearn
import numpy as np

app = FastAPI()
model = None

@app.on_event("startup")
def load_model():
    global model
    model = mlflow.sklearn.load_model("models:/fraud_detector/Production")

class PredictRequest(BaseModel):
    features: list[float]

class PredictResponse(BaseModel):
    prediction: int
    probability: float

@app.post("/predict", response_model=PredictResponse)
def predict(req: PredictRequest):
    X = np.array(req.features).reshape(1, -1)
    pred = int(model.predict(X)[0])
    prob = float(model.predict_proba(X)[0].max())
    return PredictResponse(prediction=pred, probability=prob)

@app.get("/health")
def health():
    return {"status": "ok", "model_loaded": model is not None}
```

### EJ-10-004: Testear la API
```bash
# Correr
uvicorn app:app --reload

# Testear
curl -X POST "http://localhost:8000/predict" \
  -H "Content-Type: application/json" \
  -d '{"features": [1.2, 3.4, 0.5, 2.1]}'
```
Escribir tests con `pytest` y `httpx`.

---

## Docker

### EJ-10-005: Dockerfile para modelo
```dockerfile
FROM python:3.11-slim

WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY app.py .
COPY model/ ./model/

EXPOSE 8000
CMD ["uvicorn", "app:app", "--host", "0.0.0.0", "--port", "8000"]
```
Build y testear localmente:
```bash
docker build -t ml-api:v1 .
docker run -p 8000:8000 ml-api:v1
```

### EJ-10-006: Docker Compose
```yaml
version: '3.8'
services:
  api:
    build: .
    ports:
      - "8000:8000"
    environment:
      - MLFLOW_TRACKING_URI=http://mlflow:5000
  
  mlflow:
    image: ghcr.io/mlflow/mlflow
    ports:
      - "5000:5000"
    volumes:
      - mlflow_data:/mlflow
```

---

## DVC

### EJ-10-007: Pipeline reproducible
```bash
dvc init
dvc add data/raw/dataset.csv

# Definir pipeline
dvc run -n preprocess \
  -d data/raw/dataset.csv \
  -o data/processed/clean.csv \
  python preprocess.py

dvc run -n train \
  -d data/processed/clean.csv \
  -o models/model.pkl \
  -M metrics/accuracy.json \
  python train.py
```
`dvc repro` debe reproducir exactamente los mismos resultados.

---

## CI/CD

### EJ-10-008: GitHub Actions
```yaml
# .github/workflows/ml-ci.yml
name: ML CI

on: [push]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Setup Python
        uses: actions/setup-python@v4
        with:
          python-version: '3.11'
      
      - name: Install
        run: pip install -r requirements.txt
      
      - name: Run tests
        run: pytest tests/
      
      - name: Evaluate model
        run: |
          python evaluate.py
          # Falla si accuracy < 0.90
          python -c "
          import json
          with open('metrics/accuracy.json') as f:
              acc = json.load(f)['accuracy']
          assert acc >= 0.90, f'Accuracy {acc} below threshold 0.90'
          "
```

---

## Proyecto final de etapa

Pipeline completo de producción para el modelo de la Etapa 9:
1. Dataset versionado con DVC
2. Entrenamiento loggeado en MLflow (10 experimentos, Optuna)
3. Mejor modelo promovido a "Production" en MLflow Registry
4. FastAPI que sirve predicciones desde MLflow
5. Dockerizado con docker-compose (api + mlflow server)
6. GitHub Actions: tests + evaluación automática en cada push
7. Monitoring: loggear distribution de inputs y predictions a un CSV

Bonus: deployar en Railway o Fly.io con un comando.
