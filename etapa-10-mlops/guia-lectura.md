# Etapa 10: MLOps

Duración: 1 mes | Prerequisito: Etapas 1-9

---

## Objetivo

Llevar modelos a producción. Con tu background en backend y DevOps, esta etapa será más natural. Foco en las herramientas específicas de ML que no existen en DevOps tradicional.

---

## Recursos

- **Designing Machine Learning Systems** (Chip Huyen) — libro esencial de MLOps
- Documentación oficial de MLflow: mlflow.org
- FastAPI docs: fastapi.tiangolo.com
- **Full Stack Deep Learning** — fullstackdeeplearning.com (gratuito)

---

## Prioridad de herramientas

### Tier 1 — Core (dominar)
1. **MLflow** — experiment tracking
2. **Docker** — empaquetado de modelos
3. **FastAPI** — servir modelos como API

### Tier 2 — Importante (entender bien)
4. **DVC** — versionado de datasets
5. **Airflow** — pipelines de entrenamiento
6. **GitHub Actions** — CI/CD para ML

### Tier 3 — Especializado (referencia)
7. **Kafka** — streaming en tiempo real
8. **Feast** — feature store
9. **Kubernetes** — despliegue a escala
10. **Qdrant/Pinecone** — vector databases (para RAG)

---

## Temas a dominar

### MLflow
- [ ] `mlflow.log_param`, `mlflow.log_metric`, `mlflow.log_artifact`
- [ ] Experiments y Runs
- [ ] Model Registry: staging, production, archived
- [ ] `mlflow.sklearn.log_model`, `mlflow.pytorch.log_model`
- [ ] Comparar runs en la UI

### DVC
- [ ] `dvc init`, `dvc add data/`
- [ ] Remote storage: S3, GCS, local
- [ ] `dvc repro` — reproducir pipeline
- [ ] `dvc params.yaml` — hiperparámetros versionados

### Docker para ML
- [ ] Dockerfile para modelo Python
- [ ] `COPY` del modelo serializado
- [ ] Variables de entorno para configuración
- [ ] Multi-stage builds (imagen pequeña)
- [ ] `docker-compose` para servicio + base de datos

### FastAPI
- [ ] Endpoint `POST /predict`
- [ ] Pydantic models para input/output validation
- [ ] Cargar modelo en startup (`@app.on_event("startup")`)
- [ ] Background tasks para logging
- [ ] Health check endpoint

### CI/CD para ML
- [ ] GitHub Actions: correr tests al hacer push
- [ ] Evaluar modelo en test set, fallar si accuracy < threshold
- [ ] Build y push Docker image automáticamente
- [ ] Deploy a Cloud Run / Railway / Fly.io

### Monitoring
- [ ] Data drift: ¿el input distribution cambió?
- [ ] Model drift: ¿el accuracy degradó?
- [ ] Prometheus + Grafana básico

---

## Señales de que dominas la etapa

- Puedes desplegar cualquier modelo como API en <1 hora
- Tu pipeline de ML es reproducible: `dvc repro` recrea resultados exactos
- Tienes CI/CD que impide deploy si el modelo degrada
- Entiendes qué es data drift y cómo detectarlo
