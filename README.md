# bigraytracer

Bigraytracer that utilizes advanced raytracing techniques with the aim to produce hyperrealistic images. It is composed of several subsections.

| Folder | Purpose |
| --- | --- |
| `client` | A web client for designing the raytracing scene. |
| `server` | A Java Spring Boot server that communicates between `client` editor via REST and one of `gamma` or `stingray` engines via gRPC. |
| `stingray` | A GPU-based raytracing engine written in C++. |
| `gamma` | A CPU-based raytracing engine written in Go. |