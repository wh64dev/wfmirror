# wfmirror - Personal Mirror Server Project
- 단순한 개인 파일공유 및 배포를 위한 미러서버 사이트 (GIN + Next.js + SQLite)

## Required
- Go >= 1.22.2
- Node v18 or later

## Feature
- account
    - digest login
    - jwt token authentication
    - account information will saved inner db file
- drag & drop
    - admin authentication required
    - just drag and drop file in current directory page
- private storage
    - show login page
    - define multiple directory in inner db file
- env configuration
