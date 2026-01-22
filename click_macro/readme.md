# go 매크로 작성

robotgo 라이브러리를 사용해서, 예약 도움 도구 만드는 프로젝트

## 사용 라이브러리
- gohook
- robotgo
- gosseract (Tesseract OCR)

## 사전 설치

### GCC
- macOS: Xcode Command Line Tools
- Windows: tdm gcc 설치 (https://jmeubank.github.io/tdm-gcc/)

### Tesseract OCR (이미지 텍스트 추출)

macOS에서 Homebrew를 사용하여 설치:

```bash
# Tesseract OCR 및 의존성(Leptonica 포함) 설치
brew install tesseract

# 한국어 포함 언어팩 설치
brew install tesseract-lang
```

설치 확인:

```bash
# Tesseract 버전 확인
tesseract --version

# 설치된 언어 확인
tesseract --list-langs
```

## 빌드 및 실행

### CGO 플래그 설정

gosseract 라이브러리는 CGO를 통해 Tesseract 및 Leptonica와 링크됩니다. 적절한 플래그를 설정해야 합니다:

**방법 1: 쉘 환경변수 설정**

`~/.zshrc` 또는 `~/.bashrc`에 추가:

```bash
export CGO_CPPFLAGS="-I/opt/homebrew/opt/leptonica/include -I/opt/homebrew/opt/tesseract/include"
export CGO_LDFLAGS="-L/opt/homebrew/opt/leptonica/lib -L/opt/homebrew/opt/tesseract/lib"
```

쉘 재시작:
```bash
source ~/.zshrc  # 또는 source ~/.bashrc
```

**방법 2: 명령어마다 플래그 설정**

```bash
CGO_CPPFLAGS="-I/opt/homebrew/opt/leptonica/include -I/opt/homebrew/opt/tesseract/include" \
CGO_LDFLAGS="-L/opt/homebrew/opt/leptonica/lib -L/opt/homebrew/opt/tesseract/lib" \
go test -v ./macro
```

**방법 3: pkg-config 사용**

```bash
PKG_CONFIG_PATH="/opt/homebrew/lib/pkgconfig:/opt/homebrew/opt/tesseract/lib/pkgconfig:/opt/homebrew/opt/leptonica/lib/pkgconfig" \
go test -v ./macro
```

## Tesseract OCR 사용법

### 기본 사용법

```go
import "github.com/otiai10/gosseract/v2"

func extractText() {
    client := gosseract.NewClient()
    defer client.Close()

    // 이미지 파일 경로 설정
    client.SetImage("path/to/image.png")

    // OCR 수행
    text, err := client.Text()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(text)
}
```

### 한국어 텍스트 인식

한국어 텍스트를 인식하려면 언어를 "kor"로 설정:

```go
client := gosseract.NewClient()
defer client.Close()

// 언어를 한국어로 설정
client.SetLanguage("kor")

// 이미지 파일 경로 설정
client.SetImage("last_msg.png")

// OCR 수행
text, err := client.Text()
if err != nil {
    fmt.Println("Error:", err)
    return
}

fmt.Println(text)
```

### 다중 언어 인식

여러 언어를 동시에 인식:

```go
// 한국어와 영어 모두 인식
client.SetLanguage("kor+eng")
```

### 테스트 실행

```bash
# 특정 테스트 실행
go test -v -run TestCaptureText ./macro

# 모든 테스트 실행
go test -v ./macro
```

## 문제 해결

### 에러: `leptonica/allheaders.h file not found`

Tesseract 또는 Leptonica가 설치되지 않았거나, CGO가 헤더를 찾을 수 없는 경우입니다:
1. Tesseract 설치: `brew install tesseract`
2. CGO 플래그가 올바르게 설정되었는지 확인 (빌드 및 실행 섹션 참조)

### 에러: `library 'lept' not found`

링크 문제입니다. go.mod에서 `github.com/otiai10/gosseract/v2` (v1이 아닌)를 사용하는지 확인하세요.

### 한국어 텍스트가 깨져서 나옴

언어팩이 설치되었는지 확인:
```bash
brew install tesseract-lang
```

코드에서 언어 설정:
```go
client.SetLanguage("kor")
```

## 사용팁
- 마우스 이동 시에는 중간중간 time.Sleep 필요 (0.3초 권장)
- 이벤트 등록 시, 키보드 이벤트는 `hook.AddEvent`, 마우스 이벤트는 `hook.AddMouse`
- gosseract는 반드시 v2 버전 사용: `github.com/otiai10/gosseract/v2`
