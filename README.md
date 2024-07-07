# wfmirror - Personal Mirror Server Project
- Simple personal file mirror service (GIN + Next.js)

## Required
- go >= 1.22.2
- node >= 20.15.0
- pnpm >= 9.4.0
- OS environment
    - MacOS
    - Linux
    - Windows (Not supported yet)

## Feature
- env config
- file preview
- disk usage indicator

## How to use
1. Clone current project first
```bash
~$ git clone https://github.com/wh64dev/wfmirror.git
```

2. Build backend service
```bash
~$ pnpm run app-build
```

3. Build frontend service
```bash
~$ pnpm build
```

4. Edit .env configuration
```env
SERVICE_PORT=8085
SERVICE_NAME=Example's WF Mirror
DATA_DIR=data
ALLOW_ORIGIN=http://localhost:3000
SERVER_URL=http://localhost:8085
```

4. Run both services
- Backend
```bash
./wfmirror
```
- Frontend
```bash
pnpm start
```
