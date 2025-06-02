
## ⚠️ Deprecated!

**This repository is deprecated.** <br />
A more up-to-date version is available here: https://github.com/grafana/grafana-plugin-examples/tree/master/examples/app-basic

---

# Grafana App Plugin Template

[![Build](https://github.com/grafana/grafana-starter-app/workflows/CI/badge.svg)](https://github.com/grafana/grafana-starter-app/actions?query=workflow%3A%22CI%22)

This template is a starting point for building Grafana App Plugins in Grafana 7.0+

## Getting started

### Frontend

1. Install dependencies

   ```bash
   yarn install --pure-lockfile
   ```

2. Build plugin in development mode or run in watch mode

   ```bash
   NODE_OPTIONS="--openssl-legacy-provider" yarn dev
   ```

   or

   ```bash
   NODE_OPTIONS="--openssl-legacy-provider" yarn watch
   ```

3. Build plugin in production mode

   ```bash
   NODE_OPTIONS="--openssl-legacy-provider" yarn build
   ```

### Backend

1. Update [Grafana plugin SDK for Go](https://grafana.com/docs/grafana/latest/developers/plugins/backend/grafana-plugin-sdk-for-go/) dependency to the latest minor version:

   ```bash
   go get -u github.com/grafana/grafana-plugin-sdk-go
   go mod tidy
   ```

2. Build backend plugin binaries for Linux, Windows and Darwin:

   ```bash
   mage -v
   ```

3. List all available Mage targets for additional commands:

   ```bash
   mage -l
   ```

### Run the server

1. To launch the development server using Docker, run:

   ```bash
   docker compose up --build
   ```

## Learn more

- [Build a app plugin tutorial](https://grafana.com/tutorials/build-a-app-plugin)
- [Grafana documentation](https://grafana.com/docs/)
- [Grafana Tutorials](https://grafana.com/tutorials/) - Grafana Tutorials are step-by-step guides that help you make the most of Grafana
- [Grafana UI Library](https://developers.grafana.com/ui) - UI components to help you build interfaces using Grafana Design System
