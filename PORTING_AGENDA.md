# 14-Day Porting Agenda: SATUSEHAT Integration (Go/Node/Python)
Goal: Catch up to PHP v4.9.0 features by EOD Day 14.

## Days 1-3: Core DataType Infrastructure
- Day 1: Identifier, Reference, CodeableConcept, Coding
- Day 2: Quantity, Age, Range, Period
- Day 3: Address, ContactPoint, HumanName, Attachment

## Days 4-6: Request & Auth Layer
- Day 4: SSRequest (Client/Retry/429 Backoff)
- Day 5: OAuth2Client + Token Persistence
- Day 6: DataType remaining (Annotation, Ratio, SampledData, Signature, etc.)

## Days 7-9: PayloadBuilders (Fluent API)
- Day 7: PatientBuilder, OrganizationBuilder
- Day 8: EncounterBuilder, PractitionerBuilder
- Day 9: ObservationBuilder, DiagnosticReportBuilder

## Days 10-12: Advanced Operations (Queue/Monitoring)
- Day 10: RateLimiter (Porting logic)
- Day 11: ErrorClassifier (HTTP 2xx-5xx mapping)
- Day 12: QueueWorker + QueueMonitor (Stats/DLQ)

## Days 13-14: Polish & Integration
- Day 13: Unit Testing suite (Vitest/GoTest/PyTest)
- Day 14: Documentation refresh & v1.0.0 Stable tagging

---

# Integrated Development Agenda: IVAI Ecosystem
## Phase A: Backend & API
- Task 1: Update `ivai.app` (Sanctum scopes, API routing, permission migration for `wish.ivai.app`).
- Task 2: Build News API (CMS adjustment for automated publishing).

## Phase B: Wish.ivai.app Frontend
- Task 1: Scaffold reactive UI (mobile-first, TWA-ready, Clinix-style nav/menu).
- Task 2: Integrate with `ivai.app` BFF layer.
